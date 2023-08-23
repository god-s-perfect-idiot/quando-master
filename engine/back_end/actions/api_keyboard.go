package actions

import "quando/engine/structures"

type KeyboardClient struct {
}

func GetKeyboardActions() []structures.Method {
	keyboardClient := NewKeyboard()
	return []structures.Method{
		structures.Method{
			Identifier: "quando.control.key",
			Function:   keyboardClient.Key,
			Type:       "action",
		},
	}
}

func NewKeyboard() *KeyboardClient {
	return &KeyboardClient{}
}

func (k *KeyboardClient) Key(params map[string]interface{}) float64 {
	ch := params["ch"].(string)
	upDown := params["upDown"].(string)
	onOff := params["onOff"].(float64)
	callPipe := params["callPipe"].(*chan string)
	k.key(ch, upDown, onOff, callPipe)
	return 0.0
}

func (k *KeyboardClient) key(ch string, upDown string, onOff float64, callPipe *chan string) {
	println("pressKey", ch, upDown, onOff)
	*callPipe <- "here"
}
