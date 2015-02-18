package main

import (
	"net/http"
)

// GetUserInfo retrieves and returns an EndUser struct
func DoRequest(configuration Configuration, url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)

	req.SetBasicAuth(configuration.Username, configuration.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp
}
