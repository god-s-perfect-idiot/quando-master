package devices

import (
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
			Iterator:   true,
			Arbiter:    false,
		},
	}
}

func KeyPress(params map[string]interface{}, runContext *structures.RunContext) {
	key := params["key"].(string)
	ctrl := params["ctrl"].(bool)
	alt := params["alt"].(bool)
	switch runtime.GOOS {
	case "linux":
		keyPressLinux(key, ctrl, alt, runContext)
	case "windows":
		// TODO FIXME
	}
}

func keyPressLinux(key string, ctrl bool, alt bool, runContext *structures.RunContext) {
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
		//fmt.Println("received trigger")
		//hook.End()
		for _, child := range runContext.CallNode.MainChildren {
			child.Method.CallFunc(runContext.Executable, child)
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}
