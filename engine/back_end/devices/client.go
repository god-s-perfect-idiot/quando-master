package devices

import "quando/engine/structures"

type Client struct {
	//Devices []Device
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
	//keyboardDevice := NewDevice(TypeKeyboard, "keyboard", nil)
	client := &Client{
		//Devices: []Device{
		//	keyboardDevice,
		//},
	}
	client.AddCallbacks(GetKeyboardCallbacks())
	return client
}

//func (c *Client) Run() {
//	for _, device := range c.Devices {
//		device.ListenAll()
//	}
//}
