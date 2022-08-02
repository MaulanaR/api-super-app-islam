package quran

import (
	"github.com/gofiber/fiber/v2"

	"bitbucket.org/zahironline/zahirhrm-api/app"
)

func GetHandler(c *fiber.Ctx) error {
	ctx, err := app.NewCtx(c)
	if err != nil {
		return app.ErrorHandler(c, err)
	}
	err = Get(ctx)
	if err != nil {
		return app.ErrorHandler(c, err)
	}
	return c.JSON("OK")
}

func AyatHandler(c *fiber.Ctx) error {
	ctx, err := app.NewCtx(c)
	if err != nil {
		return app.ErrorHandler(c, err)
	}
	err = GetAyat(ctx)
	if err != nil {
		return app.ErrorHandler(c, err)
	}
	return c.JSON("OK")
}

func HaditsHandler(c *fiber.Ctx) error {
	ctx, err := app.NewCtx(c)
	if err != nil {
		return app.ErrorHandler(c, err)
	}
	err = SeedHadits(ctx)
	if err != nil {
		return app.ErrorHandler(c, err)
	}
	return c.JSON("OK")
}

func DzikirHandler(c *fiber.Ctx) error {
	ctx, err := app.NewCtx(c)
	if err != nil {
		return app.ErrorHandler(c, err)
	}
	err = SeedDzikir(ctx)
	if err != nil {
		return app.ErrorHandler(c, err)
	}
	return c.JSON("OK")
}

func DoaHandler(c *fiber.Ctx) error {
	ctx, err := app.NewCtx(c)
	if err != nil {
		return app.ErrorHandler(c, err)
	}
	err = SeedDoa(ctx)
	if err != nil {
		return app.ErrorHandler(c, err)
	}
	return c.JSON("OK")
}
