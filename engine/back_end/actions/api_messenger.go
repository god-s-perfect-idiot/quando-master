package actions

import (
	"quando/engine/structures"
	"quando/internal/server/socket"
)

type MessengerClient struct {
}

func GetMessengerActions() []structures.Method {
	messengerClient := NewMessenger()
	return []structures.Method{
		structures.Method{
			Identifier: "quando.message.send",
			Function:   messengerClient.Send,
			Type:       "action",
			Iterator:   false,
			Arbiter:    false,
		},
	}
}

func NewMessenger() *MessengerClient {
	return &MessengerClient{}
}

func (m *MessengerClient) Send(params map[string]interface{}) (float64, map[string]interface{}) {
	message := params["message"].(string)
	m.send(message)
	return 0.0, nil
}

func (m *MessengerClient) send(message string) {
	socket.Broadcast(message)
	println("send", message)
}
