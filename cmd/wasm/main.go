package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	ctx := context.Background()
	infoURL, err := url.Parse("https://api.github.com/repos/golang/go/releases/latest")
	if err != nil {
		panic(err)
	}

	_, err = downloadFile(ctx, infoURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Downloaded file")
}

func downloadFile(ctx context.Context, url *url.URL) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to download info: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
