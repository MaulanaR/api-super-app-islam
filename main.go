package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"grest.dev/grest"
	"grest.dev/grest/db"

	"bitbucket.org/zahironline/zahirhrm-api/app"
	"bitbucket.org/zahironline/zahirhrm-api/middleware"
	"bitbucket.org/zahironline/zahirhrm-api/src"
)

const version = "22.07.291347"
// @title        Zahir HRM API
// @version      1.0
// @description  Zahir HRM API.
// @host         localhost:4004
// @schemes      http
// @BasePath     /
func main() {
	app.APP_VERSION = version

	server := app.NewApp(grest.App{
		ErrorHandler:    middleware.ErrorHandler,
		NotFoundHandler: middleware.NotFoundHandler,
		Config: fiber.Config{
			ReadBufferSize: 16384,
		},
	})
	server.AddRoute("/api/hrm/v1/ping", "GET", ping, nil)
	src.SetMiddleware(server)
	src.SetRoute(server)
	src.SetDBMigration()
	src.SetDBSeed()
	if app.APP_ENV == "local" || app.IS_MAIN_SERVER {
		src.RunDBMigrator()
		src.RunDBSeeder()
		src.RunScheduler()
	}
	if app.APP_ENV == "local" {
		server.AddSwagger("/docs/")
	}
	defer db.Close()
	err := server.Start(":" + app.APP_PORT)
	if err != nil {
		log.Fatal(err)
	}
}

func ping(c *fiber.Ctx) error {
	type res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Version string `json:"version"`
	}
	return c.JSON(res{200, "success", version})
}
