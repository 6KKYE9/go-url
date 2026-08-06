// go-url 是一个纯标准库实现的 URL 小工具。
// 子命令：
//   encode   对 URL 做百分号编码（常用于拼接参数）
//   decode   把百分号编码还原
//   host     提取域名（去掉协议、路径、端口）
//   hash     生成一个短链风格的哈希（取 sha1 前 8 位）
//   info     一次性给出编码后、域名、哈希
// 依赖只有 net/url、crypto/sha1、strings、fmt、os。
package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func usage() {
	fmt.Print(`go-url URL 小工具

用法:
  go-url encode <url>      百分号编码
  go-url decode <url>      还原百分号编码
  go-url host <url>        提取域名
  go-url hash <url>        生成短链风格哈希(sha1 前8位)
  go-url info <url>        一次性给出 编码/host/hash
`)
}

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}
	cmd := os.Args[1]
	raw := strings.Join(os.Args[2:], "")
	switch cmd {
	case "encode":
		fmt.Println(url.QueryEscape(raw))
	case "decode":
		decoded, err := url.QueryUnescape(raw)
		if err != nil {
			fmt.Println("解码失败:", err)
			os.Exit(1)
		}
		fmt.Println(decoded)
	case "host":
		u, err := url.Parse(raw)
		if err != nil {
			fmt.Println("解析失败:", err)
			os.Exit(1)
		}
		fmt.Println(u.Hostname())
	case "hash":
		sum := sha1.Sum([]byte(raw))
		fmt.Println(hex.EncodeToString(sum[:])[:8])
	case "info":
		u, err := url.Parse(raw)
		host := ""
		if err == nil {
			host = u.Hostname()
		}
		sum := sha1.Sum([]byte(raw))
		fmt.Println("编码:", url.QueryEscape(raw))
		fmt.Println("域名:", host)
		fmt.Println("哈希:", hex.EncodeToString(sum[:])[:8])
	case "-h", "--help", "help":
		usage()
	default:
		fmt.Println("未知命令:", cmd)
		usage()
	}
}
