package utils


type UrlRecord struct {
    Id string `json:"id"`
    ShortUrl string `json:"short_url"`
    OriginUrl string `json:"origin_url"`
    OriginUrlSha256 string `json:"origin_url_sha"`
}

type UrlRecordQuery struct {
    ShortUrl string `json:"short_url"`
    OriginUrlSha256 string `json:"origin_url_sha"`
}