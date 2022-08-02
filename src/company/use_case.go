package company

import (
	"grest.dev/grest/db"

	"bitbucket.org/zahironline/zahirhrm-api/app"
)

func Migrate(ctx *app.Ctx, m ParamMigrate) error {
	tx, err := ctx.DB()
	if err != nil {
		return err
	}
	// migrate
	err = db.Migrate(tx, "company", app.CompanySettingTable{})
	if err != nil {
		return err
	}

	return nil
}
