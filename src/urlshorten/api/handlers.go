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
    log.SetPrefix("api - handlers.DoRedirect():")
    shortUrl := req.URL.Path[1:]
    if shortUrl == "" {
        errorHandler(w)
        return
    }
    log.Printf("Got short url for redirect: %s", shortUrl)
    originUrl := db.GetOriginUrl(shortUrl)
    if originUrl == "" {
        log.Println("no record found, unable to serve")
        w.WriteHeader(http.StatusNotFound)
        return
    }
    log.Printf("Got origin url for redirect: %s", originUrl)
    http.Redirect(w, req, originUrl, http.StatusFound)
    log.Printf("Redirected to : %s", originUrl)
}

func RedirectHandler(w http.ResponseWriter, req *http.Request) {
    log.SetPrefix("api - handlers.RedirectHandler():")
    switch req.Method {
    case http.MethodGet:
        log.Println("redirecting to origin url")
        DoRedirect(w, req)
    default:
        log.Println("got not supported method ", req.Method)
        errorHandler(w)
    }
    
}
func UrlsManagementHandler(w http.ResponseWriter, req *http.Request) {
    log.SetPrefix("api - handlers.UrlsManagementHandler():")
    w.Header().Set("Content-Type", "application/json")
    switch req.Method {
    case http.MethodGet:
        log.Println("getting all url entries")
        getAllUrlRecords(w, req)
    case http.MethodPost:
        log.Println("creating new url entry")
        createUrlRecords(w, req)
    default:
        log.Println("got not supported method ", req.Method)
        errorHandler(w)
    }
}