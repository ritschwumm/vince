package api

import (
	"bytes"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gernest/vince/country"
	"github.com/gernest/vince/gate"
	"github.com/gernest/vince/geoip"
	"github.com/gernest/vince/log"
	"github.com/gernest/vince/referrer"
	"github.com/gernest/vince/remoteip"
	"github.com/gernest/vince/timeseries"
	"github.com/gernest/vince/ua"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var ErrBadJSON = errors.New("api: invalid json ")

// Request is sent by the vince script embedded in client websites. For faster
// performance we use simd if available
type Request struct {
	EventName   string `json:"n"`
	URI         string `json:"url"`
	Referrer    string `json:"r"`
	Domain      string `json:"d"`
	ScreenWidth int    `json:"w"`
	HashMode    bool   `json:"h"`
}

// Parse opportunistic parses request body to r object. This is crucial method
// any gains here translates to smooth  events ingestion pipeline.
//
// A hard size limitation of 32kb is imposed. This is arbitrary value, any change
// to it must be be supported with statistics.
func (r *Request) Parse(body io.Reader) error {
	b := bufPool.Get().(*bytes.Buffer)
	defer func() {
		b.Reset()
		bufPool.Put(b)
	}()
	// My mom used to say, don't trust the internet. Never stream decode payloads
	// directly from strangers. We copy just enough then we  process.
	b.ReadFrom(io.LimitReader(body, 32<<10))
	return json.Unmarshal(b.Bytes(), r)
}

var bufPool = &sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

// Events accepts events payloads from vince client script.
func Events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	xlg := log.Get(r.Context())
	var req Request
	err := req.Parse(r.Body)
	if err != nil {
		xlg.Err(err).Msg("Failed decoding json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	remoteIp := remoteip.Get(r)
	if req.URI == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uri, err := url.Parse(req.URI)
	if err != nil {
		xlg.Err(err).Msg("Failed parsing url")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if uri.Scheme == "data" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	host := sanitizeHost(uri.Host)
	userAgent := r.UserAgent()

	reqReferrer := req.Referrer
	refUrl, err := url.Parse(reqReferrer)
	if err != nil {
		xlg.Err(err).Msg("Failed parsing referrer url")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	path := uri.Path
	if req.HashMode && uri.Fragment != "" {
		path += "#" + uri.Fragment
	}
	if len(path) > 2000 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if req.EventName == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if req.Domain == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var domains []string
	for _, d := range strings.Split(req.Domain, ",") {
		domains = append(domains, sanitizeHost(d))
	}

	query := uri.Query()
	now := time.Now()
	agent := ua.Parse(userAgent)
	// handle referrer
	ref := referrer.ParseReferrer(req.Referrer)
	source := query.Get("utm_source")
	if source == "" {
		source = query.Get("source")
	}
	if source == "" {
		source = query.Get("ref")
	}
	if source == "" {
		if ref != nil {
			if ref.Type == "unknown" {
				source = sanitizeHost(refUrl.Host)
			} else {
				source = ref.Type
			}
		}
	}
	reqReferrer = cleanReferrer(reqReferrer)

	var countryCode country.Code
	var cityGeonameId uint32
	if remoteIp != "" {
		ip := net.ParseIP(remoteIp)
		city, err := geoip.Lookup(ip)
		if err == nil {
			countryCode = country.Lookup(city.Country.IsoCode)
			cityGeonameId = uint32(city.Country.GeoNameID)
		}
	}
	var screenSize timeseries.ScreenSize
	switch {
	case req.ScreenWidth < 576:
		screenSize = timeseries.Mobile
	case req.ScreenWidth < 992:
		screenSize = timeseries.Tablet
	case req.ScreenWidth < 1440:
		screenSize = timeseries.Laptop
	case req.ScreenWidth >= 1440:
		screenSize = timeseries.Desktop
	}
	var dropped int
	for _, domain := range domains {
		b, pass := gate.Check(r.Context(), domain)
		if !pass {
			dropped += 1
			continue
		}
		userID := int64(seedID.Gen(remoteIp, userAgent, domain, host))
		e := new(timeseries.Event)
		e.UserId = userID
		e.Name = req.EventName
		e.Hostname = host
		e.Domain = domain
		e.Pathname = path
		e.UtmMedium = query.Get("utm_medium")
		e.UtmSource = query.Get("utm_source")
		e.UtmCampaign = query.Get("utm_campaign")
		e.UtmContent = query.Get("utm_content")
		e.UtmTerm = query.Get("utm_content")
		e.OperatingSystem = agent.Os
		e.OperatingSystemVersion = agent.OsVersion
		e.Browser = agent.Browser
		e.BrowserVersion = agent.BrowserVersion
		e.ReferrerSource = source
		e.Referrer = reqReferrer
		e.CountryCode = countryCode
		e.CityGeoNameID = cityGeonameId
		e.ScreenSize = screenSize
		e.Timestamp = now
		previousUUserID := int64(seedID.GenPrevious(remoteIp, userAgent, domain, host))
		e.SessionId = b.Register(r.Context(), e, previousUUserID)
	}
	if dropped > 0 {
		w.Header().Set("x-vince-dropped", strconv.Itoa(dropped))
		w.WriteHeader(http.StatusAccepted)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func sanitizeHost(s string) string {
	return strings.TrimPrefix(strings.TrimSpace(s), "www.")
}

func cleanReferrer(s string) string {
	r, _ := url.Parse(s)
	r.Host = sanitizeHost(r.Host)
	r.Path = strings.TrimSuffix(s, "/")
	return r.String()
}
