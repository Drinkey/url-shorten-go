# url-shorten-go
Just another practice project in Go

## Tasks

- [x] 实现一个发号器，可以利用数据库自增，比如 0，1，2，给每个 URL 发一个号
- [x] 实现 10-62 进制转换功能，将其作为记录的字段存入数据库
- [x] 实现 HTTP 服务器，接收 HTTP 请求，解析 URL 短地址参数，比如 request url 是 [http://z.cn/sBc](http://z.cn/sBc)，提取器短地址参数`sBc`
- [x] 实现短地址到长地址的查询，例如解析 `sBc`为对应发的号 1000，最后查询到编号为 1000 的记录中的 URL
- [x] 实现重定向功能，HTTP response 的 code为 302, location 为 URL 地址
- [ ] 实现转换请求的数据记录
- [ ] 实现 LRU，只支持最近 1 小时内的短地址查询

