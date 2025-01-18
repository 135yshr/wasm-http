package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"honnef.co/go/js/dom/v2"
)

func main() {
	ctx := context.Background()
	u, err := url.Parse("/data/dummy.txt")
	if err != nil {
		panic(err)
	}

	body, err := downloadFile(ctx, u)
	if err != nil {
		panic(err)
	}

	div := dom.GetWindow().Document().CreateElement("div")
	div.SetInnerHTML(string(body))

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
