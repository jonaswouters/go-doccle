package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"log"
)

type (
	DocumentsResult struct {
		CurrentPage   int         `json:"currentPage"`
		HasMore       bool        `json:"hasMore"`
		NextPage      int         `json:"nextPage"`
		PageSize      int         `json:"pageSize"`
		PreviousPage  interface{} `json:"previousPage"`
		Results       int         `json:"results"`
		SortField     string      `json:"sortField"`
		SortFieldType string      `json:"sortFieldType"`
		SortOrder     string      `json:"sortOrder"`
		TotalPages    int         `json:"totalPages"`
		TotalResults  int         `json:"totalResults"`
		Documents     []struct {
			Actions []struct {
				Enabled bool   `json:"enabled"`
				ID      int    `json:"id"`
				Label   string `json:"label"`
				URL     string `json:"url"`
			} `json:"actions"`
			Categories       []string    `json:"categories"`
			ContentURL       string      `json:"contentUrl"`
			CreationDate     string      `json:"creationDate"`
			Name             string      `json:"name"`
			Payment          interface{} `json:"payment"`
			PresentationType string      `json:"presentationType"`
			Sender           struct {
				ID    string `json:"id"`
				Label string `json:"label"`
			} `json:"sender"`
			SenderDocumentType string      `json:"senderDocumentType"`
			ShortName          interface{} `json:"shortName"`
			URI                string      `json:"uri"`
		} `json:"documents"`
	}

	EndUser struct {
		ID        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	Configuration struct {
		Username string
		Password string
	}
)

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

// GetConfiguration retrieves the config.json file and parses it
func GetConfiguration() Configuration {
	configuration := Configuration{}
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

// GetDocuments retrieves and returns an DocumentsResult struct
func GetDocuments(configuration Configuration) DocumentsResult {
	url := "https://secure.doccle.be/doccle-euui/rest/v1/documents?lang=en&order=DESC&page=1&pageSize=50&sort=date"

	var resp = DoRequest(configuration, url)
	defer resp.Body.Close()

	var data = DocumentsResult{}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&data); err != nil {
		log.Println(err)
	}

	return data
}
