package src

import (
	"log"

	"grest.dev/grest/db"

	"bitbucket.org/zahironline/zahirhrm-api/app"
	"bitbucket.org/zahironline/zahirhrm-api/src/quran"
)

func RunDBMigrator() {
	tx, err := db.DB("central")
	if err != nil {
		log.Fatal(err)
	} else {
		err = db.Migrate(tx, "central", app.SettingTable{})
	}
	if err != nil {
		log.Fatal(err)
	}
}

func SetDBMigration() {
	db.RegisterTable("company", quran.Surah{})
	db.RegisterTable("company", quran.Ayat{})

	db.RegisterTable("company", quran.PerawiHadits{})
	db.RegisterTable("company", quran.Hadits{})

	db.RegisterTable("company", quran.Dzikir{})
	db.RegisterTable("company", quran.Doa{})

}
