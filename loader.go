package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/cloudflare/cfssl/log"
	"golang.org/x/exp/slices"
)

func main() {

	log.Level = log.LevelDebug
	log.Info("sandbox-loader starting ...")
	// TODO: use composite API to limit number of API calls
	// TODO: use insertTree to get relationships automatically
	// DONE: build include-list support
	// TODO: create & use a configuration file
	// DONE: implement proper authentication
	// DONE: link e.g. calculation not only to account but also to contract
	// DONE: create graph of objects
	// DONE: load graph into SB (can be done by Postman or so ...)

	log.Info("Authenticate with SF")
	bearer = "Bearer " + getBearerToken()

	account := getAccount("0017Q00000NyD8jQAF")
	cleanUpObjects(&account)
	// for k, v := range account.Body {
	// 	log.Debug("Key    : ", k)
	// 	log.Debug("Value  : ", v)
	// }
	// log.Debug("Account: ", account)
	childs := getChilds(account.Type, account.Id)

	// Cleanup IDs
	var idMapping map[string][]string
	idMapping = map[string][]string{
		"Contact":     {"AccountId"},
		"Case":        {"ContactId", "AccountId"},
		"Opportunity": {"AccountId"},
	}

	//create or open file
	f, err := os.OpenFile("test.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, v := range childs {
		// log.Debug("   child: " + v.Id + " - " + v.URL)
		account.Childs = append(account.Childs, v)
	}
	log.Debug("Length of Childs: ", len(account.Childs))

	cRequest := compRequest{GraphId: "1"}

	delete(account.Body, "attributes")
	account.Method = "POST"
	cRequest.CompRequest = append(cRequest.CompRequest, account)

	for _, v := range childs {
		delete(v.Body, "attributes")
		v.Method = "POST"
		for _, t := range idMapping[v.Type] {
			v.Body[t] = "@{" + v.Body.s(t) + ".id}"
		}
		cleanUpObjects(&v)
		// Create compound request element
		cRequest.CompRequest = append(cRequest.CompRequest, v)
	}
	log.Debug("Length of cRequest: ", len(cRequest.CompRequest))
	data, _ := json.MarshalIndent(cRequest, "", " ")
	if _, err = f.Write(data); err != nil {
		panic(err)
	}
}

// removes all "not creatable" fields from given sObject
func cleanUpObjects(obj *sObject) {

	// Get the object description
	url := baseurl + obj.Type + "/describe"
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)

	json.Unmarshal(body, &obj)

	// var dat rawObject
	//var dat map[string]map[string]map[string]interface{}
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	for _, item := range dat["fields"].([]interface{}) {
		creatable := item.(map[string]interface{})["createable"].(bool)
		name := item.(map[string]interface{})["name"].(string)
		if !creatable {
			log.Debug(name, " is not creatable")
			delete(obj.Body, name)
		}
	}

	// for k, item := range dat {
	// 	log.Debug("Boohr: ", k, item)
	// }
}

// getChilds gets all possible children of a given parent
// exlude and include lists are regarded
// DONE: return the according result
func getChilds(oType string, objc string) []sObject {
	// Getting all possible types which can be a child
	url := baseurl + oType + "/describe"
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)
	var obj ObjectDescription
	json.Unmarshal(body, &obj)

	var dat map[string]interface{}

	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	// Query each type objects to get the childs of this type
	// TODO: Selection of include List could be less cryptic

	log.Debug("Getting Childs for: ", oType)
	log.Debug("Calls so far: ", calls)
	var childs []sObject
	for _, v := range obj.Childs {
		if v.Name != "" {
			switch includeList == nil {
			case true:
				// log.Debug("Catching all objects ...")
				if !slices.Contains(excludeList, v.Obj) {
					childs = append(childs, getChildObjects(objc, v.Obj, v.Field)...)
				}
			case false:
				if slices.Contains(includeList, v.Obj) {
					// log.Debug("Catching selected type ...")
					childs = append(childs, getChildObjects(objc, v.Obj, v.Field)...)
				}
			}
		}
	}
	return childs
}

// getChildOjects gets all objects of a given type which have a SalesForce
// objectId in a given field
func getChildObjects(objId string, tpe string, nme string) []sObject {
	url := queryurl + "SELECT+id+from+" + tpe + "+where+" + nme + "+=+'" + objId + "'"
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)
	var result []sObject

	var res QueryResult
	json.Unmarshal(body, &res)

	for _, v := range res.Records {
		url := sfdcurl + v.Attributes.URL
		req, _ := http.NewRequest("GET", url, nil)
		body := getSalesForce(req)

		var dat rawObject
		if err := json.Unmarshal(body, &dat); err != nil {
			panic(err)
		}

		result = append(result, sObject{
			Type: dat.d("attributes").s("type"),
			URL:  "/services/data/v57.0/sobjects/" + dat.d("attributes").s("type") + "/",
			Body: dat,
			Id:   dat.s("Id"),
		})
	}
	return result
}

// getContacts returns all Contacts which have a relationship
// with a given account
// TODO: return the contacts in a propriet manner
func getContacts(acc string) {
	url := queryurl + "SELECT+id,+name+from+Contact+where+AccountId+=+'" + acc + "'"
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)
	var contacts QueryResult
	json.Unmarshal(body, &contacts)

	for _, v := range contacts.Records {
		log.Debug("URL of Contact: ", v.Attributes.URL)
	}
}

// getContact returns a contact by a given sObjectId
// as a map
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

// getOpportunity returns a Opportunity by a
// given sObjectId
// TODO: change the return from struct to map
func getOpportunity(oppO string) Opportunity {

	url := baseurl + "Opportunity/" + oppO
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)
	var opp Opportunity
	json.Unmarshal(body, &opp)

	log.Debug("Id      : ", opp.Id)
	log.Debug("Name    : ", opp.Name)
	log.Debug("Type    : ", opp.Type)
	log.Debug("Account : ", opp.AccountId)

	return opp
}

// getAccount returns a account as sObject
// structure given a sObjectId
func getAccount(accO string) sObject {

	url := baseurl + "Account/" + accO
	req, _ := http.NewRequest("GET", url, nil)

	body := getSalesForce(req)

	var dat rawObject
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}
	return sObject{Type: dat.d("attributes").s("type"),
		Body: dat,
		Id:   dat.s("Id"),
		URL:  "/services/data/v57.0/sobjects/Account",
	}
}

// getSalesForce returns the response for a given
// API request to SalesForce as a byte-array
func getSalesForce(req *http.Request) []byte {

	req.Header.Add("Authorization", bearer)
	calls = calls + 1
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Debug("Error on response.\n[ERROR] -", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Debug("Error while reading the response bytes:", err)
	}

	// log.Debug("GetSalesForce", resp.Status)

	switch resp.StatusCode {
	case 401:
		os.Exit(1)
	}
	return body
	// log.Debug(string([]byte(body)))

}
