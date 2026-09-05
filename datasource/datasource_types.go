package datasource

// DataSourceType defines the type of data source
type DataSourceType string

const (
	// Fandom defines type for fandom
	Fandom DataSourceType = "fandom"
	// Wikipedia defines type for wikipedia
	Wikipedia DataSourceType = "wikipedia"
	// NotSupported converts type for not supported
	NotSupported DataSourceType = "not_supported"
)

// GetAllSuportedTypes gets all the type of data source that can be supported
func GetAllSuportedTypes() []DataSourceType {
	return []DataSourceType{
		Fandom,
		Wikipedia,
	}
}

// ToString converts DataSourceType to string
func (dataSourceType DataSourceType) ToString() string {
	switch dataSourceType {
	default:
		return "not_supported"
	}
}
