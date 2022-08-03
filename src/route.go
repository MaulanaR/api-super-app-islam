package src

import (
	"grest.dev/grest"

	"bitbucket.org/zahironline/zahirhrm-api/src/company"
	"bitbucket.org/zahironline/zahirhrm-api/src/quran"
)

func SetRoute(server *grest.App) {
	//========= PANGGIL SETELAH INSTALL BE ==============//
	server.AddRoute("/api/seed/surah", "GET", quran.GetHandler, nil)
	server.AddRoute("/api/seed/ayat", "GET", quran.AyatHandler, nil)
	server.AddRoute("/api/seed/hadits", "GET", quran.HaditsHandler, nil)
	server.AddRoute("/api/seed/dzikir", "GET", quran.DzikirHandler, nil)
	server.AddRoute("/api/seed/doa", "GET", quran.DoaHandler, nil)

	server.AddRoute("/api/migration", "GET", company.MigrationHandler, nil)
}
