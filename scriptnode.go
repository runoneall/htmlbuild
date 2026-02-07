package main

import (
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func scriptnode(file string, n *html.Node) {
	var src string

	for _, attr := range n.Attr {
		if attr.Key == "src" {
			src = attr.Val
			break
		}
	}

	if src == "" {
		return
	}

	var content string
	if strings.HasPrefix(src, "https://") || strings.HasPrefix(src, "http://") {
		content = getfromhttp(src)

	} else if strings.HasPrefix(src, "//") {
		content = getfromhttp("https:" + src)

	} else {
		content = getfromfile(filepath.Join(
			filepath.Dir(file),
			src,
		))
	}

	if content == "" {
		return
	}

	node := &html.Node{
		Type: html.ElementNode,
		Data: "script",
	}

	for _, attr := range n.Attr {
		if attr.Key != "src" {
			node.Attr = append(node.Attr, attr)
		}
	}

	node.AppendChild(&html.Node{
		Type: html.TextNode,
		Data: string(content),
	})

	n.Parent.InsertBefore(node, n)
	n.Parent.RemoveChild(n)
}
