package sfdcTreeExtractor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/cloudflare/cfssl/log"
	"golang.org/x/exp/slices"
)

//GetTree returns a whole tree based on configuration

func GetTree(obj string, tpe string) []sObject {

	log.Info("Authenticate with SF")
	Config.Bearer = "Bearer " + getBearerToken()

	log.Debug("pseudo: ", Config.NoPseudo)
	log.Debug("cleanup: ", Config.NoCleanUp)

	root := getRoot(obj, tpe)

	childs = append(childs, root)
	visited = append(visited, root.Id)
	getChilds(root)

	for _, v := range childs {

		if !Config.NoCleanUp {
			cleanUpObjects(&v)
		}
		if !Config.NoPseudo {
			pseudomyse(&v)
		}

		childs = reorderObjects(childs)
	}
	log.Info("Calls made: ", counter)
	return childs
}

// GetCompositeGraph creates a request
// body for a composite grapg request
func GetCompositeGraph(childs []sObject) []byte {
	//	account.Method = "POST"
	cRequest := compGraphRequest{GraphId: "1"}

	// -----
	// Create compound request element
	for _, v := range childs {
		// log.Info("Create compound: ", v.Type)
		for _, t := range Config.Mapping[v.Type] {
			if v.Body[t] != nil {
				v.Body[t] = "@{" + v.Body.s(t) + ".id}"
			}
		}
		v.Method = "POST"

		cRequest.CompRequest = append(cRequest.CompRequest, v)
	}

	//create or open file
	graph := compGraphs{Graphs: []compGraphRequest{cRequest}}
	data, _ := json.MarshalIndent(graph, "", " ")

	// --

	log.Info("Elements to write: ", len(cRequest.CompRequest))
	log.Info("Calls made: ", counter)
	return data
}

// reorders Object slice according to
// configured order
func reorderObjects(objs []sObject) []sObject {
	var orderedObjs []sObject

	for _, k := range Config.IncludeList {
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
	od := getObjectDescription(obj)
	for _, item := range od.Fields {
		if !item.Createable {
			delete(obj.Body, item.Name)
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

// getObjectDescription retrieves and stores a SalesForce
// Objectdescription for later retrieval
func getObjectDescription(obj *sObject) ObjectDescription {

	if des, ok := ods[obj.Type]; ok {
		return des
	}
	// Get the object description
	var od ObjectDescription
	if obj.Type == "" {
		return ObjectDescription{}
	}
	url := Config.SFDCurl + "/services/data/v57.0/sobjects/" + obj.Type + "/describe"
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)

	if err := json.Unmarshal(body, &od); err != nil {
		log.Warning("Unmarshalling of od failed: ", obj.Type)
		return ObjectDescription{}
	}
	ods[obj.Type] = od
	return od
}

// getChilds gets all possible children of a given parent
// exlude and include lists are regarded
func getChilds(objc sObject) {
	// log.Debug("Getting childs for ", objc.Type)
	objcd := getObjectDescription(&objc) // Query each type objects to get the childs of this type

	for _, v := range objcd.Childs {
		if v.Name != "" {
			if slices.Contains(Config.IncludeList, v.Obj) {
				for _, c := range getChildObjects(objc.Id, v.Obj, v.Field) {
					childs = append(childs, c)
					visited = append(visited, c.Id)
				}
			}
		}
	}

	// get all childs which are not referenced as childs
	for _, v := range objcd.Fields {
		if len(v.ReferenceTo) > 0 {
			// log.Debug("Getting other childs for ", objc.Type)
			for _, w := range v.ReferenceTo {
				obId := objc.Body.s(v.Name)
				if slices.Contains(Config.IncludeList, w) && obId != "" {
					if !slices.Contains(visited, obId) {
						o := getRoot(obId, w)
						o.URL = "/services/data/v57.0/sobjects/" + o.Type
						visited = append(visited, obId)
						childs = append(childs, o)
						getChilds(o)
					}
				}
			}
		}
	}
}

// getChildOjects gets all objects of a given type which have a SalesForce
// objectId in a given field
func getChildObjects(objId string, tpe string, nme string) []sObject {
	// Query for child records
	// log.Debug("getChildRecords")
	// logDebug()

	url := Config.SFDCurl + "/services/data/v57.0/query?q=" + "SELECT+id+from+" + tpe + "+where+" + nme + "+=+'" + objId + "'"
	req, _ := http.NewRequest("GET", url, nil)
	body := getSalesForce(req)
	var result []sObject
	var cR compRequest
	var res QueryResult
	var reqList compRequestList

	json.Unmarshal(body, &res)

	log.Debug("Query for: ", tpe, " results ", len(res.Records), " children")
	// log.Debug("Query URL: ", url)
	// Result is a list of records, now get each one of them
	for k, v := range res.Records {
		if slices.Contains(visited, v.Id) {
			continue
		}

		newCR := compositeRequest{
			Method:      "GET",
			URL:         v.Attributes.URL,
			ReferenceId: v.Id,
		}

		cR.CompositeRequest = append(cR.CompositeRequest, newCR)

		if len(cR.CompositeRequest) > 15 {
			reqList.List = append(reqList.List, cR)
			cR.CompositeRequest = []compositeRequest{}
		}
		if k == len(res.Records)-1 {
			reqList.List = append(reqList.List, cR)
		}
	}

	for k, v := range reqList.List {
		log.Debug("comp req: ", len(v.CompositeRequest), " - ", k)
		if len(v.CompositeRequest) > 0 {
			comp, _ := json.MarshalIndent(v, "", " ")
			url = Config.SFDCurl + "/services/data/v57.0/composite"
			req, _ = http.NewRequest("POST", url, bytes.NewReader(comp))
			req.Header.Add("Content-Type", "application/json")
			body = getSalesForce(req)

			var resp CompositeResponse
			// log.Debug("Response: ", string(body))
			if err := json.Unmarshal(body, &resp); err != nil {
				// TODO: Printout API Error
				log.Error("Create Composite Request: ", err)
				log.Debug(string(body))
			}
			// TODO: Check if response is valid
			for _, dat := range resp.Objects {
				dat.Type = dat.Body.d("attributes").s("type")
				// TODO: Potential bug
				dat.URL = "/services/data/v57.0/sobjects/" + dat.Body.d("attributes").s("type")
				// dat.URL = dat.Body.d("attributes").s("url")
				// dat.URL = "/services/data/v57.0/sobjects/" + dat.Type
				// log.Debug(dat.Id, ":", dat.Type)
				// log.Debug(dat.Body["FirstName"], " ", dat.Body["LastName"])
				result = append(result, dat)
				visited = append(visited, dat.Id)
				getChilds(dat)
			}
		}
	}
	return result
}

// getRoot returns a sObject
// structure given a sObjectId
func getRoot(obj, tpe string) sObject {

	url := Config.SFDCurl + "/services/data/v57.0/sobjects/" + tpe + "/" + obj
	req, _ := http.NewRequest("GET", url, nil)

	body := getSalesForce(req)
	var dat rawObject
	if err := json.Unmarshal(body, &dat); err != nil {
		log.Warning("Unmarshalling of root object failed: ", obj, tpe)
		return sObject{}
	}
	log.Debug("get Root")
	logDebug()
	return sObject{Type: dat.d("attributes").s("type"),
		Body: dat,
		Id:   dat.s("Id"),
		URL:  "/services/data/v57.0/sobjects/" + tpe,
	}
}

// getSalesForce returns the response for a given
// API request to SalesForce as a byte-array
func getSalesForce(req *http.Request) []byte {
	// log.Debug(req.URL)
	req.Header.Add("Authorization", Config.Bearer)
	calls = calls + 1
	client := &http.Client{}
	resp, err := client.Do(req)
	counter += 1
	if err != nil {
		log.Error("Error on response.\n[ERROR] -", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Error("Error while reading the response bytes:", err)
	}

	log.Debug("GetSalesForce", resp.Status)

	switch resp.StatusCode {
	case 400:
		log.Error("Error with Response: ", string(body))
		os.Exit(1)
	case 401:
		log.Error("Error with Response: ", resp.Status)
		os.Exit(1)
	}
	return body
	// log.Debug(string([]byte(body)))

}

func logDebug() {
	log.Debug("--------------------------------------------------------------")
	log.Debug("Length of Childs: ", len(childs))
	for _, v := range childs {
		log.Debug(v.Type)
	}
	log.Debug("--------------------------------------------------------------")
}
