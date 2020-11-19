package db

import (
    "fmt"
    "database/sql"
    "log"
    "os"
    _ "github.com/mattn/go-sqlite3"
    "github.com/catinello/base62"
    "github.com/Drinkey/url-shorten-go/utils"
)

const DB_SCHEMA = `
CREATE TABLE urls (
    id integer not null primary key,
    short_url text not null,
    origin_url text,
    origin_url_sha text
    );
`

var DBPATH string = os.Getenv("URL_DB_PATH")

var conn *sql.DB

func init() {
    log.SetPrefix("db - init():")
    log.Println("initializing database")

    if DBPATH == "" {
        log.Panic("Specify the DB Path in environment variable URL_DB_PATH")
    }

    initDbRequired := true

    if utils.FileExist(DBPATH) {
        log.Println("database file already exist")
        initDbRequired = false
    }

    connDB, err := sql.Open("sqlite3", DBPATH)
    if err != nil {
        log.Panic(err)
    }
    conn = connDB

    if err = conn.Ping(); err != nil {
        log.Fatal("database unreachable:", err)
    }

    if initDbRequired {
        log.Println("first install, initializing database schema")
        _, err = conn.Exec(DB_SCHEMA)
        if err != nil {
            log.Fatalf("db %q: %s\n", err, DB_SCHEMA)
        }
        log.Printf("Database %s created", DBPATH)
    
        _, err = conn.Exec(`insert into urls(
                id, short_url, origin_url, origin_url_sha
            ) values (
                1, "abc", "test", "123123123123123123")`)
        if err != nil {
            log.Panic(err)
        }
    }
    
}

func generateIndex() int {
    log.SetPrefix("db - generateIndex():")
    rows, err := conn.Query("select id from urls order by id desc limit 1")
    if err != nil {
        log.Fatal(err)
    }
    var index int = 0
    for rows.Next() {
        err = rows.Scan(&index)
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Got latest id %d", index)
    }

    return index
}

func GetAll(urlSha string) []utils.UrlRecord {
    log.SetPrefix("db - GetAll():")
    var sql string
    if urlSha != "" {
        sql = "select * from urls where origin_url_sha='" + urlSha +"'"
    }else{
        sql = "select * from urls"
    }

    rows, err := conn.Query(sql)
    if err != nil {
        log.Panic(err)
    }
    var results []utils.UrlRecord
    for rows.Next() {
        var r utils.UrlRecord
        err = rows.Scan(&r.Id, &r.ShortUrl, &r.OriginUrl, &r.OriginUrlSha256)
        if err != nil {
            log.Fatal(err)
        }
        results = append(results, r)
    }

    return results
}

// Get DB record by specified property in UrlRecord struct
// If OriginSha specified, query by sha
// If ShortUrl specified, query by short URL
func Get(url utils.UrlRecordQuery) utils.UrlRecord {
    log.SetPrefix("db - Get():")
    var sql string
    if url.ShortUrl != "" {
        sql = "select * from urls where short_url='" + url.ShortUrl +"' order by id desc limit 1"
    }else if url.OriginUrlSha256 != "" {
        sql = "select * from urls where origin_url_sha='" + url.OriginUrlSha256 +"' order by id desc limit 1"
    }

    rows, err := conn.Query(sql)
    if err != nil {
        log.Fatal(err)
    }

    var r utils.UrlRecord
    for rows.Next() {
        err = rows.Scan(&r.Id, &r.ShortUrl, &r.OriginUrl, &r.OriginUrlSha256)
        if err != nil {
            log.Fatal(err)
            return utils.UrlRecord{}
        }
    }
    return r
}

func Create(url utils.UrlRecord) (utils.UrlRecord, bool) {
    log.SetPrefix("db - Create():")
    index := generateIndex()
    base62 := base62.Encode(index)
    sha := utils.Sha256Sum(url.OriginUrl)

    values := fmt.Sprintf("%d, '%s', '%s', '%s'", index+1, base62, url.OriginUrl, sha)
    sqlStmt := fmt.Sprintf("insert into urls values (%s)", values)
    log.Println(sqlStmt)

    log.Println("creating new record")
    _, err := conn.Exec(sqlStmt)
    if err != nil {
        log.Fatal(err)
        return utils.UrlRecord{}, false
    }
    query := utils.UrlRecordQuery{OriginUrlSha256: sha}
    return Get(query), true
}

func GetOriginUrl(shortUrl string) string {
    log.SetPrefix("db - GetOriginUrl():")
    r := Get(utils.UrlRecordQuery{ShortUrl: shortUrl})
    if r.OriginUrl == "" {
        log.Println("record not already exist")
        return ""
    }
    return r.OriginUrl
}