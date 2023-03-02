package main

const sfdcurl = "https://schnuddelhuddelde-dev-ed.my.salesforce.com"
const baseurl = sfdcurl + "/services/data/v57.0/sobjects/"
const queryurl = sfdcurl + "/services/data/v57.0/query?q="

const bearer = "Bearer " + "<<access-token>>"

var excludeList = []string{"ActivityHistory", "AttachedContentDocument", "CombinedAttachment", "NoteAndAttachment", "OpenActivity", "ProcessInstanceHistory"}

var includeList = []string{"Case", "Contact", "Opportunity", "Account"}

var includeObjects = []ChildRelationship{
	{Name: "Account"},
	{Name: "Case"},
	{Name: "Contact"},
	{Name: "Opportunity"},
}
