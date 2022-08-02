package src

import (
	"github.com/robfig/cron/v3"
)

func RunScheduler() {
	c := cron.New()

	// remove expired token every 00.05 WIB
	// c.AddFunc("CRON_TZ=Asia/Jakarta 5 0 * * *", app.RemoveExpiredToken)

	c.Start()
}
