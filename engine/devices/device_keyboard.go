package devices

import (
	"quando/engine/structures"

	"github.com/eiannone/keyboard"
)

//const TypeKeyboard = "keyboard"

//type Keyboard struct {
//	//*BaseDevice
//	KeyPress events.Event
//}

func GetKeyboardCallbacks() []structures.Method {
	//keyboardFrame := &Keyboard{}
	//keyboardFrame.generateEvents()
	return []structures.Method{
		structures.Method{
			Identifier: "quando.key.handleKey",
			Function:   KeyPress,
			Type:       "callback",
		},
	}
}

//func (k *Keyboard) generateEvents() *[]events.Event {
//	var eventList []events.Event
//	keyPress := events.NewEvent("key_press", events.Nugget{})
//	keyPress.Function = func() {
//		if err := keyboard.Open(); err != nil {
//			println("Could not start keyboard listener")
//		}
//		defer func() {
//			_ = keyboard.Close()
//		}()
//		for {
//			char, key, err := keyboard.GetKey()
//			if err != nil {
//				panic(err)
//			}
//			println(char, key)
//			keyPress.Emit(events.Nugget{
//				Message: fmt.Sprintf("%q", char),
//				Payload: key,
//			})
//		}
//	}
//	k.KeyPress = *keyPress
//	eventList = append(eventList, *keyPress)
//	return &eventList
//}

//func NewKeyboard(b *BaseDevice, properties interface{}) Device {
//	device := &Keyboard{
//		BaseDevice: b,
//	}
//	device.SetProperties(properties)
//	device.SetEvents(*device.generateEvents())
//	return device
//}

func KeyPress(params map[string]interface{}) float64 {
	key := params["key"].(string)
	keyRune := []rune(key)[0]
	keyPress(keyRune)
	return 0.0
}

func keyPress(key rune) {
	if err := keyboard.Open(); err != nil {
		println("Could not start keyboard listener")
	}
	defer func() {
		_ = keyboard.Close()
	}()
	for {
		char, k, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		println(char, k)
		if char == key {
			break
		}
	}
}
