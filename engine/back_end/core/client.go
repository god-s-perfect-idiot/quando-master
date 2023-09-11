package core

import (
	"quando/engine/back_end/actions"
	"quando/engine/back_end/devices"
)

type Client struct {
	apiClient    *actions.Client
	deviceClient *devices.Client
}

func NewClient() *Client {
	apiClient := actions.NewClient()
	deviceClient := devices.NewClient()
	client := &Client{
		apiClient:    apiClient,
		deviceClient: deviceClient,
	}
	return client
}
