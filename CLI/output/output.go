package output

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/yourorg/webscan/internal/models"
)

func Print(result *models.Result, jsonOut bool, verbose bool) error {
	if jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(result)
	}

	fmt.Println("== WebScan Result ==")
	fmt.Printf("Target:          %s\n", result.Target)
	fmt.Printf("Final URL:       %s\n", result.FinalURL)
	fmt.Printf("Status Code:     %d\n", result.StatusCode)
	fmt.Printf("Response Time:   %d ms\n", result.ResponseTimeMS)
	if result.Server != "" {
		fmt.Printf("Server:          %s\n", result.Server)
	}
	if result.Title != "" {
		fmt.Printf("Title:           %s\n", result.Title)
	}

	if result.TLS != nil && result.TLS.Enabled {
		fmt.Println("TLS:             enabled")
		fmt.Printf("TLS Issuer:      %s\n", result.TLS.Issuer)
		fmt.Printf("TLS Subject:     %s\n", result.TLS.Subject)
		fmt.Printf("Expires In:      %d days\n", result.TLS.DaysUntilExpiry)
	} else {
		fmt.Println("TLS:             not detected")
	}

	if len(result.MissingHeaders) > 0 {
		fmt.Println("Missing Headers:")
		for _, h := range result.MissingHeaders {
			fmt.Printf("  - %s\n", h)
		}
	} else {
		fmt.Println("Missing Headers: none")
	}

	if verbose {
		fmt.Println()
		fmt.Println("Security Headers:")
		for _, h := range result.SecurityHeaders {
			status := "missing"
			if h.Present {
				status = "present"
			}
			val := h.Value
			if len(val) > 80 {
				val = val[:80] + "..."
			}
			if val != "" {
				fmt.Printf("  - %-24s %-8s %s\n", h.Name, status, strings.TrimSpace(val))
			} else {
				fmt.Printf("  - %-24s %-8s\n", h.Name, status)
			}
		}
	}

	return nil
}
