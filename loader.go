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

	account.Method = "POST"
	cRequest := compRequest{GraphId: "1"}
	cRequest.CompRequest = append(cRequest.CompRequest, account)

	childs := getChilds(account.Type, account.Id)
	childs = reorderObjects(childs)

	// Create compound request element
	for _, v := range childs {
		log.Debug("Create compound: ", v.Type)
		for _, t := range idMapping[v.Type] {
			v.Body[t] = "@{" + v.Body.s(t) + ".id}"
		}
		cleanUpObjects(&v)
		v.Method = "POST"

		cRequest.CompRequest = append(cRequest.CompRequest, v)
	}
	graph := compGraphs{Graphs: []compRequest{cRequest}}
	data, _ := json.MarshalIndent(graph, "", " ")

	//create or open file
	f, err := os.OpenFile("test.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// write to json file
	if _, err = f.Write(data); err != nil {
		panic(err)
	}
}

// reorders Object slice according to
// configured order
func reorderObjects(objs []sObject) []sObject {
	var orderedObjs []sObject

	for _, k := range includeList {
		for _, o := range objs {
			if o.Type == k {
				orderedObjs = append(orderedObjs, o)
			}
		}
	}
	return orderedObjs
}

// removes all "not creatable" fields from given sObject
func cleanUpObjects(obj *sObject) {

	// Get the object description
	url := baseurl + obj.Type + "/describe"
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	for _, item := range dat["fields"].([]interface{}) {
		creatable := item.(map[string]interface{})["createable"].(bool)
		name := item.(map[string]interface{})["name"].(string)
		if !creatable {
			delete(obj.Body, name)
		}
	}
	delete(obj.Body, "attributes")

	// only non nil fields are needed in json file
	for k, v := range obj.Body {
		if v == nil {
			delete(obj.Body, k)
		}
	}
}

// getChilds gets all possible children of a given parent
// exlude and include lists are regarded
func getChilds(oType string, objc string) []sObject {
	// DONE: return the according result
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
	var childs []sObject
	for _, v := range obj.Childs {
		if v.Name != "" {
			switch includeList == nil {
			case true:
				if !slices.Contains(excludeList, v.Obj) {
					childs = append(childs, getChildObjects(objc, v.Obj, v.Field)...)
				}
			case false:
				if slices.Contains(includeList, v.Obj) {
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
