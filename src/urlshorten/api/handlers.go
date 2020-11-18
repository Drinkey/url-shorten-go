package api

import (
	"net/http"
    "log"
    "fmt"
    "io"
    "github.com/Drinkey/url-shorten-go/db"
)

func errorHandler(w http.ResponseWriter) {
    w.WriteHeader(http.StatusNotImplemented)
    err := fmt.Sprintf("%d %s", http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
    io.WriteString(w, err)
}

func DoRedirect(w http.ResponseWriter, req *http.Request){
    shortUrl := req.URL.Path[1:]
    if shortUrl == "" {
        errorHandler(w)
        return
    }
    log.Printf("handler: Got short url for redirect: %s", shortUrl)
    originUrl := db.GetOriginUrl(shortUrl)
    if originUrl == "" {
        log.Println("handlers: no record found, unable to serve")
        w.WriteHeader(http.StatusNotFound)
        return
    }
    log.Printf("handlers: Got origin url for redirect: %s", originUrl)
    http.Redirect(w, req, originUrl, http.StatusFound)
    log.Printf("handlers: Redirected to : %s", originUrl)
}

func RedirectHandler(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case http.MethodGet:
        log.Println("handlers: redirecting to origin url")
        DoRedirect(w, req)
    default:
        log.Println("handlers: got not supported method ", req.Method)
        errorHandler(w)
    }
    
}
func UrlsManagementHandler(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    switch req.Method {
    case http.MethodGet:
        log.Println("handlers: getting all url entries")
        getAllUrlRecords(w, req)
    case http.MethodPost:
        log.Println("handlers: creating new url entry")
        createUrlRecords(w, req)
    default:
        log.Println("handlers: got not supported method ", req.Method)
        errorHandler(w)
    }
}