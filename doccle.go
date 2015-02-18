package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var configuration = GetConfiguration()
	var documentsResult = GetDocuments(configuration)

	for _, document := range documentsResult.Documents {
		fmt.Printf("%s%s\n", "https://secure.doccle.be/doccle-euui", document.ContentURL)
	}

}

// Configuration struct
type Configuration struct {
	Username string
	Password string
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
