package main

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"

	"golang.org/x/net/html"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: htmlbuild <file>")
		return
	}

	file := args[1]
	src, err := os.Open(file + ".html")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	defer src.Close()

	doc, err := html.Parse(src)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	traverse(file, doc)

	var buf bytes.Buffer
	if err := html.Render(&buf, doc); err != nil {
		slog.Error(err.Error())
		return
	}

	minifydst(file, &buf)
}
