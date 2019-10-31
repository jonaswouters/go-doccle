# README

## What is go-doccle

In the /pkg/doccle folder, there is a simple library that allows you to get all your documents, get all your new documents, download them and archive them.

The retriever is a simple implementation of the library that allows you to download and archive your documents.

## Simple usage example

*Make sure you create a configuration file `config.json` with your credentials for doccle.*

`./retriever -path="~/Desktop/doccledocuments/" -archive -new`

* The -new parameter only downloads the documents in the new queue
* The -archive parameter will archive the documents after downloading (Only useful in combination with the -new parameter)
* The -path parameter supports the home folder with ~, Make sure to always close with a /

## Contributions

This is only a simple downloader script. It won't get a lot of attention as long as it's working for me.
Any contributions are welcome.
