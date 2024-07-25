package immudb

import (
	"context"
	"fmt"
	"os"
	"strconv"

	immudb "github.com/codenotary/immudb/pkg/client"
	"github.com/codenotary/immudb/pkg/api/schema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	client immudb.ImmuClient
}

func NewClient(url, username, password string) (*Client, error) {
	portStr := os.Getenv("IMMUDB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
    	return nil, fmt.Errorf("invalid IMMUDB_PORT value: %v", err)
	}
	opts := immudb.DefaultOptions().
		WithAddress(os.Getenv("IMMUDB_ADDRESS")).
		WithPort(port)

	client, err := immudb.NewImmuClient(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to create immudb client: %v", err)
	}

	ctx := context.Background()
	_, err = client.Login(ctx, []byte(username), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("failed to login: %v", err)
	}

	return &Client{client: client}, nil
}

func (c *Client) Set(ctx context.Context, key string, value []byte) error {
	_, err := c.client.Set(ctx, []byte(key), value)
	return err
}

func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	entry, err := c.client.Get(ctx, []byte(key))
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, fmt.Errorf("key not found: %s", key)
		}
		return nil, err
	}
	return entry.Value, nil
}

func (c *Client) Scan(ctx context.Context, prefix string) (map[string][]byte, error) {
	req := &schema.ScanRequest{
		Prefix: []byte(prefix),
	}
	entries, err := c.client.Scan(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform scan: %v", err)
	}

	result := make(map[string][]byte)
	for _, entry := range entries.Entries {
		result[string(entry.Key)] = entry.Value
	}

	return result, nil
}

func (c *Client) Close() error {
	return c.client.CloseSession(context.Background())
}