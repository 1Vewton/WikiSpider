package utils

import (
	"fmt"
	"reflect"
)

// GetValueFromFieldByString gets value from struct by field name
func GetValueFromFieldByString(
	data any,
	field string,
) (any, error) {
	reflectedData := reflect.ValueOf(data)
	value := reflectedData.FieldByName(field)
	if value.IsValid() {
		return value.Interface(), nil
	} else {
		return nil, fmt.Errorf(
			"%s field does not exists in this struct",
			field,
		)
	}
}
