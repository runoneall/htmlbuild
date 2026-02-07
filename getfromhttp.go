package main

import (
	"io"
	"log/slog"
	"net/http"
	"time"
)

func getfromhttp(url string) string {
	slog.Info("embed http file", "url", url)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return ""
	}

	if resp.StatusCode != http.StatusOK {
		return ""
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(content)
}
