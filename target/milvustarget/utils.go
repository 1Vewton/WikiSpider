package milvustarget

import (
	"fmt"
)

// CreateQueryTextString creates text for query string field
func CreateQueryTextString(
	field string,
	value string,
) string {
	processedValue := fmt.Sprintf(
		"%s%s",
		value,
		"%",
	)
	return fmt.Sprintf(
		"%s like \"%s\"",
		field,
		processedValue,
	)
}
