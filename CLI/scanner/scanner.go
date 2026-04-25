package scanner

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/yourorg/webscan/internal/models"
)

type Options struct {
	Timeout            time.Duration
	UserAgent          string
	InsecureSkipVerify bool
	FollowRedirects    bool
	MaxBodyBytes       int64
}

type Scanner struct {
	client *http.Client
	opts   Options
}

func New(opts Options) (*Scanner, error) {
	if opts.Timeout <= 0 {
		opts.Timeout = 15 * time.Second
	}
	if opts.UserAgent == "" {
		opts.UserAgent = "WebScan/1.0"
	}
	if opts.MaxBodyBytes <= 0 {
		opts.MaxBodyBytes = 1 << 20
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: opts.InsecureSkipVerify, // intentional for controlled testing only
		},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   opts.Timeout,
	}

	if !opts.FollowRedirects {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	return &Scanner{
		client: client,
		opts:   opts,
	}, nil
}

func (s *Scanner) Run(ctx context.Context, rawTarget string) (*models.Result, error) {
	target, err := normalizeTarget(rawTarget)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", s.opts.UserAgent)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	start := time.Now()
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(io.LimitReader(resp.Body, s.opts.MaxBodyBytes))
	title := extractTitle(body)

	headerChecks, missing := analyzeHeaders(resp.Header)
	tlsInfo := extractTLSInfo(resp.TLS)

	result := &models.Result{
		Target:          target,
		FinalURL:        resp.Request.URL.String(),
		StatusCode:      resp.StatusCode,
		ResponseTimeMS:   time.Since(start).Milliseconds(),
		Server:          resp.Header.Get("Server"),
		Title:           title,
		SecurityHeaders: headerChecks,
		MissingHeaders:  missing,
		TLS:             tlsInfo,
		Timestamp:       time.Now().UTC(),
	}

	return result, nil
}

func normalizeTarget(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", errors.New("target URL is required")
	}

	if !strings.Contains(raw, "://") {
		raw = "https://" + raw
	}

	u, err := url.Parse(raw)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return "", fmt.Errorf("unsupported scheme: %s", u.Scheme)
	}
	if u.Host == "" {
		return "", errors.New("URL host is empty")
	}

	return u.String(), nil
}

func extractTitle(body []byte) string {
	re := regexp.MustCompile(`(?is)<title[^>]*>(.*?)</title>`)
	m := re.FindSubmatch(body)
	if len(m) < 2 {
		return ""
	}
	title := strings.TrimSpace(string(m[1]))
	title = strings.ReplaceAll(title, "\n", " ")
	title = strings.ReplaceAll(title, "\t", " ")
	title = strings.Join(strings.Fields(title), " ")
	return title
}

func extractTLSInfo(state *tls.ConnectionState) *models.TLSInfo {
	if state == nil || len(state.PeerCertificates) == 0 {
		return &models.TLSInfo{Enabled: false}
	}

	cert := state.PeerCertificates[0]
	days := int(time.Until(cert.NotAfter).Hours() / 24)

	return &models.TLSInfo{
		Enabled:         true,
		Issuer:          cert.Issuer.String(),
		Subject:         cert.Subject.String(),
		NotBefore:       cert.NotBefore,
		NotAfter:        cert.NotAfter,
		DaysUntilExpiry: days,
	}
}
