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
	case Fandom:
		return "fandom"
	case Wikipedia:
		return "wikipedia"
	case NotSupported:
		return "not_supported"
	default:
		return "not_supported"
	}
}

// ToDataSourceType converts string to datasource type
func ToDataSourceType(rawDataSource string) DataSourceType {
	switch rawDataSource {
	case "fandom":
		return Fandom
	case "wikipedia":
		return Wikipedia
	case "not_supported":
		return NotSupported
	default:
		return NotSupported
	}
}
