package tpcgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrJSONUnmarshal   = errors.New("json unmarshal")
	ErrNoKeyError      = errors.New("no api key found")
	ErrUnAuthorized    = errors.New("unauthorized")
	ErrNotFound        = errors.New("resource not found")
	ErrBadForm         = errors.New("invalid form")
	ErrAlreadyReported = errors.New("already reported")
)

func (s *Session) sendVATSIMRequest(method, url string, body interface{}) (*http.Response, error) {
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
	req.Header.Set("User-Agent", s.CoreAPISession.UserAgent)
	if s.CoreAPISession.ApiKey != "" {
		req.Header.Set(s.CoreAPISession.ApiKeyHeader, s.CoreAPISession.ApiKey)
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

func (s *Session) sendFCPRequest(method, url string, body interface{}) (*http.Response, error) {
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
	req.Header.Set("User-Agent", s.FCPSession.UserAgent)
	if s.FCPSession.ApiKey != "" {
		req.Header.Set(s.FCPSession.ApiKeyHeader, s.FCPSession.ApiKey)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case 200:
	case 201:
	case http.StatusUnauthorized:
		return nil, ErrUnAuthorized
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusUnprocessableEntity:
		return nil, ErrBadForm
	case http.StatusAlreadyReported:
		return nil, ErrAlreadyReported
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	return resp, nil
}

func (s *Session) sendCoreAPIRequest(method, url string, body interface{}) (*http.Response, error) {
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
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.CoreAPISession.UserAgent)
	if s.CoreAPISession.ApiKey != "" {
		req.Header.Set(s.CoreAPISession.ApiKeyHeader, s.CoreAPISession.ApiKey)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case 200:
	case 201:
	case http.StatusUnauthorized:
		return nil, ErrUnAuthorized
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusUnprocessableEntity:
		return nil, ErrBadForm
	case http.StatusAlreadyReported:
		return nil, ErrAlreadyReported
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	return resp, nil
}
