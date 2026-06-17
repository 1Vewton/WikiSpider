package config

// Settings struct
type Settings struct {
	wiki_url            *string // url of the wiki
	email               *string // email address
	project_name        *string // name of the project
	project_url         *string // url of the project
	request_packge_name *string // name of the request package
	version             *string // version of the project
	maximum_retry       *int    // maximum retry count
}

var settings Settings

// Get wiki url
func GetWikiUrl() string {
	var raw_wiki_url string
	if settings.wiki_url == nil {
		raw_wiki_url = GetEnvString(
			"WIKI_URL",
			"https://en.wikipedia.org/w/api.php?",
		)
		settings.wiki_url = &raw_wiki_url
	}
	// Return the wiki url
	return *settings.wiki_url
}

// Get email
func GetEmail() string {
	// Return the email
	return *settings.email
}

// Get project name
func GetProjectName() string {
	// Return the project name
	return *settings.project_name
}

// Get project url
func GetProjectUrl() string {
	// Return the project url
	return *settings.project_url
}

// Get request package name
func GetRequestPackageName() string {
	// Return the request package name
	return *settings.request_packge_name
}

// Get version
func GetVersion() string {
	// Return the version
	return *settings.version
}

// Get maximum retry
func GetMaximumRetry() int {
	// Return the maximum retry count
	return *settings.maximum_retry
}
