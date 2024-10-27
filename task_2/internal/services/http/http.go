package http

import (
	"net/http"
	"time"
)

type Checker interface {
	CheckURL(url string) (int, error)
}

type SimpleHTTPChecker struct{}

func NewSimpleHTTPChecker() *SimpleHTTPChecker {
	return &SimpleHTTPChecker{}
}

func (c *SimpleHTTPChecker) CheckURL(url string) (int, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
