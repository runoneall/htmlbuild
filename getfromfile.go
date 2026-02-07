package main

import (
	"log/slog"
	"os"
	"strings"
)

func getfromfile(path string) string {
	slog.Info("embed local file", "path", path)

	content, err := os.ReadFile(strings.Split(path, "?")[0])
	if err != nil {
		slog.Warn(err.Error())
		return ""
	}

	return string(content)
}
