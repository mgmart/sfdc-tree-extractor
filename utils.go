package main

import (
	"errors"
	"reflect"
	"os"
	
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

func writeFile(data []byte, name string) {

	log.Info("Writing output to ", name)


	//create or open file
	f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// write to json file
	if _, err = f.Write(data); err != nil {
		panic(err)
	}
}
