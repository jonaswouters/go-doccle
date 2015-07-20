package doccle

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	// API_URL Url of the API
	API_URL = "https://secure.doccle.be/doccle-euui"
)

type (
	DocumentsResult struct {
		CurrentPage   int         `json:"currentPage",xml:"username"`
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
		Documents     []Document  `json:"documents"`
	}

	Document struct {
		Actions            []Action    `json:"actions"`
		Categories         []string    `json:"categories"`
		ContentURL         string      `json:"contentUrl"`
		CreationDate       string      `json:"creationDate"`
		Name               string      `json:"name"`
		Payment            interface{} `json:"payment"`
		PresentationType   string      `json:"presentationType"`
		Sender             Sender      `json:"sender"`
		SenderDocumentType string      `json:"senderDocumentType"`
		ShortName          interface{} `json:"shortName"`
		URI                string      `json:"uri"`
	}

	Action struct {
		Enabled bool   `json:"enabled"`
		ID      int    `json:"id"`
		Label   string `json:"label"`
		URL     string `json:"url"`
	}

	Sender struct {
		ID    string `json:"id"`
		Label string `json:"label"`
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

// DoRequest makes a request
func DoRequest(configuration Configuration, url string, method string) *http.Response {
	req, err := http.NewRequest(method, url, nil)

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
	url := strings.Join([]string{API_URL, "/rest/v1/documents?lang=en&order=DESC&page=1&pageSize=50&sort=date"}, "")

	var resp = DoRequest(configuration, url, "GET")
	defer resp.Body.Close()

	var data = DocumentsResult{}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&data); err != nil {
		log.Println(err)
	}

	return data
}

// GetNewDocuments retrieves and returns an DocumentsResult struct with new documents only
func GetNewDocuments(configuration Configuration) DocumentsResult {
	url := strings.Join([]string{API_URL, "/rest/v1/documents/new?lang=en&order=DESC&page=1&pageSize=50&sort=date"}, "")

	var resp = DoRequest(configuration, url, "GET")
	defer resp.Body.Close()

	var data = DocumentsResult{}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&data); err != nil {
		log.Println(err)
	}

	return data
}

// Download the document's file
func (document Document) Download(configuration Configuration, path string, filename string) (int64, error) {
	url := strings.Join([]string{API_URL, document.ContentURL}, "")
	var resp = DoRequest(configuration, url, "GET")
	defer resp.Body.Close()

	out, err := os.Create(strings.Join([]string{path, filename}, ""))
	defer out.Close()

	if err != nil {
		return 0, err
	}

	n, err := io.Copy(out, resp.Body)

	return n, err
}

// Archive the document
func (document Document) Archive(configuration Configuration) {
	for _, action := range document.Actions {
		if action.Label == "ARCHIVE" {
			url := strings.Join([]string{API_URL, action.URL}, "")
			var resp = DoRequest(configuration, url, "PUT")
			defer resp.Body.Close()
		}
	}
}
