package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var configuration = GetConfiguration()
	var documentsResult = GetDocuments(configuration)

	for _, document := range documentsResult.Documents {
		url := strings.Join([]string{"https://secure.doccle.be/doccle-euui", document.ContentURL}, "")
		var resp = DoRequest(configuration, url)
		defer resp.Body.Close()

		var filename = strings.Join([]string{strings.Replace(document.Name, "/", "-", 999), ".pdf"}, "")

		out, err := os.Create(filename)
		defer out.Close()

		n, err := io.Copy(out, resp.Body)
		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Printf("%s (%d)\n", filename, n)
	}

}
