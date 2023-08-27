package devices

import (
	"fmt"
	hook "github.com/robotn/gohook"
	"quando/engine/structures"
	"runtime"
)

func GetKeyboardCallbacks() []structures.Method {
	return []structures.Method{
		structures.Method{
			Identifier: "quando.key.handleKey",
			Function:   KeyPress,
			Type:       "callback",
		},
	}
}

func KeyPress(params map[string]interface{}) float64 {
	key := params["key"].(string)
	ctrl := params["ctrl"].(bool)
	alt := params["alt"].(bool)
	switch runtime.GOOS {
	case "linux":
		keyPressLinux(key, ctrl, alt)
	case "windows":
		// TODO FIXME
	}
	return 0.0
}

func keyPressLinux(key string, ctrl bool, alt bool) {
	keys := []string{
		key,
	}
	if ctrl {
		keys = append(keys, "ctrl")
	}
	if alt {
		keys = append(keys, "alt")
	}
	hook.Register(hook.KeyDown, keys, func(e hook.Event) {
		fmt.Println("received trigger")
		hook.End()
	})

	s := hook.Start()
	<-hook.Process(s)
}
