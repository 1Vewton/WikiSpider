package jsontarget

import (
	"fmt"
	"testing"
	"time"

	"github.com/1Vewton/WikiSpider/target"
)

// TestCRUD tests the CRUD of the json target
func TestCRUD(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	tmpDir := t.TempDir()
	filePath := fmt.Sprintf(
		"%s/%s",
		tmpDir,
		"test.json",
	)
	target := NewJSONTarget(filePath)
	err := target.Connect(
		ctx,
		5*time.Second,
	)
	if err != nil {
		t.Error(err)
	}
	data1 := map[string]string{
		"114514": "1919810",
	}
	data2 := map[string]string{
		"114514": "1919810",
	}
	err = target.NewDoc(
		ctx,
		[]any{
			data1,
			data2,
		},
	)
	if err != nil {
		t.Error(err)
	}
	data := target.GetData()
	if len(data) != 2 {
		t.Errorf(
			"expected length of data to be %d, got %d",
			2,
			len(data),
		)
	}
	newTarget := NewJSONTarget(
		filePath,
	)
	err = newTarget.Connect(
		ctx,
		5*time.Second,
	)
	if err != nil {
		t.Error(err)
	}
	data3 := newTarget.GetData()
	if len(data3) != 2 {
		t.Errorf(
			"expected length of data to be %d, got %d",
			2,
			len(data3),
		)
	}
}

// TestInterface tests whether the milvustarget can fit into milvus target
func TestInterface(t *testing.T) {
	t.Parallel()
	var testTarget interface{} = &JSONTarget{}
	_, ok := testTarget.(target.Target)
	if !ok {
		t.Error("milvus target does not implements methods of target")
	}
}
