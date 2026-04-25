package scanner

import "net/http"

var requiredSecurityHeaders = []string{
	"Content-Security-Policy",
	"X-Frame-Options",
	"X-Content-Type-Options",
	"Referrer-Policy",
	"Permissions-Policy",
	"Strict-Transport-Security",
}

func analyzeHeaders(h http.Header) ([]HeaderCheck, []string) {
	checks := make([]HeaderCheck, 0, len(requiredSecurityHeaders))
	missing := make([]string, 0)

	for _, name := range requiredSecurityHeaders {
		val := h.Get(name)
		present := val != ""
		checks = append(checks, HeaderCheck{
			Name:    name,
			Present: present,
			Value:   val,
		})
		if !present {
			missing = append(missing, name)
		}
	}

	return checks, missing
}
