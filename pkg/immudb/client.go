package immudb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	http_utils "github.com/anuraj2023/bank-account-management-be/internal/utils"
)

type Client struct {
	apiKey string
	url    string
}

type AccountSearchResults struct {
	Page      int `json:"page"`
	PerPage   int `json:"perPage"`
	Revisions []struct {
		Document struct {
			ID         string `json:"_id"`
			VaultMD    struct {
				Creator string `json:"creator"`
				TS      int64  `json:"ts"`
			} `json:"_vault_md"`
			AccName    string  `json:"acc_name"`
			AccNumber  string  `json:"acc_number"`
			Address    string  `json:"address"`
			Amount     float64 `json:"amount"`
			IBAN       string  `json:"iban"`
			Type       string  `json:"type"`
		} `json:"document"`
		Revision      string `json:"revision"`
		TransactionID string `json:"transactionId"`
	} `json:"revisions"`
	SearchID string `json:"searchId"`
}

func NewClient(url, apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		url:    url,
	}
}

func (c *Client) getHeaders() map[string]string {
	return map[string]string{
		"accept":       "application/json",
		"Content-Type": "application/json",
		"X-API-Key":    c.apiKey,
	}
}

func (c *Client) AccNumberAlreadyExists(ctx context.Context, accNumber string) (bool, error) {
	fmt.Println("In AccNumberAlreadyExists")
	query := map[string]interface{}{
		"searchId": "1234567892",
		"query": map[string]interface{}{
			"expressions": []map[string]interface{}{
				{
					"fieldComparisons": []map[string]interface{}{
						{
							"field":    "acc_number",
							"operator": "EQ",
							"value":    accNumber,
						},
					},
				},
			},
		},
		"page":    1,
		"perPage": 1,
	}
	fmt.Println("query: ", query)
	jsonData, err := json.Marshal(query)
	if err != nil {
		fmt.Println("error in marshalling query: ", err)
		return false, fmt.Errorf("failed to marshal query: %v", err)
	}
	fmt.Println("jsonData: ", string(jsonData))

	url := fmt.Sprintf("%s/default/collection/default/documents/search", c.url)
	headers := c.getHeaders()
	respBody, err := http_utils.MakeRequest(ctx, http.MethodPost, url, headers, jsonData)
	if err != nil {
		fmt.Println("error in making request: ", err)
		return false, err
	}

	if respBody.StatusCode != http.StatusOK {
		fmt.Println("respBody.StatusCode is: ", respBody.StatusCode)
		fmt.Println("response body: ", string(respBody.Body))
		return false, fmt.Errorf("failed to search documents, status code: %d, response: %s", respBody.StatusCode, string(respBody.Body))
	}

	var result AccountSearchResults
	err = json.Unmarshal(respBody.Body, &result)
	if err != nil {
		fmt.Println("Failed to decode response: ", err)
		return false, fmt.Errorf("failed to decode response: %v", err)
	}

	return len(result.Revisions) > 0, nil
}



func (c *Client) Save(ctx context.Context, data []byte) error {
	
	var account map[string]interface{}
	if err := json.Unmarshal(data, &account); err != nil {
		return fmt.Errorf("failed to unmarshal data: %v", err)
	}

	// accNumber, ok := account["acc_number"].(string)
	// if !ok {
	// 	return fmt.Errorf("account number not found or invalid")
	// }

	// exists, err := c.AccNumberAlreadyExists(ctx, accNumber)
	// if err != nil {
	// 	return fmt.Errorf("failed to check if account number already exists: %v", err)
	// }

	// If account number exists already, no need to create a new record
	// if exists {
	// 	return fmt.Errorf("account number %s already exists", accNumber)
	// }

	url := fmt.Sprintf("%s/default/collection/default/document", c.url)
	headers := c.getHeaders()
	respBody, err := http_utils.MakeRequest(ctx, http.MethodPut, url, headers, data)
	if err != nil {
		return err
	}

	if respBody.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to set document, status code: %d, response: %s", respBody.StatusCode, string(respBody.Body))
	}
	return nil
}

func (c *Client) GetAll(ctx context.Context, page, perPage int) ([]map[string]interface{}, error) {
	query := map[string]int{
		"page":    page,
		"perPage": perPage,
	}
	jsonData, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query: %v", err)
	}

	url := fmt.Sprintf("%s/default/collection/default/documents/search", c.url)
	headers := c.getHeaders()
	respBody, err := http_utils.MakeRequest(ctx, http.MethodPost, url, headers, jsonData)
	if err != nil {
		return nil, err
	}

	if respBody.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get documents, status code: %d, response: %s", respBody.StatusCode, string(respBody.Body))
	}

	var result AccountSearchResults
	err = json.Unmarshal(respBody.Body, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	var documents []map[string]interface{}
	for _, revision := range result.Revisions {
		document := map[string]interface{}{
			"id":         revision.Document.ID,
			"creator":    revision.Document.VaultMD.Creator,
			"ts":         revision.Document.VaultMD.TS,
			"acc_name":   revision.Document.AccName,
			"acc_number": revision.Document.AccNumber,
			"address":    revision.Document.Address,
			"amount":     revision.Document.Amount,
			"iban":       revision.Document.IBAN,
			"type":       revision.Document.Type,
		}
		documents = append(documents, document)
	}

	return documents, nil
}
