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
