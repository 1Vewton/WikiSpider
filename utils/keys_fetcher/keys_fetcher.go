package keys_fetcher

// Fetch keys
func FetchKeys(m map[string]any) []string {
	keys := make([]string, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}
