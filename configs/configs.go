package configs

import "os"

const (
	ASSETS_DIR = "/assets/"
	STATIC_DIR = "/static/"
	PORT       = "4000"
)

func Port() string {
	val, ok := os.LookupEnv("PORT")
	if ok {
		return val
	}
	return PORT
}
