package http_utils

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ResponseBody struct {
	Body       []byte
	StatusCode int
}

func SetHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

func MakeRequest(ctx context.Context, method, url string, headers map[string]string, body []byte) (*ResponseBody, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	SetHeaders(req, headers)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return &ResponseBody{
		Body:       respBody,
		StatusCode: resp.StatusCode,
	}, nil
}
