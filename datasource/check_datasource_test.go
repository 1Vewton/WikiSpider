package datasource

import (
	"testing"
)

// TestDataSourceTypeGetting tests the data source getting function
func TestDataSourceTypeGetting(t *testing.T) {
	t.Parallel()
	testDataSource1 := GetDataSource(
		"https://en.wikipedia.org/w/api.php?",
	)
	if testDataSource1 != Wikipedia {
		t.Errorf(
			"expected %s for data source of https://en.wikipedia.org/w/api.php?, got %s",
			Wikipedia,
			testDataSource1,
		)
	}
	testDataSource2 := GetDataSource(
		"https://minecraft.fandom.com/api.php?",
	)
	if testDataSource2 != Fandom {
		t.Errorf(
			"expected %s for data source of https://minecraft.fandom.com/api.php?, got %s",
			Fandom,
			testDataSource2,
		)
	}
}

// TestDataSourceTypeConverting tests the datasource type converting
func TestDataSourceTypeConverting(t *testing.T) {
	t.Parallel()
	dataSourceTestCase1 := Fandom
	result1 := dataSourceTestCase1.ToString()
	if result1 != "fandom" {
		t.Errorf(
			"expected %s, got %s",
			result1,
			"fandom",
		)
	}
	dataSourceTestCase2 := Wikipedia
	result2 := dataSourceTestCase2.ToString()
	if result2 != "wikipedia" {
		t.Errorf(
			"expected %s, got %s",
			result2,
			"wikipedia",
		)
	}
	dataSourceTestCase3 := "fandom"
	result3 := ToDataSourceType(dataSourceTestCase3)
	if result3 != Fandom {
		t.Errorf(
			"expected %s, got %s",
			result3.ToString(),
			Fandom.ToString(),
		)
	}
	dataSourceTestCase4 := "wikipedia"
	result4 := ToDataSourceType(dataSourceTestCase4)
	if result4 != Wikipedia {
		t.Errorf(
			"expected %s, got %s",
			result4.ToString(),
			Wikipedia.ToString(),
		)
	}
}
