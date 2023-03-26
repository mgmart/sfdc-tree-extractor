//
//  authenticate.go
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

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func getBearerToken() string {
	params := url.Values{
		"response_type": {"code"},
		"format":        {"json"},
		"grant_type":    {"password"},
		"client_id":     {Config.ClientId},
		"client_secret": {Config.ClientSecret},
		"username":      {Config.UserName},
		"password":      {Config.Password},
	}

	burl := Config.SFDCurl + "/services/oauth2/token" + "?" + params.Encode()

	req, _ := http.NewRequest("POST", burl, nil)
	body := getSalesForce(req)

	// log.Debug("Body: ", string(body))
	var token Token
	if err := json.Unmarshal(body, &token); err != nil {
		panic(err)
	}

	// log.Debug("Bearer: ", token.Bearer)

	return token.Bearer
}
