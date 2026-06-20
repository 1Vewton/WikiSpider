// Manage the data sources
package datasource

// Data Source interface
type DataSource interface {
	GetText() (string, []string, error) // Get the text of certain page
	GetReferences() ([]string, error)   // Get the references in certain page
	GetTargetURL() string               // Get the target url to request
}
