package cmd

import (
	"context"
	"time"

	"github.com/yourorg/webscan/internal/output"
	"github.com/yourorg/webscan/internal/scanner"
	"github.com/spf13/cobra"
)

var (
	target            string
	timeout           time.Duration
	jsonOut           bool
	insecureTLS       bool
	followRedirects   bool
	userAgent         string
	maxBodyBytes      int64
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Run a passive scan against a target URL",
	Example: `webscan scan --url https://example.com
webscan scan -u example.com --json
webscan scan -u https://example.com --timeout 15s --verbose`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sc, err := scanner.New(scanner.Options{
			Timeout:           timeout,
			UserAgent:         userAgent,
			InsecureSkipVerify: insecureTLS,
			FollowRedirects:   followRedirects,
			MaxBodyBytes:      maxBodyBytes,
		})
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(cmd.Context(), timeout)
		defer cancel()

		result, err := sc.Run(ctx, target)
		if err != nil {
			return err
		}

		return output.Print(result, jsonOut, verbose)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&target, "url", "u", "", "target URL to scan")
	scanCmd.Flags().DurationVar(&timeout, "timeout", 15*time.Second, "request timeout")
	scanCmd.Flags().BoolVar(&jsonOut, "json", false, "print JSON output")
	scanCmd.Flags().BoolVar(&insecureTLS, "insecure", false, "skip TLS certificate verification")
	scanCmd.Flags().BoolVar(&followRedirects, "follow-redirects", true, "follow HTTP redirects")
	scanCmd.Flags().StringVar(&userAgent, "user-agent", "WebScan/1.0", "custom user agent")
	scanCmd.Flags().Int64Var(&maxBodyBytes, "max-body-bytes", 1<<20, "maximum response body bytes to read")

	_ = scanCmd.MarkFlagRequired("url")
}
