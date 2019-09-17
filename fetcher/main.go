package main

import (
	"io/ioutil"
	"log"

	"product-scraping/fetcher/app"

	yaml "gopkg.in/yaml.v2"
)

func GetConfigParams() (app.Config, error) {
	filename := "config/config.yml" // from /pkg/
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return app.Config{}, err
	}
	var config app.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return app.Config{}, err
	}
	return config, nil
}


func main() {
	appobj := &app.App{}
	conf, err := GetConfigParams()
	if err != nil {
		log.Fatal("Failed to fetch main config parameters", err)
	}
	appobj.Initialize(&conf)
	appobj.Run()
	close(appobj.Quit)
}
