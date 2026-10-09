package jsontarget

import (
	"context"
	"encoding/json"
	"os"
	"slices"
	"sync"
	"time"
)

// JSONTarget defines the target for storing in json file
type JSONTarget struct {
	sync.RWMutex
	jsonFileName string
	Data         []any
}

// NewJSONTarget creates new json target
func NewJSONTarget(
	jsonFileName string,
) *JSONTarget {
	return &JSONTarget{
		jsonFileName: jsonFileName,
	}
}

// GetData gets data from target
func (target *JSONTarget) GetData() []any {
	target.RLock()
	defer target.RUnlock()
	result := slices.Clone(target.Data)
	return result
}

// Connect is just a simple function made it implements the target interface
func (target *JSONTarget) Connect(
	ctx context.Context,
	timeout time.Duration,
) error {
	target.Lock()
	defer target.Unlock()
	_, err := os.Stat(target.jsonFileName)
	if err != nil {
		if os.IsNotExist(err) {
			data, err := json.Marshal(
				target,
			)
			if err != nil {
				return err
			}
			err = os.WriteFile(
				target.jsonFileName,
				data,
				0644,
			)
			return err
		}
	}
	data, err := os.ReadFile(
		target.jsonFileName,
	)
	if err != nil {
		return err
	}
	err = json.Unmarshal(
		data,
		&target,
	)
	return err
}

// Close closes the connection
func (target *JSONTarget) Close(
	ctx context.Context,
) {
}

// NewDoc adds new doc
func (target *JSONTarget) NewDoc(
	ctx context.Context,
	datas []any,
) error {
	target.Lock()
	defer target.Unlock()
	target.Data = append(target.Data, datas...)
	jsonData, err := json.Marshal(
		target,
	)
	if err != nil {
		return err
	}
	err = os.WriteFile(
		target.jsonFileName,
		jsonData,
		0644,
	)
	return err
}
