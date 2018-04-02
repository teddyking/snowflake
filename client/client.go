package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/teddyking/snowflake"
)

type Client struct {
	hostAddress string
}

func New(hostAddress string) *Client {
	return &Client{
		hostAddress: hostAddress,
	}
}

func (c *Client) PostSuite(suite *snowflake.Suite) error {
	body, err := json.Marshal(suite)
	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf("%s/v1/suites", c.hostAddress), "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected HTTP response code: %d", resp.StatusCode)
	}

	return nil
}
