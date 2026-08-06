// go-url 是处理 URL 的小工具。
// 子命令：
//   encode  对 URL 做百分号编码（常用于拼接参数）
//   decode  把百分号编码还原
//   host    提取域名（去掉协议、路径、端口）
//   path    提取路径部分
//   query   解析查询参数，逐行打印 key=value
//   hash    生成一个短链风格的哈希（取 sha1 前 8 位）
//   info    一次性给出 编码/host/path/hash
package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"
)

func usage() {
	fmt.Print(`go-url URL 小工具

用法:
  go-url encode <url>      百分号编码
  go-url decode <url>      还原百分号编码
  go-url host <url>        提取域名
  go-url path <url>        提取路径
  go-url query <url>       解析查询参数
  go-url hash <url>        生成短链风格哈希(sha1 前8位)
  go-url info <url>        一次性给出 编码/host/path/hash
`)
}

// doEncode 百分号编码。
func doEncode(raw string) string { return url.QueryEscape(raw) }

// doDecode 还原百分号编码。
func doDecode(raw string) (string, error) { return url.QueryUnescape(raw) }

// doHost 提取域名。
func doHost(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	return u.Hostname(), nil
}

// doPath 提取路径部分。
func doPath(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	return u.Path, nil
}

// doQuery 解析查询参数，返回按 key 排序的 key=value 行。
func doQuery(raw string) ([]string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	keys := make([]string, 0, len(q))
	for k := range q {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var lines []string
	for _, k := range keys {
		lines = append(lines, fmt.Sprintf("%s=%s", k, q.Get(k)))
	}
	return lines, nil
}

// doHash 生成短链风格哈希（sha1 前 8 位）。
func doHash(raw string) string {
	sum := sha1.Sum([]byte(raw))
	return hex.EncodeToString(sum[:])[:8]
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
		fmt.Println(doEncode(raw))
	case "decode":
		decoded, err := doDecode(raw)
		if err != nil {
			fmt.Println("解码失败:", err)
			os.Exit(1)
		}
		fmt.Println(decoded)
	case "host":
		h, err := doHost(raw)
		if err != nil {
			fmt.Println("解析失败:", err)
			os.Exit(1)
		}
		fmt.Println(h)
	case "path":
		p, err := doPath(raw)
		if err != nil {
			fmt.Println("解析失败:", err)
			os.Exit(1)
		}
		fmt.Println(p)
	case "query":
		lines, err := doQuery(raw)
		if err != nil {
			fmt.Println("解析失败:", err)
			os.Exit(1)
		}
		for _, l := range lines {
			fmt.Println(l)
		}
	case "hash":
		fmt.Println(doHash(raw))
	case "info":
		h, _ := doHost(raw)
		p, _ := doPath(raw)
		fmt.Println("编码:", doEncode(raw))
		fmt.Println("域名:", h)
		fmt.Println("路径:", p)
		fmt.Println("哈希:", doHash(raw))
	case "-h", "--help", "help":
		usage()
	default:
		fmt.Println("未知命令:", cmd)
		usage()
	}
}
