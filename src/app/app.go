package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/idzharbae/quickbid/src/app/config"
)

// App layer glues all dependencies together

type QuickBid struct {
	Cfg      config.Config
	UseCases *UseCases
	Repos    *Repositories
}

func NewQuickBidApp(cfgPath string) *QuickBid {
	cfg := readConfig(cfgPath)

	repos := newRepositories(cfg)
	usecases := newUseCases(repos)

	return &QuickBid{
		Cfg:      cfg,
		Repos:    repos,
		UseCases: usecases,
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
