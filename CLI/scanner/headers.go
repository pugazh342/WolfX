package scanner

import "net/http"

func checkHeaders(h http.Header) []string {
    required := []string{
        "Content-Security-Policy",
        "X-Frame-Options",
        "X-XSS-Protection",
        "Strict-Transport-Security",
    }

    var missing []string

    for _, r := range required {
        if h.Get(r) == "" {
            missing = append(missing, r)
        }
    }

    return missing
}
