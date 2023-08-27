package actions

import (
	"quando/engine/structures"
	"strconv"
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

func (k *KeyboardClient) Key(params map[string]interface{}) (float64, map[string]interface{}) {
	ch := params["ch"].(string)
	upDown := params["upDown"].(string)
	onOff := params["onOff"].(float64)
	callPipe := params["callPipe"].(*chan map[string]interface{})
	k.key(ch, upDown, onOff, callPipe)
	return 0.0, nil
}

func (k *KeyboardClient) key(ch string, upDown string, onOff float64, callPipe *chan map[string]interface{}) {
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
	println("press", press)
	body := []byte(`{"ch": "` + ch + `", "press": ` + strconv.FormatBool(press) + `}`)
	route := "control/key"
	payload := make(map[string]interface{})
	payload["route"] = route
	payload["body"] = body
	*callPipe <- payload
}
