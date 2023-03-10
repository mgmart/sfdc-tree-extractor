package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/cloudflare/cfssl/log"
	"golang.org/x/exp/slices"
)

func main() {

	// TODO: process cmdline arguments
	// TODO: make LogLevel configurable
	log.Level = log.LevelDebug
	log.Info("sandbox-loader starting ...")

	objectPtr := flag.String("object", "0017Q00000NyD8jQAF", "Rootobject to retrieve")
	typePtr := flag.String("type", "Account", "The type of object")

	flag.Parse()

	log.Debug("object: ", *objectPtr)
	log.Debug("type: ", *typePtr)

	getConfiguration()

	log.Info("Authenticate with SF")
	bearer = "Bearer " + getBearerToken()

	// TODO: Get account from commandline
	// account := getAccount("0017Q00000NyD8jQAF")
	account := getRoot(*objectPtr, *typePtr)
	cleanUpObjects(&account)
	pseudomyse(&account)
	account.Method = "POST"
	cRequest := compRequest{GraphId: "1"}
	cRequest.CompRequest = append(cRequest.CompRequest, account)

	childs := getChilds(account.Type, account.Id)
	childs = reorderObjects(childs)

	// Create compound request element
	for _, v := range childs {
		log.Debug("Create compound: ", v.Type)
		for _, t := range config.Mapping[v.Type] {
			v.Body[t] = "@{" + v.Body.s(t) + ".id}"
		}
		cleanUpObjects(&v)
		pseudomyse(&v)
		v.Method = "POST"

		cRequest.CompRequest = append(cRequest.CompRequest, v)
	}

	//create or open file
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

	for _, k := range config.IncludeList {
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
	url := config.SFDCurl + "/services/data/v57.0/sobjects/" + obj.Type + "/describe"
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

func pseudomyse(obj *sObject) {

	switch obj.Type {
	case "Account":
		for key := range obj.Body {
			switch key {
			case "Name":
				obj.Body["Name"] = gofakeit.Company()
			case "Phone":
				obj.Body["Phone"] = gofakeit.PhoneFormatted()
			case "BillingCity":
				obj.Body["BillingCity"] = gofakeit.City()
			case "BillingState":
				obj.Body["BillingState"] = gofakeit.State()
			case "BillingStreet":
				obj.Body["BillingStreet"] = gofakeit.Street()
			case "Fax":
				obj.Body["Fax"] = gofakeit.PhoneFormatted()
			case "Website":
				obj.Body["Website"] = gofakeit.URL()
			case "AccountNumber":
				obj.Body["AccountNumber"] = strconv.Itoa(gofakeit.Number(1111111, 9999999))
			}
		}
	case "Contact":
		for key := range obj.Body {
			switch key {
			case "FirstName":
				obj.Body["FirstName"] = gofakeit.FirstName()
			case "Phone":
				obj.Body["Phone"] = gofakeit.PhoneFormatted()
			case "LastName":
				obj.Body["LastName"] = gofakeit.LastName()
			case "Fax":
				obj.Body["Fax"] = gofakeit.PhoneFormatted()
			case "Email":
				obj.Body["Email"] = gofakeit.Email()

			}
		}

	}
}

// getChilds gets all possible children of a given parent
// exlude and include lists are regarded
func getChilds(oType string, objc string) []sObject {
	// DONE: return the according result
	// Getting all possible types which can be a child
	url := config.SFDCurl + "/services/data/v57.0/sobjects/" + oType + "/describe"
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)
	var obj ObjectDescription
	json.Unmarshal(body, &obj)

	var dat map[string]interface{}

	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	// Query each type objects to get the childs of this type
	var childs []sObject
	for _, v := range obj.Childs {
		if oType == "CampaignMember" || oType == "Campaign" {
			// log.Debug("Child is :", v.Obj)
		}
		if v.Name != "" {
			if slices.Contains(config.IncludeList, v.Obj) {
				// log.Debug("Getting children: ", v.Obj, " - ", oType)
				childs = append(childs, getChildObjects(objc, v.Obj, v.Field)...)
			}
		}
	}

	// get all childs which are not referenced as childs
	// TODO: incorporate includelist
	for _, v := range obj.Fields {
		if len(v.ReferenceTo) > 0 {
			for _, w := range v.ReferenceTo {
				log.Debug("Getting children for: ", oType, ":", w, " - ", v.Name)
				// TODO: objc has to be changed to value of field v.Name
				childs = append(childs, getChildObjects(objc, oType, v.Name)...)
			}
		}

	}
	return childs
}

// getChildOjects gets all objects of a given type which have a SalesForce
// objectId in a given field
func getChildObjects(objId string, tpe string, nme string) []sObject {
	// Query for child records
	url := config.SFDCurl + "/services/data/v57.0/query?q=" + "SELECT+id+from+" + tpe + "+where+" + nme + "+=+'" + objId + "'"
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)
	var result []sObject

	if tpe == "CampaignMember" && nme == "ContactId" {
		log.Debug(url)
	}
	var res QueryResult
	json.Unmarshal(body, &res)

	// Result is a list of records, now get each one of them
	for _, v := range res.Records {
		url := config.SFDCurl + v.Attributes.URL
		req, _ := http.NewRequest("GET", url, nil)
		body := getSalesForce(req)

		var dat rawObject
		if err := json.Unmarshal(body, &dat); err != nil {
			panic(err)
		}

		cObj := sObject{
			Type: dat.d("attributes").s("type"),
			URL:  "/services/data/v57.0/sobjects/" + dat.d("attributes").s("type") + "/",
			Body: dat,
			Id:   dat.s("Id"),
		}
		result = append(result, cObj)
		if tpe == "CampaignMember" || tpe == "Campaign" {
			log.Debug("addRecurs: ", tpe, " - ", cObj.Id)
			getChilds(tpe, cObj.Id)
		}
	}
	return result
}

// getAccount returns a account as sObject
// structure given a sObjectId
func getAccount(accO string) sObject {

	url := config.SFDCurl + "/services/data/v57.0/sobjects/" + "Account/" + accO
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

// getAccount returns a account as sObject
// structure given a sObjectId
func getRoot(obj, tpe string) sObject {

	url := config.SFDCurl + "/services/data/v57.0/sobjects/" + tpe + "/" + obj
	req, _ := http.NewRequest("GET", url, nil)

	body := getSalesForce(req)

	var dat rawObject
	if err := json.Unmarshal(body, &dat); err != nil {
		log.Error("Unmarshalling of root object failed: ", obj, tpe)
		panic(err)
	}
	return sObject{Type: dat.d("attributes").s("type"),
		Body: dat,
		Id:   dat.s("Id"),
		URL:  "/services/data/v57.0/sobjects/" + tpe,
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
		log.Debug("Error with Response: ", resp.Status)
		os.Exit(1)
	}
	return body
	// log.Debug(string([]byte(body)))

}
