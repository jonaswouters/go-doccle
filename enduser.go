package main

import (
	"encoding/json"
	"log"
)

// EndUser object
type EndUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// GetUserInfo retrieves and returns an EndUser struct
func GetUserInfo(configuration Configuration) EndUser {
	url := "https://secure.doccle.be/doccle-euui/rest/v1/user"

	var resp = DoRequest(configuration, url)
	defer resp.Body.Close()

	var data struct {
		EndUser EndUser `json:"endUser"`
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&data); err != nil {
		log.Println(err)
	}

	return data.EndUser
}
