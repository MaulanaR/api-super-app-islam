package app

import (
	"gorm.io/gorm"
	"grest.dev/grest/db"
)

func (c *Ctx) DB(connName ...string) (*gorm.DB, error) {
	//if c.IsTest {
	//	return c.NewMockDB()
	//}
	if c.Company.Slug == "" || (len(connName) > 0 && connName[0] == "central") {
		// Control the transaction manually (set begin transaction, commit and rollback on middleware)
		if c.centralTx != nil {
			return c.centralTx, nil
		}
		// Autocommit if use goroutine, etc
		return db.DB("central")
	}

	if len(connName) > 0 {
		return db.DB(connName[0])
	}

	// Control the transaction manually (set begin transaction, commit and rollback on middleware)
	if c.companyTx != nil {
		return c.companyTx, nil
	}
	// Autocommit if use goroutine, etc
	return db.DB(c.Company.Slug)
}

func (c *Ctx) TxBegin() error {
	centralTx, err := db.DB("central")
	if err != nil {
		return err
	} else {
		c.centralTx = centralTx.Begin()
	}
	if c.Company.ConnName != "" {
		companyTx, err := db.DB(c.Company.ConnName)
		if err != nil {
			conf := db.Config{}
			conf.Driver = c.Company.DbDriver
			conf.Host = c.Company.DbHost
			conf.Port = int(c.Company.DbPort)
			conf.User, err = Decrypt(c.Company.DbUsername)
			if err != nil {
				return err
			}
			conf.Password, err = Decrypt(c.Company.DbPassword)
			if err != nil {
				return err
			}
			conf.DbName = c.Company.DbDatabase
			err := ConnectDB(c.Company.ConnName, conf)
			if err != nil {
				return err
			}
			companyTx, err = db.DB(c.Company.ConnName)
			if err != nil {
				return err
			}
		}
		c.companyTx = companyTx.Begin()
	}
	return nil
}

func (c *Ctx) TxCommit() {
	if c.centralTx != nil {
		c.centralTx.Commit()
	}
	if c.companyTx != nil {
		c.companyTx.Commit()
	}

	// reset to nil to use gorm autocommit if use goroutine, etc
	c.centralTx = nil
	c.companyTx = nil
}

func (c *Ctx) TxRollback() {
	if c.centralTx != nil {
		c.centralTx.Rollback()
	}
	if c.companyTx != nil {
		c.companyTx.Rollback()
	}

	// reset to nil to use gorm autocommit if use goroutine, etc
	c.centralTx = nil
	c.companyTx = nil
}
