package config

// Settings struct
type Settings struct {
	wiki_url            string
	email               string
	project_name        string
	project_url         string
	request_packge_name string
	version             string
	maximum_retry       int
}

var settings Settings

// Get wiki url
func GetWikiUrl() string {
	// Return the wiki url
	return settings.wiki_url
}

// Get email
func GetEmail() string {
	// Return the email
	return settings.email
}

// Get project name
func GetProjectName() string {
	// Return the project name
	return settings.project_name
}

// Get project url
func GetProjectUrl() string {
	// Return the project url
	return settings.project_url
}

// Get request package name
func GetRequestPackageName() string {
	// Return the request package name
	return settings.request_packge_name
}

// Get version
func GetVersion() string {
	// Return the version
	return settings.version
}

// Get maximum retry
func GetMaximumRetry() int {
	// Return the maximum retry count
	return settings.maximum_retry
}
