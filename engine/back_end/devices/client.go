package devices

import "quando/engine/structures"

var killChannel *chan map[string]interface{}

type Client struct {
	Methods []structures.Method
}

func (c *Client) AddCallbacks(methods []structures.Method) {
	for _, method := range methods {
		c.Methods = append(c.Methods, method)
	}
}

func (c *Client) GetCallbacks() []structures.Method {
	return c.Methods
}

func NewClient() *Client {
	client := &Client{}
	client.AddCallbacks(GetKeyboardCallbacks())
	client.AddCallbacks(GetTimerCallbacks())
	client.AddCallbacks(GetPickerCallbacks())
	client.AddCallbacks(GetMouseCallbacks())
	return client
}
