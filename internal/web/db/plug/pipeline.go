package plug

import (
	"net/http"
	"strings"

	"github.com/vinceanalytics/vince/app"
	"github.com/vinceanalytics/vince/internal/web/db"
)

type Handler func(db *db.Config, w http.ResponseWriter, r *http.Request)

type Middleware func(h Handler) Handler

type Pipeline []Middleware

func (p Pipeline) With(m ...Middleware) Pipeline {
	return append(p, m...)
}

func (p Pipeline) Then(h Handler) func(db *db.Config, w http.ResponseWriter, r *http.Request) {
	for i := range p {
		h = p[len(p)-1-i](h)
	}
	return h
}

func Browser() Pipeline {
	return Pipeline{
		FetchSession,
		FetchFlash,
		SecureHeaders,
		SessionTimeout,
		FetchFlash,
	}
}

func API() Pipeline {
	return Pipeline{}
}

func InternalStats() Pipeline {
	return Pipeline{
		FetchSession,
	}
}

func FetchSession(h Handler) Handler {
	return func(db *db.Config, w http.ResponseWriter, r *http.Request) {
		db.Load(w, r)
		h(db, w, r)
	}
}

func FetchFlash(h Handler) Handler {
	return func(db *db.Config, w http.ResponseWriter, r *http.Request) {
		db.Flash(w)
		h(db, w, r)
	}
}

var file = http.FileServerFS(app.Public)
var icons = http.FileServerFS(app.Images)
var scripts = http.FileServerFS(app.Scripts)

func Static(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/public/") {
			file.ServeHTTP(w, r)
			return
		}
		if strings.HasPrefix(r.URL.Path, "/js/") {
			w.Header().Set("x-content-type-options", "nosniff")
			w.Header().Set("cross-origin-resource-policy", "cross-origin")
			w.Header().Set("access-control-allow-origin", "*")
			w.Header().Set("cache-control", "public, max-age=86400, must-revalidate")
			scripts.ServeHTTP(w, r)
			return
		}
		if strings.HasPrefix(r.URL.Path, "/images/") {
			icons.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func SessionTimeout(h Handler) Handler {
	return func(db *db.Config, w http.ResponseWriter, r *http.Request) {
		db.SessionTimeout(w)
		h(db, w, r)
	}
}

func SecureHeaders(h Handler) Handler {
	return func(db *db.Config, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("x-frame-options", "SAMEORIGIN")
		w.Header().Set("x-xss-protection", "1; mode=block")
		w.Header().Set("x-content-type-options", "nosniff")
		w.Header().Set("x-download-options", "noopen")
		w.Header().Set("x-permitted-cross-domain-policies", "none")
		w.Header().Set("cross-origin-window-policy", "deny")
		h(db, w, r)
	}
}

func CSRF(h Handler) Handler {
	return func(db *db.Config, w http.ResponseWriter, r *http.Request) {
		db.SaveCsrf(w)
		h(db, w, r)
	}
}

func VerifyCSRF(h Handler) Handler {
	return func(db *db.Config, w http.ResponseWriter, r *http.Request) {
		if !db.IsValidCsrf(r) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		h(db, w, r)
	}
}

func RequireAccount(h Handler) Handler {
	return func(db *db.Config, w http.ResponseWriter, r *http.Request) {
		if !db.Authorize(w, r) {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		h(db, w, r)
	}
}

func RequireLogout(h Handler) Handler {
	return func(db *db.Config, w http.ResponseWriter, r *http.Request) {
		if !db.IsLoggedOut(w) {
			http.Redirect(w, r, "/sites", http.StatusFound)
			return
		}
		h(db, w, r)
	}
}

func AcceptJSON(h Handler) Handler {
	return func(db *db.Config, w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("content-type") != "application/json" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		h(db, w, r)
	}
}
