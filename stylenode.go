package main

import (
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func stylenode(file string, n *html.Node) {
	var isStyle bool
	var href string
	var keptAttrs []html.Attribute

	for _, attr := range n.Attr {
		if attr.Key == "rel" && attr.Val == "stylesheet" {
			isStyle = true
			continue
		}

		if attr.Key == "href" {
			href = attr.Val
			continue
		}

		keptAttrs = append(keptAttrs, attr)
	}

	if !(isStyle && href != "") {
		return
	}

	var content string
	if strings.HasPrefix(href, "https://") || strings.HasPrefix(href, "http://") {
		content = getfromhttp(href)

	} else if strings.HasPrefix(href, "//") {
		content = getfromhttp("https:" + href)

	} else {
		content = getfromfile(filepath.Join(
			filepath.Dir(file),
			href,
		))
	}

	if content == "" {
		return
	}

	node := &html.Node{
		Type: html.ElementNode,
		Data: "style",
		Attr: keptAttrs,
	}

	node.AppendChild(&html.Node{
		Type: html.TextNode,
		Data: string(content),
	})

	n.Parent.InsertBefore(node, n)
	n.Parent.RemoveChild(n)
}
