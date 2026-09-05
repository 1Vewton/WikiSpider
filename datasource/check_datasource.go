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

// GetDataSource for url
func GetDataSource(
	url string,
) DataSourceType {
	supported := GetAllSuportedTypes()
	for _, supportedType := range supported {
		if CheckContains(url, supportedType.ToString()) {
			return supportedType
		}
	}
	return NotSupported
}
