//
//  object.go
//  sfdcTreeExtractor
//
//  Created by Mario Martelli on 25.02.23.
//  Copyright © 2023 Mario Martelli. All rights reserved.
//
//  This file is part of EverOrg.
//
//  sfdcTreeextractor is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  EverOrg is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with sfdcTreeExtractor. If not, see <http://www.gnu.org/licenses/>.

package sfdcTreeExtractor

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

type ObjectFields struct {
	Name        string   `json:"name"`
	ReferenceTo []string `json:"referenceTo"`
	Createable  bool     `json:"createable"`
}

type ObjectDescription struct {
	Name   string              `json:"label"`
	Childs []ChildRelationship `json:"childRelationships"`
	Fields []ObjectFields      `json:"fields"`
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
	switch d[k].(type) {
	case string:
		return d[k].(string)
	}
	return ""
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

type compGraphs struct {
	Graphs []compGraphRequest `json:"graphs"`
}
type compGraphRequest struct {
	GraphId     string    `json:"graphId"`
	CompRequest []sObject `json:"compositeRequest"`
}

type compositeRequest struct {
	Method      string `json:"method"`
	URL         string `json:"url"`
	ReferenceId string `json:"referenceId"`
}

type CompositeResponse struct {
	Objects []sObject `json:"compositeResponse"`
}

type compRequest struct {
	CompositeRequest []compositeRequest `json:"compositeRequest"`
}

type compRequestList struct {
	List []compRequest
}

type Configuration struct {
	SFDCurl      string                       `json:"sfdcurl"`
	UserName     string                       `json:"username"`
	Password     string                       `json:"password"`
	ClientId     string                       `json:"clientid"`
	ClientSecret string                       `json:"clientsecret"`
	Bearer       string                       `json:"accesstoken"`
	IncludeList  []string                     `json:"includelist"`
	LogLevel     string                       `json:"loglevel"`
	Mapping      map[string][]string          `json:"mapping"`
	Pseudo       map[string]map[string]string `json:"pseudo"`
	NoCleanUp    bool
	NoPseudo     bool
}
