package query

import (
	"net/url"
)

// ExtractDomain extracts a domain from URL string.
func ExtractDomain(u string) string {
	if u == "" {
		return ""
	}

	parsed, err := url.Parse(u)
	if err != nil {
		return ""
	}

	return parsed.Hostname()
}
