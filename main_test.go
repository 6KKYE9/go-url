package main

import (
	"strings"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	in := "https://example.com/搜索?q=go 语言"
	enc := doEncode(in)
	if !strings.Contains(enc, "%E6%90%9C") {
		t.Errorf("中文未编码: %q", enc)
	}
	dec, err := doDecode(enc)
	if err != nil {
		t.Fatal(err)
	}
	if dec != in {
		t.Errorf("往返不一致: %q", dec)
	}
}

func TestHost(t *testing.T) {
	h, err := doHost("https://blog.example.com:8080/path?x=1")
	if err != nil {
		t.Fatal(err)
	}
	if h != "blog.example.com" {
		t.Errorf("host=%q want blog.example.com", h)
	}
}

func TestPath(t *testing.T) {
	p, err := doPath("https://example.com/a/b/c?x=1")
	if err != nil {
		t.Fatal(err)
	}
	if p != "/a/b/c" {
		t.Errorf("path=%q want /a/b/c", p)
	}
}

func TestQuery(t *testing.T) {
	lines, err := doQuery("https://example.com/?b=2&a=1&a=3")
	if err != nil {
		t.Fatal(err)
	}
	got := strings.Join(lines, "\n")
	// a 有多个值，Query.Get 取第一个；按字母序 a 在前
	if !strings.Contains(got, "a=1") || !strings.Contains(got, "b=2") {
		t.Errorf("query 解析异常: %q", got)
	}
}

func TestHash(t *testing.T) {
	h := doHash("https://example.com/abc")
	if len(h) != 8 {
		t.Errorf("hash 长度应为 8, got %d (%q)", len(h), h)
	}
	// 不同输入不同哈希
	if doHash("x") == doHash("y") {
		t.Error("不同输入哈希应不同")
	}
}
