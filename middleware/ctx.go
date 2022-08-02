package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"bitbucket.org/zahironline/zahirhrm-api/app"
)

func NewCtx(c *fiber.Ctx) error {
	lang := c.Get("Content-Language")
	if lang == "" {
		lang = "en"
	}
	companyID := ""
	dataID := ""
	endPoint := c.Path()
	path := strings.Split(endPoint, "/")
	pathLen := len(path)
	if pathLen > 4 {
		endPoint = path[4]
		if pathLen > 5 && path[4] == "companies" {
			companyID = path[5]
			if pathLen > 6 {
				endPoint = path[4] + "/:company_id/" + path[6]
			}
			if pathLen > 7 {
				dataID = path[7]
			}
		} else if pathLen == 6 {
			dataID = path[5]
		}
	}
	ctx := app.Ctx{
		Lang: lang,
		Action: app.Action{
			Method:    c.Method(),
			EndPoint:  endPoint,
			CompanyID: companyID,
			DataID:    dataID,
		},
	}
	c.Locals("ctx", &ctx)
	return c.Next()
}
