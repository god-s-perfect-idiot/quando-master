package actions

import "quando/engine/structures"

type Client struct {
	Methods []structures.Method
}

func (c *Client) AddActions(methods []structures.Method) {
	for _, method := range methods {
		c.Methods = append(c.Methods, method)
	}
}

func (c *Client) GetActions() []structures.Method {
	return c.Methods
}

func NewClient() *Client {
	client := &Client{}
	client.AddActions(GetKeyboardActions())
	client.AddActions(GetLoggerActions())
	client.AddActions(GetMessengerActions())
	client.AddActions(GetMouseActions())
	return client
}
