package main

type Attributes struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type ObjectList struct {
	Id         string     `json:"Id"`
	Name       string     `json:"Name"`
	Attributes Attributes `json:"attributes"`
}

type QueryResult struct {
	Records []ObjectList `json:"records"`
}

type ChildRelationship struct {
	Name  string `json:"relationshipName"`
	Obj   string `json:"childSObject"`
	Field string `json:"field"`
}

type ObjectDescription struct {
	Childs []ChildRelationship `json:"childRelationships"`
}

type InvalidRequestResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type Token struct {
	Bearer    string `json:"access_token"`
	Type      string `json:"token_type"`
	URL       string `json:"instance_url"`
	Id        string `json:"id"`
	Issued    string `json:"issued_at"`
	Signature string `json:"signature"`
}
