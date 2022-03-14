package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
	"github.com/idzharbae/quickbid/src/app/config"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`
const configPath = "../files/config/quickbid.json"

func main() {
	flag.Usage = usage
	flag.Parse()

	var cfg config.Config
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}

	fileStr, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	err = json.Unmarshal(fileStr, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.DB.Address, cfg.DB.Port),
		Database: cfg.DB.DBName,
		User:     cfg.DB.Username,
		Password: cfg.DB.Password,
	})

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
	fmt.Print(usageText)
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
