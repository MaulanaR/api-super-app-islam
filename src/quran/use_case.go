package quran

import (
	"embed"
	_ "embed"
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/gorm/clause"
	"grest.dev/grest"
	"grest.dev/grest/convert"
	"grest.dev/grest/httpclient"

	"bitbucket.org/zahironline/zahirhrm-api/app"
)

//go:embed master_hadits/*.json
//go:embed master_doa/*.json
//go:embed master_dzikir/*.json
var content embed.FS

func Get(ctx *app.Ctx) error {
	tx, err := ctx.DB()
	if err != nil {
		return err
	}
	//GET DATA
	var data []Surah
	c := httpclient.New("GET", "https://api.npoint.io/99c279bb173a6e28359c/data")
	if app.PROXY_IS_DEBUG {
		c.IsDebug = true
	}
	c.Send()
	convert.NewJSON(c.BodyResponse).ToFlat().Unmarshal(&data)

	err = tx.Create(&data).Error
	if err != nil {
		return grest.NewError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func GetAyat(ctx *app.Ctx) error {
	tx, err := ctx.DB()
	if err != nil {
		return err
	}
	//GET DATA
	var data []Surah
	tx.Model(&Surah{}).Find(&data)
	for _, d := range data {
		var ayats []Ayat
		c := httpclient.New("GET", "https://api.npoint.io/99c279bb173a6e28359c/surat/"+strconv.Itoa(int(d.Nomor.Int64)))
		if app.PROXY_IS_DEBUG {
			c.IsDebug = true
		}
		c.Send()
		c.UnmarshalJson(&ayats)

		for _, ayat := range ayats {
			ayat.NomorSurah.Set(d.Nomor.Int64)
			ayat.IDAyat.Set(strconv.Itoa(int(d.Nomor.Int64)) + "_" + strconv.Itoa(int(ayat.Nomor.Int64)))
			err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&ayat).Error
			if err != nil {
				return grest.NewError(http.StatusInternalServerError, err.Error())
			}
		}
	}
	return nil
}

func SeedHadits(ctx *app.Ctx) error {
	tx, err := ctx.DB()
	if err != nil {
		return err
	}
	//GET DATA PERAWI
	var perawi []PerawiHadits
	byteValue, err := content.ReadFile("master_hadits/perawi-hadits.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &perawi)
	if err != nil {
		return err
	}

	for _, p := range perawi {
		err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&p).Error
		if err != nil {
			return grest.NewError(http.StatusInternalServerError, err.Error())
		}

		//GET DATA DETAIL HADITS
		var hadits []Hadits
		bVal, err := content.ReadFile("master_hadits/" + p.Kode.String + ".json")
		if err != nil {
			return err
		}
		err = json.Unmarshal(bVal, &hadits)
		if err != nil {
			return err
		}

		for _, hadit := range hadits {
			hadit.Kode.Set(p.Kode.String + "_" + strconv.Itoa(int(hadit.Nomor.Int64)))
			hadit.KodePerawi.Set(p.Kode.String)
			err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&hadit).Error
			if err != nil {
				return grest.NewError(http.StatusInternalServerError, err.Error())
			}
		}
	}
	return nil
}

func SeedDzikir(ctx *app.Ctx) error {
	tx, err := ctx.DB()
	if err != nil {
		return err
	}
	//GET DATA DZIKIR
	var dzikirs []Dzikir
	byteValue, err := content.ReadFile("master_dzikir/dzikir.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &dzikirs)
	if err != nil {
		return err
	}

	err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&dzikirs).Error
	if err != nil {
		return grest.NewError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func SeedDoa(ctx *app.Ctx) error {
	tx, err := ctx.DB()
	if err != nil {
		return err
	}
	//GET DATA DOA
	var doas []Doa
	byteValue, err := content.ReadFile("master_doa/doa.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &doas)
	if err != nil {
		return err
	}

	err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&doas).Error
	if err != nil {
		return grest.NewError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
