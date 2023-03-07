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

type rawObject map[string]interface{}

func (d rawObject) d(k string) rawObject {
	return d[k].(map[string]interface{})
}

func (d rawObject) s(k string) string {
	return d[k].(string)
}

func (d rawObject) a(k string) []interface{} {
	return d[k].([]interface{})
}

type sObject struct {
	Type   string    `json:"-"`
	URL    string    `json:"url"`
	Body   rawObject `json:"body"`
	Id     string    `json:"referenceId"`
	Method string    `json:"method"`
	Childs []sObject `json:"-"`
}

type compRequest struct {
	GraphId     string    `json:"graphId"` // : "1",
	CompRequest []sObject `json:"compositeRequest"`
}

type MappingTable struct {
	Object     string   `json:""`
	References []string `json:""`
}
type Configuration struct {
	SFDCurl      string       `json:""`
	UserName     string       `json:""`
	Password     string       `json:""`
	ClientId     string       `json:""`
	ClientSecret string       `json:""`
	ExcludeList  []string     `json:""`
	IncludeList  []string     `json:""`
	LogLevel     string       `json:""`
	Mapping      MappingTable `json:""`
}
