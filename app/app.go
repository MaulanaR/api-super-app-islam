package app

import (
	"time"

	"github.com/go-redis/redis/v8"
	"grest.dev/grest"
	"grest.dev/grest/cache"
	"grest.dev/grest/telegram"

	"bitbucket.org/zahironline/zahirhrm-api/app/translations"
)

const (
	//YYYY-MM-DD HH:ii:ss-Timezone
	FormatFullDateTimeZone = "2006-01-02 15:04:05-07"
	//YYYY-MM-DD
	FormatDateOnly = "2006-01-02"
)

func NewApp(cfg ...grest.App) *grest.App {
	configureEnv()
	configureTranslation()
	configureValidator()
	configureDB()
	configureCache()
	configureCrypto()
	// configureFileStore()
	configureTelegram()
	return grest.New(cfg...)
}

func configureTranslation() {
	translations.AddEnUS()
	translations.AddIdID()
}

func configureCache() {
	cache.Configure(cache.Config{
		RedisOptions: &redis.Options{
			Addr:     REDIS_HOST + ":" + REDIS_PORT,
			Username: REDIS_USERNAME,
			Password: REDIS_PASSWORD,
			DB:       REDIS_CACHE_DB,
		},
		DefaultExpiration: 24 * time.Hour,
	})
}

func configureTelegram() {
	telegram.Configure(TELEGRAM_ALERT_TOKEN, TELEGRAM_ALERT_CHANNEL_ID)
}
