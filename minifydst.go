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

	m.Add("text/html", &html.Minifier{
		KeepComments:            false,
		KeepConditionalComments: false,
		KeepSpecialComments:     false,
		KeepDefaultAttrVals:     false,
		KeepDocumentTags:        false,
		KeepEndTags:             false,
		KeepQuotes:              false,
		KeepWhitespace:          false,
	})

	m.Add("text/css", &css.Minifier{
		Inline: true,
	})

	m.Add("application/javascript", &js.Minifier{
		KeepVarNames: false,
	})

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
