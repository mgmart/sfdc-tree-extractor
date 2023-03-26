//
//  pseudomysation_test.go
//  sfdcTreeExtractor
//
//  Created by Mario Martelli on 05.03.23.
//  Copyright © 2023 Mario Martelli. All rights reserved.
//
//  This file is part of EverOrg.
//
//  sfdcTreeextractor is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  EverOrg is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with sfdcTreeExtractor. If not, see <http://www.gnu.org/licenses/>.

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
