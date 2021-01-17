package security

import (
	"net/http"

	"github.com/unrolled/secure"
)

// SecureHandler adds security headers to the HTTP response
func SecureHandler(production bool) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {

		options := secure.Options{
			IsDevelopment:        !production,
			FrameDeny:            true,
			ContentTypeNosniff:   true,
			BrowserXssFilter:     true,
			AllowedHostsAreRegex: true,
		}

		secureHandler := secure.New(options)

		return secureHandler.Handler(h)
	}
}
