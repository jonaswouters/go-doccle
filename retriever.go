package main

import (
	"fmt"
	"strings"
	"github.com/jonaswouters/go-doccle/doccle"
)

func main() {
	var configuration = doccle.GetConfiguration()
	var documentsResult = doccle.GetDocuments(configuration)

	for _, document := range documentsResult.Documents {
		var filename = strings.Join([]string{strings.Replace(document.Name, "/", "-", 999), ".pdf"}, "")
		n, err := document.Download(configuration, filename)

		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Printf("%s (%d)\n", filename, n)
	}

}
