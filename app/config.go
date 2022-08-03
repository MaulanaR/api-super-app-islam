package app

import (
	"time"

	"github.com/joho/godotenv"
	"grest.dev/grest"
)

// default config
var (
	APP_VERSION    = "22.02.041727"
	APP_ENV        = "production"
	APP_PORT       = "4004"
	IS_MAIN_SERVER = false // set to true to run migration, seed and task scheduling
	IS_ON_PREMISE  = false

	JWT_KEY       = []byte("")
	CRYPTO_KEY    = ""
	CRYPTO_PREFIX = ""

	DB_DRIVER            = "postgres"
	DB_HOST              = "127.0.0.1"
	DB_PORT              = 5432
	DB_DATABASE          = "db_ilmu"
	DB_USERNAME          = "postgres"
	DB_PASSWORD          = ""
	DB_MAX_OPEN_CONNS    = 25
	DB_MAX_IDLE_CONNS    = 25
	DB_CONN_MAX_LIFETIME = time.Hour // on .env = "1h". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	DB_IS_DEBUG          = false

	REDIS_HOST      = "127.0.0.1"
	REDIS_PORT      = "6379"
	REDIS_CACHE_DB  = 3
	REDIS_REPORT_DB = 3
	REDIS_USERNAME  = ""
	REDIS_PASSWORD  = ""

	TELEGRAM_ALERT_TOKEN      = ""
	TELEGRAM_ALERT_CHANNEL_ID = ""

	PROXY_IS_DEBUG = false
)

func configureEnv() {
	godotenv.Load()

	grest.LoadEnv("APP_ENV", &APP_ENV)
	grest.LoadEnv("APP_PORT", &APP_PORT)
	grest.LoadEnv("IS_MAIN_SERVER", &IS_MAIN_SERVER)
	grest.LoadEnv("IS_ON_PREMISE", &IS_ON_PREMISE)

	grest.LoadEnv("JWT_KEY", &JWT_KEY)
	grest.LoadEnv("CRYPTO_KEY", &CRYPTO_KEY)
	grest.LoadEnv("CRYPTO_PREFIX", &CRYPTO_PREFIX)

	grest.LoadEnv("DB_DRIVER", &DB_DRIVER)
	grest.LoadEnv("DB_HOST", &DB_HOST)
	grest.LoadEnv("DB_PORT", &DB_PORT)
	grest.LoadEnv("DB_DATABASE", &DB_DATABASE)
	grest.LoadEnv("DB_USERNAME", &DB_USERNAME)
	grest.LoadEnv("DB_PASSWORD", &DB_PASSWORD)
	grest.LoadEnv("DB_MAX_OPEN_CONNS", &DB_MAX_OPEN_CONNS)
	grest.LoadEnv("DB_MAX_IDLE_CONNS", &DB_MAX_IDLE_CONNS)
	grest.LoadEnv("DB_CONN_MAX_LIFETIME", &DB_CONN_MAX_LIFETIME)
	grest.LoadEnv("DB_IS_DEBUG", &DB_IS_DEBUG)

	grest.LoadEnv("REDIS_HOST", &REDIS_HOST)
	grest.LoadEnv("REDIS_PORT", &REDIS_PORT)
	grest.LoadEnv("REDIS_CACHE_DB", &REDIS_CACHE_DB)
	grest.LoadEnv("REDIS_REPORT_DB", &REDIS_REPORT_DB)
	grest.LoadEnv("REDIS_USERNAME", &REDIS_USERNAME)
	grest.LoadEnv("REDIS_PASSWORD", &REDIS_PASSWORD)

	grest.LoadEnv("TELEGRAM_ALERT_TOKEN", &TELEGRAM_ALERT_TOKEN)
	grest.LoadEnv("TELEGRAM_ALERT_CHANNEL_ID", &TELEGRAM_ALERT_CHANNEL_ID)

	grest.LoadEnv("PROXY_IS_DEBUG", &PROXY_IS_DEBUG)

}
