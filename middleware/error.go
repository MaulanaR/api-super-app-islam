package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"grest.dev/grest"

	"bitbucket.org/zahironline/zahirhrm-api/app"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	return app.ErrorHandler(c, err)
}

func NotFoundHandler(c *fiber.Ctx) error {
	lang := c.Get("Content-Language")
	if lang == "" {
		lang = "en"
	}
	err := grest.Error{}
	err.Err.Code = http.StatusNotFound
	err.Err.Message = grest.Trans(lang, "404_not_found")
	return c.Status(err.Err.Code).JSON(err)
}

func Recover(c *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if err, ok = r.(error); !ok {
				err = fmt.Errorf("%v", r)
			}
			errMsg := err.Error()
			errTrace := app.GetTrace(0)
			ghe, HrmErr := c.Locals("grestHRMErrors").(*grest.Error)
			if HrmErr {
				errMsg = ghe.Error()
				errTrace = ghe.TraceSimple()
			}
			app.SendAlert(c, errMsg, errTrace)
		}
	}()
	return c.Next()
}
