// config.go

package config

import "os"

func Get(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v

	}

	return fallback

}
