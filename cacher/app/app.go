package app

import (
	"fmt"
	"log"
	"os"
	cas "product-scraping/lib_cassandra"
	util "product-scraping/lib_utilities"
)

type App struct {
	LogFile     *os.File
	Conf        *Config
	CasCursor   *cas.CasDb
	Quit        chan bool
}

func (a *App) initializeLogging() {
	log.Println("Initializing Logging...")
	var err error
	a.LogFile, err = util.InitializeLogFile(a.Conf.LogFilePath)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		os.Exit(1)
	}
}

func (a *App) initializeCasCursor() {
	log.Println("Establishing Cassandra connection...")
	var err error
	config := cas.LocalCasaConfig()
	a.CasCursor, err = cas.InitializeCasaConn(config)
	if err != nil {
		log.Fatal("Failed to establish Cassandra connection")
	}
}

func (a *App) Initialize(conf *Config) {
	a.Quit = make(chan bool)
	a.Conf = conf
	a.initializeLogging()
	a.initializeCasCursor()

}

func (a *App) Run() {
	a.startScrapeJob()
}