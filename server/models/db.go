package models

import (
	"crypto/tls"
	"fmt"

	"github.com/asxcandrew/herd/server/config"
	"github.com/go-pg/pg"
)

var db *pg.DB

func DB() *pg.DB {
	if db == nil {
		options := pg.Options{
			Addr:     fmt.Sprintf("%s:%s", config.C.DBHost, config.C.DBPort),
			User:     config.C.DBUser,
			Password: config.C.DBPass,
			Database: config.C.DBName,
		}
		if config.C.TLS {
			options.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		}
		db = pg.Connect(&options)
	}
	return db
}
