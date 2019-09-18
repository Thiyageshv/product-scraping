package app

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	cas "product-scraping/lib_cassandra"
	util "product-scraping/lib_utilities"
)

type App struct {
	LogFile     *os.File
	Conf        *Config
	CasCursor   *cas.CasDb
	Router      *mux.Router
	Quit        chan bool
}


func (a *App) initRouter() {

	// endpoints
	r := mux.NewRouter()
	r.HandleFunc("/fetcher/api/v1/getInfo", a.getInfoEntry)
	r.HandleFunc("/fetcher/api/v1/getMetrics", a.getMetricsEntry)
	http.Handle("/fetcher/api/v1/", r)
	a.Router = r
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
	for err != nil {
		time.Sleep(time.Duration(a.Conf.RetryInterval) * time.Second)
		log.Println("Retrying connecting to Cassandra...")
		a.CasCursor, err = cas.InitializeCasaConn(config)
	}
}

func (a *App) Initialize(conf *Config) {
	a.Quit = make(chan bool)
	a.Conf = conf
	a.initializeLogging()
	a.initRouter()
	a.initializeCasCursor()
}

func (a *App) Run() {
	log.Println("serving...")
	http.ListenAndServe(":6000", nil)
}