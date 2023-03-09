package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cloudflare/cfssl/log"
)

func getConfiguration() bool {

	// Open config
	log.Level = log.LevelDebug
	log.Debug("Before MistService")
	fmt.Println("Ooops")

	configFile, err := os.Open("config.json")

	// if  os.Open returns an error then handle it
	if err != nil {
		log.Error("Open Config File: ", err)
		os.Exit(1)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)
	json.Unmarshal(byteValue, &config)

	if ok, err := IsEmpty(config); err != nil {
		log.Error("Something is missing from configuration: ", err)
		return ok
	} else if ok {
		log.Error("Something is missing from configuration: ", ok)
		log.Debug("Config:\n", config)
		return ok
	}
	return true
}
