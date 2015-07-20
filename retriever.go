package main

import (
	"flag"
	"fmt"
	"github.com/jonaswouters/go-doccle/doccle"
	"os/user"
	"strings"
)

func main() {
	// Parameters
	archivePtr := flag.Bool("archive", false, "Archive the downloaded documents")
	newPtr := flag.Bool("new", false, "Only download new documents")
	pathPtr := flag.String("path", "", "Path where the files should be downloaded")
	flag.Parse()

	// Get documents
	var configuration = doccle.GetConfiguration()
	var documentsResult doccle.DocumentsResult
	if *newPtr {
		documentsResult = doccle.GetNewDocuments(configuration)
	} else {
		documentsResult = doccle.GetDocuments(configuration)
	}

	// Home folder
	usr, _ := user.Current()
	dir := usr.HomeDir
	path := *pathPtr
	if path[:2] == "~/" {
		*pathPtr = strings.Replace(*pathPtr, "~", dir, 1)
	}

	for _, document := range documentsResult.Documents {
		var filename = strings.Join([]string{strings.Replace(document.Name, "/", "-", 999), ".pdf"}, "")
		filename = strings.Join([]string{document.Sender.Label, filename}, " - ")
		n, err := document.Download(configuration, *pathPtr, filename)

		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Printf("Downloaded %s (%d)\n", filename, n)

		if *archivePtr {
			fmt.Printf("Archiving document %s\n", filename)
			document.Archive(configuration)
		}
	}

}
