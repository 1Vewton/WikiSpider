package utils

import (
	"testing"
)

// TestStruct defines the data struct for testing
type TestStruct struct {
	Field1 string
}

// TestGetValueFromFieldByString tests the get value from field by string
func TestGetValueFromFieldByString(
	t *testing.T,
) {
	testData := TestStruct{
		Field1: "abc",
	}
	data, errGet := GetValueFromFieldByString(
		testData,
		"Field1",
	)
	if errGet != nil {
		t.Error(errGet)
	}
	if data != testData.Field1 {
		t.Errorf(
			"Expected %s, got %s",
			data,
			testData.Field1,
		)
	}
	_, errGet2 := GetValueFromFieldByString(
		testData,
		"field2",
	)
	if errGet2 == nil {
		t.Error("Expected error")
	}
}
