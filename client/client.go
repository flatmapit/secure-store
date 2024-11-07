package client

import "crypto"

type Client struct {
}

func (client *Client) GetKey(url string) (crypto.PublicKey, error) {
	return nil, nil
}

func (client *Client) sendText(str string) error {
	return nil
}
