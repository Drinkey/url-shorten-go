package api

import (
    "log"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "urlshorten/db"
    "urlshorten/utils"
)

func JsonResponse(w http.ResponseWriter, code int, message interface{}) {
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(message)
}

func error(msg string) map[string]string {
    err := map[string]string {
        "error": "only JSON format in request body accepted",
    }
    return err
}

func getAllUrlRecords(w http.ResponseWriter, req *http.Request) {
    JsonResponse(w, http.StatusOK, db.GetAll(""))
}

func createUrlRecords(w http.ResponseWriter, req *http.Request) {
    if req.Header.Get("content-type") != "application/json" {
        JsonResponse(w, http.StatusBadRequest, 
            error("only JSON format in request body accepted"))
        return
    }

    var newUrl utils.UrlRecord
    body, _ := ioutil.ReadAll(req.Body)
    
    if err := json.Unmarshal(body, &newUrl); err != nil {
        JsonResponse(w, http.StatusBadRequest,
            error("read request body error"))
        return
    }
    
    newUrlSha := utils.Sha256Sum(newUrl.OriginUrl)
    urls := db.GetAll(newUrlSha)
    if len(urls) != 0 {
        log.Println("URL already exist ", newUrl.OriginUrl)
        JsonResponse(w, http.StatusOK, urls[0])
        return
    }

    r, success := db.Create(newUrl)
    if success != true {
        log.Println("Failed to create record for "+newUrl.OriginUrl)
        JsonResponse(w, http.StatusInternalServerError,
            error("failed to create record for "+newUrl.OriginUrl))
        return
    }
    JsonResponse(w, http.StatusCreated, r)
}
