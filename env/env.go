package env

import "os"

//GetValue returns os env value by given key. when empty returns given default value
func GetValue(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}

	return defaultValue
}
