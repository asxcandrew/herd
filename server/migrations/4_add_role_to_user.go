package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("updating table users...")
		_, err := db.Exec(`ALTER TABLE users
									ADD COLUMN role varchar;
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("reverting table users...")
		_, err := db.Exec(`ALTER TABLE users
									DROP COLUMN role RESTRICT;
		`)
		return err
	})
}
