package sfdcTreeExtractor

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cloudflare/cfssl/log"
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

	log.Debug("Response: ", token)

	return token.Bearer
}
