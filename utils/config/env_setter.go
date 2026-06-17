package config

// Set the env variable for a string env var
func SetEnvString(
	key string,
	default_value string,
	key_pointer **string,
) {
	if *key_pointer == nil {
		var raw_value string
		raw_value = GetEnvString(
			key,
			default_value,
		)
		*key_pointer = &raw_value
	}
}

// Set the env variable for a int env var
func SetEnvInteger(
	key string,
	default_value int,
	key_pointer **int,
) {
	if *key_pointer == nil {
		var raw_value int
		raw_value = GetEnvInteger(
			key,
			default_value,
		)
		*key_pointer = &raw_value
	}
}
