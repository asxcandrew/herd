package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"

	"github.com/asxcandrew/herd/server/config"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

const usageText = `This program runs command on the db. Supported commands are:
  - up - runs all available migrations.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.

Usage:
  go run *.go <command> [args]
`

func main() {
	var db *pg.DB

	flag.Usage = usage
	flag.Parse()

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

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Printf(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
