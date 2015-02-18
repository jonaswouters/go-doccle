package main

import (
	"encoding/json"
	"log"
	"net/http"
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

	req, err := http.NewRequest("GET", url, nil)

	req.SetBasicAuth(configuration.Username, configuration.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
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
