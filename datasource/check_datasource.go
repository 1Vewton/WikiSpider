package datasource

import (
	"strings"
)

// CheckContains checks if url contains keyword
func CheckContains(
	url string,
	keyword string,
) bool {
	return strings.Contains(
		url,
		keyword,
	)
}
