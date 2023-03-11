package main

var config = Configuration{}
var bearer = "Bearer " + "<<access-token>>"
var calls = 0
var visited []string
var counter = 0
var childs []sObject

var ods = make(map[string]ObjectDescription)
