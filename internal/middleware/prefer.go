package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var prefer = http.CanonicalHeaderKey("Prefer")

// Prefer is a middleware that checks if the client has set
// the Prefer header https://www.rfc-editor.org/rfc/rfc7240.html
//
// If the head is set they have provided the `wait` preference, the
// the server will wait for the specified period of time before responing
func Prefer(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if prefs := r.Header.Get(prefer); prefs != "" {
			// by default wait for 0 seconds
			wait := 0 * time.Second
			// parse the preferences, looking for 'wait=N'
			for _, pref := range strings.Split(prefs, ",") {
				parts := strings.Split(pref, "=")
				if strings.TrimSpace(parts[0]) == "wait" && len(parts) == 2 {
					i, err := strconv.Atoi(parts[1])
					if err != nil {
						panic(fmt.Sprintf("could not parse 'wait' preference from Prefer header: %s", err))
					}
					wait = time.Duration(i) * time.Second
				}
			}
			time.Sleep(wait)
		}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
