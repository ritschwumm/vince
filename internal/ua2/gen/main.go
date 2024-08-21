package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	makeBot(root)
	makeClient(root)
	makeDevice(root)
	makeOs(root)
	makeVendor(root)
}

func makeVendor(root string) {
	var b bytes.Buffer
	fmt.Fprintln(&b, "// DO NOT EDIT Code generateVendord by ua/os/make_os.go")
	fmt.Fprintln(&b, " package ua2")
	generateVendor(&b, root, "vendorfragments.yml")
	r, err := format.Source(b.Bytes())
	if err != nil {
		fail("failed formatting go source ", err.Error())
	}
	os.WriteFile("ua_vendor.go", r, 0600)
}

type Vendor struct {
	name string
	re   []string
}

type VSLice []*Vendor

func (x VSLice) Len() int           { return len(x) }
func (x VSLice) Less(i, j int) bool { return x[i].name < x[j].name }
func (x VSLice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func generateVendor(b *bytes.Buffer, root, path string) {
	var m map[string][]string

	readUA(root, path, &m)

	var items []*Vendor
	for k, v := range m {
		items = append(items, &Vendor{
			name: k, re: v,
		})
	}
	sort.Sort(VSLice(items))

	var s bytes.Buffer
	var started bool
	for i, d := range items {
		if started {
			s.WriteByte('|')
		} else {
			started = true
		}
		if i != 0 {
			s.WriteByte('|')
		}
		for _, r := range d.re {
			s.WriteString(r)
		}
	}
	if IsStdRe(s.String()) {
		fmt.Fprintf(b, " var vendorAllRe= MatchRe(`%s`)\n", Clean(s.String()))
	} else {
		fmt.Fprintf(b, " var vendorAllRe= MatchRe2(`%s`)\n", Clean(s.String()))
	}
	fmt.Fprintf(b, "var vendorAll=[]*vendorRe{\n")
	var buf bytes.Buffer
	for _, d := range items {
		buf.Reset()
		s.Reset()
		for k, v := range d.re {
			if k != 0 {
				s.WriteByte('|')
			}
			s.WriteString(v)
		}
		r := Clean(s.String())
		if IsStdRe(s.String()) {
			fmt.Fprintf(&buf, "re:MatchRe(`%s`)", r)
		} else {
			fmt.Fprintf(&buf, "re: MatchRe2(`%s`)", r)
		}
		fmt.Fprintf(b, "{%s,name:%q", &buf, d.name)
		fmt.Fprintf(b, "},\n")
	}
	fmt.Fprintln(b, "}")
}

var allOs = map[string]struct{}{}
var allOsVersion = map[string]struct{}{}

func makeOs(root string) {
	var b bytes.Buffer

	fmt.Fprintln(&b, "// DO NOT EDIT Code generated by ua/os/make_os.go")
	fmt.Fprintln(&b, " package ua2")

	generate(&b, root, "oss.yml")
	var ls []string
	for k := range allOs {
		ls = append(ls, k)
	}
	sort.Strings(ls)
	var o bytes.Buffer
	json.NewEncoder(&o).Encode(ls)
	// os.WriteFile("os.json", o.Bytes(), 0600)
	ls = ls[:0]
	for k := range allOsVersion {
		ls = append(ls, k)
	}
	sort.Strings(ls)
	o.Reset()
	json.NewEncoder(&o).Encode(ls)
	// os.WriteFile("os_version.json", o.Bytes(), 0600)

	r, err := format.Source(b.Bytes())
	if err != nil {
		fail("failed  formatting go source ", err.Error())
	}
	os.WriteFile("ua_os.go", r, 0600)
}

type OsReg struct {
	Regex   string `yaml:"regex" json:"regex"`
	Name    string `yaml:"name" json:"name"`
	Version string `yaml:"version" json:"version"`
}

func generate(b *bytes.Buffer, root, path string) {
	var items []*OsReg
	readUA(root, path, &items)
	var s strings.Builder
	for i, d := range items {
		if !strings.Contains(d.Name, "$") {
			allOs[d.Name] = struct{}{}
		}
		if d.Version != "" && !strings.Contains(d.Version, "$") {
			allOsVersion[d.Version] = struct{}{}
		}
		if i != 0 {
			s.WriteByte('|')
		}
		s.WriteString(d.Regex)
	}
	if IsStdRe(s.String()) {
		fmt.Fprintf(b, " var osAllRe= MatchRe(`%s`)\n", Clean(s.String()))
	} else {
		fmt.Fprintf(b, " var osAllRe= MatchRe2(`%s`)\n", Clean(s.String()))
	}
	fmt.Fprintf(b, "var osAll=[]*clientRe{\n")
	var buf bytes.Buffer
	for _, d := range items {
		buf.Reset()
		r := Clean(d.Regex)
		if IsStdRe(d.Regex) {
			fmt.Fprintf(&buf, "re:MatchRe(`%s`)", r)
		} else {
			fmt.Fprintf(&buf, "re: MatchRe2(`%s`)", r)
		}
		fmt.Fprintf(b, "{%s,name:%q", &buf, d.Name)
		if d.Version != "" {
			fmt.Fprintf(b, ",version:%q", d.Version)
		}
		fmt.Fprintf(b, "},\n")
	}
	fmt.Fprintln(b, "}")
}

func makeDevice(root string) {
	var b bytes.Buffer

	fmt.Fprintln(&b, "// DO NOT EDIT Code generated by ua/device/make_device.go")
	fmt.Fprintln(&b, " package ua2")

	genCamera(&b, root)
	genCar(&b, root)
	genConsole(&b, root)
	genTV(&b, root)
	genMobile(&b, root)
	genNotebook(&b, root)
	genPortableMediaPlayer(&b, root)
	genShell(&b, root)
	r, err := format.Source(b.Bytes())
	if err != nil {
		fail("failed to format go source ", err.Error())
	}
	os.WriteFile("ua_device.go", r, 0600)
}

type Model struct {
	Regex  string `yaml:"regex" json:"regex"`
	Model  string `yaml:"model" json:"model"`
	Device string `yaml:"device" json:"device"` //mobile
	Brand  string `yaml:"brand" json:"brand"`   //mobile
}

type DeviceReg struct {
	Regex        string   `yaml:"regex" json:"regex"`
	Model        string   `yaml:"model" json:"model"`
	Device       string   `yaml:"device" json:"device"`
	Manufacturer string   `yaml:"-" json:"-"`
	Models       []*Model `yaml:"models" json:"models"`
}

type DeviceRegSlice []*DeviceReg

func (x DeviceRegSlice) Len() int           { return len(x) }
func (x DeviceRegSlice) Less(i, j int) bool { return x[i].Manufacturer < x[j].Manufacturer }
func (x DeviceRegSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func genCamera(b *bytes.Buffer, root string) {
	fmt.Fprintln(b)
	genericDevice(b, "Camera", root, "device/cameras.yml")
}

func genCar(b *bytes.Buffer, root string) {
	fmt.Fprintln(b)
	genericDevice(b, "Car", root, "device/car_browsers.yml")
}

func genConsole(b *bytes.Buffer, root string) {
	fmt.Fprintln(b)
	genericDevice(b, "Console", root, "device/consoles.yml")
}

func genTV(b *bytes.Buffer, root string) {
	fmt.Fprintln(b)
	fmt.Fprintf(b, " var deviceIsTVRe= MustCompile(`%s`)\n", Clean(`HbbTV/([1-9]{1}(?:.[0-9]{1}){1,2})`))
	genericDevice(b, "TV", root, "device/televisions.yml")
}

func genMobile(b *bytes.Buffer, root string) {
	fmt.Fprintln(b)
	genericDevice(b, "Mobile", root, "device/mobiles.yml")
}

func genNotebook(b *bytes.Buffer, root string) {
	fmt.Fprintln(b)
	genericDevice(b, "Notebook", root, "device/notebooks.yml")
}
func genPortableMediaPlayer(b *bytes.Buffer, root string) {
	fmt.Fprintln(b)
	genericDevice(b, "PortableMediaPlayer", root, "device/portable_media_player.yml")
}

func genShell(b *bytes.Buffer, root string) {
	fmt.Fprintln(b)
	fmt.Fprintf(b, " var deviceIsShellTV= MustCompile(`%s`)\n", Clean(`[a-z]+[ _]Shell[ _]\w{6}`))
	genericDevice(b, "Shell", root, "device/shell_tv.yml")
}

func genericDevice(b *bytes.Buffer, name string, root, path string) {
	items := loadDevice(root, path)
	var s strings.Builder
	for i, d := range items {
		if i != 0 {
			s.WriteByte('|')
		}
		s.WriteString(d.Regex)
	}
	var buf bytes.Buffer

	if IsStdRe(s.String()) {
		fmt.Fprintf(b, " var device%sAllRe= MatchRe(`%s`)\n", name, Clean(s.String()))
	} else {
		fmt.Fprintf(b, " var device%sAllRe= MatchRe2(`%s`)\n", name, Clean(s.String()))
	}
	fmt.Fprintf(b, "var device%sAll=[]*deviceRe{\n", name)
	for _, d := range items {
		buf.Reset()
		r := Clean(d.Regex)
		if IsStdRe(d.Regex) {
			fmt.Fprintf(&buf, "re:MatchRe(`%s`)", r)
		} else {
			fmt.Fprintf(&buf, "re: MatchRe2(`%s`)", r)
		}
		fmt.Fprintf(b, "{%s,company:%q,device:%q, ", &buf, d.Manufacturer, d.Device)
		if d.Model != "" {
			fmt.Fprintf(b, "model:%q},\n", d.Model)
		} else {
			if len(d.Models) > 0 {
				fmt.Fprintf(b, "models:[]*deviceModel{")
				for _, m := range d.Models {

					fmt.Fprintf(b, "{model:%q,", m.Model)
					buf.Reset()
					r = Clean(m.Regex)
					if IsStdRe(m.Regex) {
						fmt.Fprintf(&buf, "re:MatchRe(`%s`)", r)
					} else {
						fmt.Fprintf(&buf, "re: MatchRe2(`%s`)", r)
					}
					fmt.Fprintf(b, "%s},\n", &buf)
				}
				fmt.Fprintf(b, "},")
			}
			fmt.Fprintf(b, "},\n")
		}

	}
	fmt.Fprintln(b, "}")
}

func loadDevice(root, path string) []*DeviceReg {
	var v map[string]*DeviceReg
	readUA(root, path, &v)

	var items []*DeviceReg
	for key, value := range v {
		value.Manufacturer = key
		items = append(items, value)
	}
	sort.Sort(DeviceRegSlice(items))
	return items
}

var all = map[string]struct{}{}
var versions = map[string]struct{}{}

func makeClient(root string) {
	var b bytes.Buffer
	fmt.Fprintln(&b, "// DO NOT EDIT Code generated by ua/client/make_client.go")
	fmt.Fprintln(&b, " package ua2")
	genBrowserEngine(&b, root)
	genBrowser(&b, root)
	genFeedReader(&b, root)
	genLibraries(&b, root)
	genMediaPlayers(&b, root)
	genMobileApps(&b, root)
	genPim(&b, root)
	var ls []string
	for k := range all {
		ls = append(ls, k)
	}
	sort.Strings(ls)
	var o bytes.Buffer
	json.NewEncoder(&o).Encode(ls)
	// os.WriteFile("browsers.json", o.Bytes(), 0600)
	ls = ls[:0]
	for k := range versions {
		ls = append(ls, k)
	}
	sort.Strings(ls)
	o.Reset()
	json.NewEncoder(&o).Encode(ls)
	// os.WriteFile("browser_version.json", o.Bytes(), 0600)
	r, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("ua_client.go", r, 0600)
}

type ClientReg struct {
	Regex   string  `yaml:"regex" json:"regex"`
	Name    string  `yaml:"name" json:"name"`
	Version string  `yaml:"version" json:"version"`
	Url     string  `yaml:"url" json:"url"`
	Type    string  `yaml:"type" json:"type"`
	Engine  *Engine `yaml:"engine" json:"engine"`
}

type Engine struct {
	Default  string            `yaml:"default" json:"default"`
	Versions map[string]string `yaml:"versions" json:"versions"`
}

func genBrowserEngine(b *bytes.Buffer, root string) {
	generic(b, "BrowserEngine", root, "client/browser_engine.yml")
}

func genBrowser(b *bytes.Buffer, root string) {
	generic(b, "Browser", root, "client/browsers.yml")
}
func genFeedReader(b *bytes.Buffer, root string) {
	generic(b, "FeedReader", root, "client/feed_readers.yml")
}

func genLibraries(b *bytes.Buffer, root string) {
	generic(b, "Library", root, "client/libraries.yml")
}

func genMediaPlayers(b *bytes.Buffer, root string) {
	generic(b, "MediaPlayer", root, "client/mediaplayers.yml")
}
func genMobileApps(b *bytes.Buffer, root string) {
	generic(b, "MobileApp", root, "client/mobile_apps.yml")
}
func genPim(b *bytes.Buffer, root string) {
	generic(b, "Pim", root, "client/pim.yml")
}

func generic(b *bytes.Buffer, name string, root string, path string) {
	var items []*ClientReg
	readUA(root, path, &items)
	var s strings.Builder
	for i, d := range items {
		if !strings.Contains(d.Name, "$") {
			all[d.Name] = struct{}{}
		}
		if d.Version != "" && !strings.Contains(d.Version, "$") {
			versions[d.Version] = struct{}{}
		}
		// set kind
		if d.Type == "" {
			d.Type = name
		}
		if i != 0 {
			s.WriteByte('|')
		}
		s.WriteString(d.Regex)
	}
	if IsStdRe(s.String()) {
		fmt.Fprintf(b, " var client%sAllRe= MatchRe(`%s`)\n", name, Clean(s.String()))
	} else {
		fmt.Fprintf(b, " var client%sAllRe= MatchRe2(`%s`)\n", name, Clean(s.String()))
	}
	fmt.Fprintf(b, "var client%sAll=[]*clientRe{\n", name)
	var buf bytes.Buffer
	for _, d := range items {
		buf.Reset()
		r := Clean(d.Regex)
		if IsStdRe(d.Regex) {
			fmt.Fprintf(&buf, "re:MatchRe(`%s`)", r)
		} else {
			fmt.Fprintf(&buf, "re: MatchRe2(`%s`)", r)
		}
		fmt.Fprintf(b, "{%s,name:%q", &buf, d.Name)
		if d.Version != "" {
			fmt.Fprintf(b, ",version:%q", d.Version)
		}
		if d.Type != "" {
			fmt.Fprintf(b, ",kind:%q", d.Type)
		}
		if d.Url != "" {
			fmt.Fprintf(b, ",url:%q", d.Url)
		}
		if d.Engine != nil {
			fmt.Fprintf(b, ",engine:&clientEngine{")
			if d.Engine.Default != "" {
				fmt.Fprintf(b, "def:%q,", d.Engine.Default)
			}
			if d.Engine.Versions != nil {
				var ls []string
				for k := range d.Engine.Versions {
					ls = append(ls, k)
				}
				sort.Strings(ls)
				fmt.Fprintf(b, "versions:map[string]string{")
				for _, k := range ls {
					fmt.Fprintf(b, "%q:%q,\n", k, d.Engine.Versions[k])
				}
				fmt.Fprint(b, "},")
			}
			fmt.Fprint(b, "},")
		}
		fmt.Fprintf(b, "},\n")
	}
	fmt.Fprintln(b, "}")
}

func makeBot(root string) {
	type Producer struct {
		Name string `yaml:"name" json:"name"`
		Url  string `yaml:"url" json:"url"`
	}

	type Bot struct {
		Regex    string   `yaml:"regex" json:"regex"`
		Name     string   `yaml:"name" json:"name"`
		Category string   `yaml:"category" json:"category"`
		Url      string   `yaml:"url" json:"url"`
		Producer Producer `yaml:"producer" json:"producer"`
	}

	var r []*Bot
	readUA(root, "bots.yml", &r)

	var s bytes.Buffer

	var buf bytes.Buffer
	fmt.Fprintln(&buf, "// DO NOT EDIT Code generated by ua/bot/make_bot.go")
	fmt.Fprintln(&buf, " package ua2")
	for i, x := range r {
		if i != 0 {
			s.WriteByte('|')
		}
		s.WriteString(x.Regex)
	}

	if IsStdRe(s.String()) {
		fmt.Fprintf(&buf, " var allBotsReStandardMatch= MustCompile(`%s`)\n", Clean(s.String()))
	} else {
		fmt.Fprintf(&buf, " var allBotsReStandardMatch= MustCompile2(`%s`)\n", Clean(s.String()))
	}
	fmt.Fprintln(&buf, "var botsReList=[]*botRe{")
	for _, m := range r {
		s.Reset()
		if IsStdRe(m.Regex) {
			fmt.Fprintf(&s, "re:MatchRe(`%s`)", Clean(m.Regex))
		} else {
			fmt.Fprintf(&s, "re:MatchRe2(`%s`)", Clean(m.Regex))
		}
		fmt.Fprintf(&buf, "{%s, name:%q,category:%q,url:%q,producerName:%q,producerURL:%q, },\n",
			&s, m.Name, m.Category, m.Url, m.Producer.Name, m.Producer.Url,
		)
	}
	fmt.Fprintln(&buf, "}")
	f, err := format.Source(buf.Bytes())
	if err != nil {
		fail("failed to format go source ", err.Error())
	}
	os.WriteFile("ua_bots.go", f, 0600)
}

func readUA(root, name string, out any) {
	path := filepath.Join(root, name)
	f, err := os.ReadFile(path)
	if err != nil {
		fail("failed to read ua file ", path)
	}
	err = yaml.Unmarshal(f, out)
	if err != nil {
		fail("failed to  decode ", path, err.Error())
	}
}

func fail(msg string, args ...any) {
	slog.Error(msg, args...)
	os.Exit(1)
}

const meta = "\\.+*?()|[]{}^$#"

func IsRe(s string) bool {
	return strings.ContainsAny(s, meta)
}

func Clean(re string) string {
	rg := strings.Replace(re, `/`, `\/`, -1)
	rg = strings.Replace(rg, `++`, `+`, -1)
	rg = strings.Replace(rg, `\_`, `_`, -1)
	// if we find `\_` again, the original was `\\_`,
	// so restore that so the regex engine does not attempt to escape `_`
	rg = strings.Replace(rg, `\_`, `\\_`, -1)

	// only match if useragent begins with given regex or there is no letter before it
	return `(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:` + rg + ")"
}

func IsStdRe(s string) bool {
	r := Clean(s)
	_, err := regexp.Compile(r)
	return err == nil
}
