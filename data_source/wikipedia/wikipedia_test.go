package wikipedia

import (
	"testing"

	"github.com/1Vewton/WikiSpider/data_source"
)

func TestInterface(t *testing.T) {
	var wikipedia_test interface{} = &WikipediaRequest{}
	if _, ok := wikipedia_test.(data_source.DataSource); !ok {
		t.Errorf("WikipediaRequest does not implement Request interface")
	}
}
