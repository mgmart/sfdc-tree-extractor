package main

import (
	"errors"
	"reflect"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/cloudflare/cfssl/log"
)

func IsEmpty(object interface{}) (bool, error) {
	//First check normal definitions of empty
	if object == nil {
		return true, nil
	} else if object == "" {
		return true, nil
	} else if object == false {
		return true, nil
	}
	//Then see if it's a struct
	if reflect.ValueOf(object).Kind() == reflect.Struct {
		// and create an empty copy of the struct object to compare against
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			log.Debug("is Empty: ", object)
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, errors.New("Check not implementend for this struct")
}

func randomise() {
	// Create structs with random injected data
	log.Level = log.LevelDebug

	// log.Debug(gofakeit.Name())
	// log.Debug(gofakeit.Email())
	// log.Debug(gofakeit.StreetName(), " ", gofakeit.StreetNumber())
	// log.Debug(gofakeit.Zip())
	// log.Debug(gofakeit.City())
	// log.Debug(gofakeit.Contact())

	var cmp FakeCompany
	gofakeit.Struct(&cmp)
	log.Debug(cmp)
}
