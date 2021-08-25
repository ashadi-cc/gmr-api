package env

import "os"

func GetValue(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}

	return defaultValue
}
