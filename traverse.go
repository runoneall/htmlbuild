package main

import (
	"golang.org/x/net/html"
)

func traverse(file string, n *html.Node) {
	if n.Type == html.ElementNode {
		switch n.Data {

		case "link":
			stylenode(file, n)

		case "script":
			scriptnode(file, n)

		}
	}

	for c := n.FirstChild; c != nil; {
		next := c.NextSibling
		traverse(file, c)
		c = next
	}
}
