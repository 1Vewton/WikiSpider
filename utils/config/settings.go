package config

// Settings struct
type Settings struct {
	wiki_url                 *string // url of the wiki
	email                    *string // email address
	project_name             *string // name of the project
	project_url              *string // url of the project
	request_packge_name      *string // name of the request package
	version                  *string // version of the project
	maximum_retry            *int    // maximum retry count
	embedding_model_name     *string // embedding model name
	embedding_model_base_url *string // embedding model base url
	embedding_model_api_key  *string // api key of the embedding model
}

var settings Settings

// Get wiki url
func GetWikiURL() string {
	SetEnvString(
		"WIKI_URL",
		"https://en.wikipedia.org/w/api.php?",
		&settings.wiki_url,
	)
	// Return the wiki url
	return *settings.wiki_url
}

// Get email
func GetEmail() string {
	SetEnvString(
		"Email",
		"1vewton4dev@gmail.com",
		&settings.email,
	)
	// Return the email
	return *settings.email
}

// Get project name
func GetProjectName() string {
	SetEnvString(
		"PROJECT_NAME",
		"WikiSpider",
		&settings.email,
	)
	// Return the project name
	return *settings.project_name
}

// Get project url
func GetProjectURL() string {
	SetEnvString(
		"PROJECT_URL",
		"https://github.com/1Vewton/WikiSpider",
		&settings.project_url,
	)
	// Return the project url
	return *settings.project_url
}

// Get request package name
func GetRequestPackageName() string {
	SetEnvString(
		"REQUEST_PACKGE_NAME",
		"Go-http-client/1.1",
		&settings.request_packge_name,
	)
	// Return the request package name
	return *settings.request_packge_name
}

// Get version
func GetVersion() string {
	SetEnvString(
		"VERSION",
		"0.1.0",
		&settings.version,
	)
	// Return the version
	return *settings.version
}

// Get maximum retry
func GetMaximumRetry() int {
	SetEnvInteger(
		"MAXIMUM_RETRY",
		3,
		&settings.maximum_retry,
	)
	// Return the maximum retry count
	return *settings.maximum_retry
}

// Get embedding model name
func GetEmbeddingModelName() string {
	SetEnvString(
		"EMBEDDING_MODEL_NAME",
		"",
		&settings.embedding_model_name,
	)
	return *settings.embedding_model_name
}

// Get embedding model api key
func GetEmbeddingAPIKey() string {
	SetEnvString(
		"EMBEDDING_MODEL_API_KEY",
		"",
		&settings.embedding_model_api_key,
	)
	return *settings.embedding_model_api_key
}

// Get embedding model base URL
func GetEmbeddingBaseURL() string {
	SetEnvString(
		"EMBEDDING_MODEL_BASE_URL",
		"",
		&settings.embedding_model_base_url,
	)
	return *settings.embedding_model_base_url
}
