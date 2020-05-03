package configs

import "os"

const (
	ASSETS_DIR         = "/assets/"
	STATIC_DIR         = "/static/"
	PORT               = "4000"
	DEFAULT_SECRET_KEY = "THIS_IS_SUPER_SECRET_KEY"
)

func Port() string {
	val, ok := os.LookupEnv("PORT")
	if ok {
		return val
	}
	return PORT
}

func SecretKey() string {
	val, ok := os.LookupEnv("SECRET_KEY")

	if ok {
		return val
	} else {
		return DEFAULT_SECRET_KEY
	}
}
