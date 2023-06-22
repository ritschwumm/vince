package plug

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/vinceanalytics/vince/internal/params"
)

func Browser(ctx context.Context) Pipeline {
	return Pipeline{
		Firewall(ctx),
		FetchSession,
		FetchFlash,
		PutSecureBrowserHeaders,
		SessionTimeout,
		Auth,
		LastSeen,
	}
}

func SharedLink() Pipeline {
	return Pipeline{
		PutSecureBrowserHeaders,
	}
}

func Protect() Pipeline {
	return Pipeline{
		CSRF,
		Captcha,
	}
}

func API(ctx context.Context) Pipeline {
	return Pipeline{
		Firewall(ctx),
	}
}

func InternalStatsAPI() Pipeline {
	return Pipeline{
		AcceptJSON,
		FetchSession,
		AuthorizedSiteAccess(),
	}
}

func (p Pipeline) Re(exp string, method string, f func(w http.ResponseWriter, r *http.Request)) Plug {
	for k, v := range replace {
		exp = strings.ReplaceAll(exp, k, v)
	}
	re := regexp.MustCompile(exp)
	pipe := p.Pass(f)
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if checkMethod(method, r) && re.MatchString(r.URL.Path) {
				r = r.WithContext(params.Set(r.Context(),
					params.Re(re, r.URL.Path),
				))
				pipe.ServeHTTP(w, r)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func (p Pipeline) GET(re string, f func(w http.ResponseWriter, r *http.Request)) Plug {
	return p.Re(re, http.MethodGet, f)
}

func (p Pipeline) PUT(re string, f func(w http.ResponseWriter, r *http.Request)) Plug {
	return p.Re(re, http.MethodPut, f)
}

func (p Pipeline) POST(re string, f func(w http.ResponseWriter, r *http.Request)) Plug {
	return p.Re(re, http.MethodPost, f)
}

func (p Pipeline) DELETE(re string, f func(w http.ResponseWriter, r *http.Request)) Plug {
	return p.Re(re, http.MethodDelete, f)
}

func (p Pipeline) PathGET(path string, handler func(w http.ResponseWriter, r *http.Request)) Plug {
	return p.Path(http.MethodGet, path, handler)
}

func (p Pipeline) PathPOST(path string, handler func(w http.ResponseWriter, r *http.Request)) Plug {
	return p.Path(http.MethodPost, path, handler)
}

func (p Pipeline) PathPUT(path string, handler func(w http.ResponseWriter, r *http.Request)) Plug {
	return p.Path(http.MethodPut, path, handler)
}

func (p Pipeline) PathDELETE(path string, handler func(w http.ResponseWriter, r *http.Request)) Plug {
	return p.Path(http.MethodDelete, path, handler)
}

func (p Pipeline) Path(method, path string, handler func(w http.ResponseWriter, r *http.Request)) Plug {
	pipe := p.Pass(handler)
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if checkMethod(method, r) && r.URL.Path == path {
				pipe.ServeHTTP(w, r)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func (p Pipeline) Prefix(path string, handler func(w http.ResponseWriter, r *http.Request)) Plug {
	pipe := p.Pass(handler)
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, path) {
				pipe.ServeHTTP(w, r)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func PREFIX(prefix string, pipe ...Plug) Plug {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, prefix) {
				Pipeline(pipe).Pass(NOOP).ServeHTTP(w, r)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

var replace = map[string]string{
	":recipient":     "(?P<recipient>[^.]+)",
	":slug":          "(?P<slug>[^.]+)",
	":id":            "(?P<id>[^.]+)",
	":owner":         "(?P<owner>[^.]+)",
	":invitation_id": "(?P<id>[^.]+)",
	":new_role":      "(?P<new_role>[^.]+)",
	":site":          `(?P<site>(?:[a-z0-9]+(?:-[a-z0-9]+)*\.)+[a-z]{2,})`,
	":goal_id":       "(?P<goal_id>[^.]+)",
}

func Ok(ok bool, pipe Plug) Plug {
	if ok {
		return pipe
	}
	return func(h http.Handler) http.Handler {
		return h
	}
}

const form = "application/x-www-form-urlencoded"

func checkMethod(method string, r *http.Request) bool {
	if method == r.Method {
		return true
	}
	if r.Header.Get("content-type") == form {
		return r.FormValue("_method") == method
	}
	return false
}
