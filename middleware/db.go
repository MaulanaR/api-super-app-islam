package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"bitbucket.org/zahironline/zahirhrm-api/app"
)

func SetDB(c *fiber.Ctx) error {
	ctx, ok := c.Locals(app.CtxKey).(*app.Ctx)
	if !ok {
		return fiber.NewError(http.StatusInternalServerError, "SetDB: ctx is not found")
	}
	ctx.TxBegin()
	err := c.Next()
	if err != nil || (c.Response().StatusCode() >= http.StatusBadRequest || c.Response().StatusCode() < http.StatusOK) {
		ctx.TxRollback()
	} else {
		ctx.TxCommit()
	}
	return nil
}
