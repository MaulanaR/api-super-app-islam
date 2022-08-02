package company

import (
	"github.com/gofiber/fiber/v2"

	"bitbucket.org/zahironline/zahirhrm-api/app"
)

func GetListHandler(c *fiber.Ctx) error {
	return c.JSON(app.GeneralResponse{Code: 200, Message: "Success"})
}

func GetByIDHandler(c *fiber.Ctx) error {
	return c.JSON(app.GeneralResponse{Code: 200, Message: "Success"})
}

func CreateHandler(c *fiber.Ctx) error {
	return c.JSON(app.GeneralResponse{Code: 200, Message: "Success"})
}

func DeleteByIDHandler(c *fiber.Ctx) error {
	return c.JSON(app.GeneralResponse{Code: 200, Message: "Success"})
}

func MigrationHandler(c *fiber.Ctx) error {
	ctx, err := app.NewCtx(c)
	if err != nil {
		return app.ErrorHandler(c, err)
	}
	err = Migrate(ctx, ParamMigrate{})
	if err != nil {
		return app.ErrorHandler(c, err)
	}

	return c.JSON(app.GeneralResponse{Code: 200, Message: "Success"})
}

func GetFeatureHandler(c *fiber.Ctx) error {
	return c.JSON(app.GeneralResponse{Code: 200, Message: "Success"})
}

func GetSettingHandler(c *fiber.Ctx) error {
	return c.JSON(app.GeneralResponse{Code: 200, Message: "Success"})
}

func PartialyUpdateSettingHandler(c *fiber.Ctx) error {
	return c.JSON(app.GeneralResponse{Code: 200, Message: "Success"})
}
