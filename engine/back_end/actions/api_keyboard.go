package actions

import (
	"quando/engine/structures"
	"quando/internal/server/devices/keyboard"
)

type KeyboardClient struct {
}

func GetKeyboardActions() []structures.Method {
	keyboardClient := NewKeyboard()
	return []structures.Method{
		structures.Method{
			Identifier: "quando.control.key",
			Function:   keyboardClient.Key,
			Type:       "action",
			Iterator:   false,
			Arbiter:    false,
		},
	}
}

func NewKeyboard() *KeyboardClient {
	return &KeyboardClient{}
}

func (k *KeyboardClient) Key(params map[string]interface{}, _ *structures.RunContext) {
	ch := params["ch"].(string)
	upDown := params["upDown"].(string)
	onOff := params["onOff"].(float64)
	k.key(ch, upDown, onOff)
}

func (k *KeyboardClient) key(ch string, upDown string, onOff float64) {
	println("pressKey", ch, upDown, onOff)
	var press bool
	switch upDown {
	case "down":
		press = true
	case "either":
		if onOff < 0.5 {
			press = false
		} else {
			press = true
		}
	}
	keyboard.PressRelease(ch, press)
}
