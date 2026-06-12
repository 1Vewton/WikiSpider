// Manage the data sources
package data_source

// Data Source interface
type DataSource interface {
	GetText() (string, []string, error) // Get the text of certain page
	GetReferences() ([]string, error)   // Get the references in certain page
}
