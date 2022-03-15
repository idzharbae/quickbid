package app

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/idzharbae/quickbid/src/app/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

// App layer glues all dependencies together

type QuickBid struct {
	Cfg      config.Config
	UseCases *UseCases
	Repos    *Repositories
	Bridges  *Bridges
}

func NewQuickBidApp(cfgPath string) *QuickBid {
	cfg := readConfig(cfgPath)

	pgxPool, err := pgxpool.Connect(context.Background(), cfg.DB.ToConnectionString())
	if err != nil {
		log.Fatalf("[newRepositories]: %v", err)
	}

	bridges := newBridges(pgxPool)
	repos := newRepositories(pgxPool)
	usecases := newUseCases(repos, bridges)

	return &QuickBid{
		Cfg:      cfg,
		Repos:    repos,
		UseCases: usecases,
		Bridges:  bridges,
	}
}

func readConfig(cfgPath string) config.Config {
	var cfg config.Config
	file, err := os.Open(cfgPath)
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

	return cfg
}
