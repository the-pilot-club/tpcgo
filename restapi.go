package tpcgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func sendRequest(method, url string, body interface{}, ApiKeyHeader string, ApiKey string, UserAgent string) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if UserAgent != "" {
		req.Header.Set("User-Agent", UserAgent)
	}
	if ApiKeyHeader != "" || ApiKey != "" {
		req.Header.Set(ApiKeyHeader, ApiKey)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	return resp, nil
}
