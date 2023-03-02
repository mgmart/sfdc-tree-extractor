package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println("hello")
	// TODO: use composite API to limit number of API calls
	// TODO: use insertTree to get relationships automatically
	// TODO: build include-list support
	// TODO: create & use a configuration file

	//	opp := getOpportunity("0067Q000009kU6HQAU")

	//opp := getOpportunity("0067Q000009kU6NQAU")
	//_ = getAccount(opp.AccountId)
	//getContacts(opp.AccountId)

	// log.Println("Record value: å", dat["records"])
	//	log.Println("JSON map:\n", dat)

	getChilds("Account")
}

// getChilds gets all possible children of a given parent
// exlude and include lists are regarded
func getChilds(objc string) {
	// Getting all possible types which can be a child
	url := baseurl + objc + "/describe"
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)
	var obj ObjectDescription
	json.Unmarshal(body, &obj)

	var dat map[string]interface{}

	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	// log.Println(obj.Childs)
	// Query each type objects to get the childs of this type
	// TODO: Selection of include List could be less cryptic
	for _, v := range obj.Childs {
		if v.Name != "" {
			switch includeList == nil {
			case true:
				log.Println("Catching all objects ...")
				if !slices.Contains(excludeList, v.Obj) {
					getChildObjects(objc, v.Obj, v.Field)
				}
			case false:

				if slices.Contains(includeList, v.Obj) {
					log.Println("Catching selected type ...")
					getChildObjects(objc, v.Obj, v.Field)
				}
			}
		}
	}
}

// getChildOjects gets all objects of a given type which have a SalesForce
// objectId in a given field
func getChildObjects(objId string, tpe string, nme string) {
	url := queryurl + "SELECT+id+from+" + tpe + "+where+" + nme + "+=+'" + objId + "'"
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)

	// log.Println("Query: ", url)
	var res QueryResult
	json.Unmarshal(body, &res)
	//	log.Println(string(body))

	for _, v := range res.Records {
		// log.Println("URL of Object: ", v.Attributes.URL)
		url := sfdcurl + v.Attributes.URL
		req, _ := http.NewRequest("GET", url, nil)
		body := getSalesForce(req)

		var dat map[string]interface{}

		if err := json.Unmarshal(body, &dat); err != nil {
			panic(err)
		}
		// log.Println("Child: ", dat["Name"], " id: ", dat["Id"])
		// log.Println("Child: ", dat)
	}
}

func getContacts(acc string) {
	url := queryurl + "SELECT+id,+name+from+Contact+where+AccountId+=+'" + acc + "'"
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)
	var contacts QueryResult
	json.Unmarshal(body, &contacts)

	for _, v := range contacts.Records {
		log.Println("URL of Contact: ", v.Attributes.URL)
	}
}

func getContact(cntO string) map[string]interface{} {

	url := baseurl + "Contact/" + cntO
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)
	var acc Account
	json.Unmarshal(body, &acc)

	var dat map[string]interface{}

	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}
	return dat
}

func getOpportunity(oppO string) Opportunity {

	url := baseurl + "Opportunity/" + oppO
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)
	var opp Opportunity
	json.Unmarshal(body, &opp)

	log.Println("Id     : ", opp.Id)
	log.Println("Name   : ", opp.Name)
	log.Println("Type   : ", opp.Type)
	log.Println("Account: ", opp.AccountId)

	return opp
}

func getAccount(accO string) Account {

	url := baseurl + "Account/" + accO
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)
	var acc Account
	json.Unmarshal(body, &acc)

	log.Println("Id         : ", acc.Id)
	log.Println("Name       : ", acc.Name)
	log.Println("Type       : ", acc.Type)
	log.Println("Description: ", acc.Description)

	return acc
}

func getSalesForce(req *http.Request) []byte {

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	return body
	// log.Println(string([]byte(body)))

}
