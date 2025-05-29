package butils

func SafeStringFromMap(src map[string]interface{}, key string, dftValue string) string {
	if val, ok := src[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}

	return dftValue
}
