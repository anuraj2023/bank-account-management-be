package immudb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	apiKey string
	url    string
}

func NewClient(url string, apiKey string) (*Client, error) {
	return &Client{
		apiKey: apiKey,
		url:    url,
	}, nil
}

func (c *Client) Save(ctx context.Context, data []byte) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf("%s/default/collection/default/document", c.url), bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to set document, status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

func (c *Client) GetAll(ctx context.Context, page, perPage int) ([]map[string]interface{}, error) {
	fmt.Println("Inside ListDocuments")
	query := map[string]int{
		"page":    page,
		"perPage": perPage,
	}
	jsonData, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query: %v", err)
	}

	fmt.Printf("Query JSON: %s\n", string(jsonData))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/default/collection/default/documents/search", c.url), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", c.apiKey)

	fmt.Printf("Request URL: %s\n", req.URL.String())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Response Status: %s\n", resp.Status)

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get documents, status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	var result struct {
		Documents []map[string]interface{} `json:"documents"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	fmt.Printf("Fetched Documents: %v\n", result.Documents)

	return result.Documents, nil
}
