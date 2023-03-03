package main

const sfdcurl = "https://schnuddelhuddelde-dev-ed.my.salesforce.com"
const baseurl = sfdcurl + "/services/data/v57.0/sobjects/"
const queryurl = sfdcurl + "/services/data/v57.0/query?q="

// client credentials for SalesForce API
const clientId = <client-id>
const clientSecret = <client-secret>	
const userName = <userName>
const password = <password>

var bearer = "Bearer " + "<<access-token>>"

var excludeList = []string{"ActivityHistory", "AttachedContentDocument", "CombinedAttachment", "NoteAndAttachment", "OpenActivity", "ProcessInstanceHistory"}

var includeList = []string{"Case", "Contact", "Opportunity", "Account"}
