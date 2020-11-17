package main

import (
    "log"
    "net/http"
    _ "github.com/mattn/go-sqlite3"
    "urlshorten/api"
)

func main() {
    log.Println("Starting HttpService")
    http.HandleFunc("/", api.RedirectHandler)
    http.HandleFunc("/urls", api.UrlsManagementHandler)
    log.Fatal(http.ListenAndServe(":80", nil))
}