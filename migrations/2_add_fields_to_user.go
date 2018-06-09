package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("updating table users...")
		_, err := db.Exec(`ALTER TABLE users
									ADD COLUMN name varchar,
									ADD COLUMN bio varchar;
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("reverting table users...")
		_, err := db.Exec(`ALTER TABLE users
									DROP COLUMN name RESTRICT,
									DROP COLUMN bio RESTRICT;
		`)
		return err
	})
}
