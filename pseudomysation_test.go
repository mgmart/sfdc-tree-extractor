package sfdcTreeExtractor

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cloudflare/cfssl/log"
)

func TestFakerCompany(t *testing.T) {

	got := fakeCompany()
	want := ""

	if got == want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPseudo(t *testing.T) {

	getConfiguration()
	newPseudo()
}

func getConfiguration() bool {

	// Open config
	configFile, err := os.Open("../sfdc-testdata-generator/config.json")

	// if  os.Open returns an error then handle it
	if err != nil {
		log.Error("Open Config File: ", err)
		os.Exit(1)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)
	err = json.Unmarshal(byteValue, &Config)
	if err != nil {
		log.Error("Unmarhalling Config: ", err)
		os.Exit(1)
	}
	return true
}
