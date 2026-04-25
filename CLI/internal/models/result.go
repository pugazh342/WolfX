package models

import "time"

type HeaderCheck struct {
	Name    string `json:"name"`
	Present bool   `json:"present"`
	Value   string `json:"value,omitempty"`
}

type TLSInfo struct {
	Enabled         bool      `json:"enabled"`
	Issuer          string    `json:"issuer,omitempty"`
	Subject         string    `json:"subject,omitempty"`
	NotBefore       time.Time `json:"not_before,omitempty"`
	NotAfter        time.Time `json:"not_after,omitempty"`
	DaysUntilExpiry int       `json:"days_until_expiry,omitempty"`
}

type Result struct {
	Target          string        `json:"target"`
	FinalURL        string        `json:"final_url"`
	StatusCode      int           `json:"status_code"`
	ResponseTimeMS   int64         `json:"response_time_ms"`
	Server          string        `json:"server,omitempty"`
	Title           string        `json:"title,omitempty"`
	SecurityHeaders []HeaderCheck `json:"security_headers"`
	MissingHeaders  []string      `json:"missing_headers,omitempty"`
	TLS             *TLSInfo      `json:"tls,omitempty"`
	Timestamp       time.Time     `json:"timestamp"`
}
