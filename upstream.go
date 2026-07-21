package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Upstream(query []byte) ([]byte, error) {
	upstreamURL := "https://1.1.1.1/dns-query"

	req, err := http.NewRequest("POST", upstreamURL, bytes.NewReader(query)); if err != nil{
		return nil, fmt.Errorf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/dns-message")
	req.Header.Set("Accept", "application/dns-message")

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req); if err != nil{
		return nil, fmt.Errorf("Failed to get response: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("Status respond not OK: %d", res.StatusCode)
	}

	dnsRes, err := io.ReadAll(res.Body); if err != nil{
		return nil, fmt.Errorf("Failed to read response: %v", err)
	}

	return dnsRes, nil
}