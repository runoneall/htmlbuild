package main

import (
	"io"
	"log/slog"
	"os"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
)

func minifydst(file string, src io.Reader) {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("application/javascript", js.Minify)

	dst, err := os.Create(file + ".min.html")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	defer dst.Close()

	if err := m.Minify("text/html", dst, src); err != nil {
		slog.Error(err.Error())
		return
	}
}
