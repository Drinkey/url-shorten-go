package main

import (
    "log"
    "net/http"
    _ "github.com/mattn/go-sqlite3"
    "github.com/Drinkey/url-shorten-go/api"
)

func main() {
    log.SetPrefix("main: ")
    log.Println("Starting HttpService")
    http.HandleFunc("/", api.RedirectHandler)
    http.HandleFunc("/urls", api.UrlsManagementHandler)
    go log.Fatal(http.ListenAndServe(":80", nil))
}