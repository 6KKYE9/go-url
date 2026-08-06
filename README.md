# go-url

URL 的几个常见处理：编码、解码、取域名、算短链哈希。

```powershell
go run . encode "https://example.com/搜索?q=go 语言"
go run . decode "https%3A%2F%2Fexample.com%2F%25E6%2590%259C%25E7%25B4%25A2"
go run . host "https://blog.example.com:8080/path?x=1"
go run . path "https://example.com/a/b/c?x=1"
go run . query "https://example.com/?b=2&a=1"
go run . hash "https://example.com/abc"
go run . info "https://example.com/文章?title=hello"
```

子命令：

| 命令 | 说明 |
|------|------|
| `encode <url>` | 百分号编码（适合拼进查询参数） |
| `decode <url>` | 还原百分号编码 |
| `host <url>` | 只取域名，去掉协议/端口/路径 |
| `path <url>` | 提取路径部分 |
| `query <url>` | 解析查询参数，逐行打印 key=value（按 key 排序） |
| `hash <url>` | 用 sha1 取前 8 位当短链指纹 |
| `info <url>` | 一次给出编码后 / 域名 / 路径 / 哈希 |

实现上几个点：

- 编码解码直接用 `net/url` 的 `QueryEscape` / `QueryUnescape`，和浏览器地址栏行为一致。
- 域名提取走 `url.Parse` 的 `Hostname()`，比自己切字符串稳。
- `hash` 用 `crypto/sha1` 取前 8 位十六进制，仅作指纹用途，不保证不冲突。

写脚本顺手处理 URL 时能用上，比如从一堆链接里批量抽域名。
