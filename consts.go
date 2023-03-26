//
//  consts.go
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

var Config = Configuration{}

// var bearer = "Bearer " + "<<access-token>>"
var calls = 0
var visited []string
var counter = 0
var childs []sObject

var ods = make(map[string]ObjectDescription)
