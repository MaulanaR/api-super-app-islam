package src

import (
	"grest.dev/grest"

	"bitbucket.org/zahironline/zahirhrm-api/middleware"
)

func SetMiddleware(server *grest.App) {
	server.AddMiddleware(middleware.NewCtx)
	server.AddMiddleware(middleware.Recover)
}
