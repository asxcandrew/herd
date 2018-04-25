package models

import (
	"crypto/tls"

	"github.com/asxcandrew/herd/server/config"
	"github.com/go-pg/pg"
)

var db *pg.DB

func DB() *pg.DB {
	if db == nil {
		options := pg.Options{
			Addr:     config.Root().DB.Addr,
			User:     config.Root().DB.User,
			Password: config.Root().DB.Pass,
			Database: config.Root().DB.Name,
		}
		if config.Root().DB.TLS {
			options.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		}
		db = pg.Connect(&options)
	}
	return db
}
