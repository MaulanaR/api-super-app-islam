package app

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"grest.dev/grest/db"
)

func configureDB() {
	c := db.Config{}
	c.Driver = DB_DRIVER
	c.Host = DB_HOST
	c.Port = DB_PORT
	c.User = DB_USERNAME
	c.Password = DB_PASSWORD
	c.TimeZone, _ = time.LoadLocation("UTC")
	c.DbName = DB_DATABASE
	err := ConnectDB("central", c)
	if err != nil {
		log.Fatal(err)
	}
	if !IS_ON_PREMISE {
		c.DbName = "zahironline_provider"
		err = ConnectDB("membership", c)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ConnectDB(connName string, c db.Config) error {
	dialector := postgres.Open(c.DSN())
	if c.Driver == "mysql" {
		dialector = mysql.Open(c.DSN())
	}

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return err
	}

	if DB_IS_DEBUG {
		gormDB = gormDB.Debug()
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(DB_MAX_OPEN_CONNS)
	sqlDB.SetMaxIdleConns(DB_MAX_OPEN_CONNS)
	sqlDB.SetConnMaxLifetime(DB_CONN_MAX_LIFETIME)

	db.Configure(connName, gormDB)
	return nil
}
