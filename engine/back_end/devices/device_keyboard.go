package devices

import (
	hook "github.com/robotn/gohook"
	"quando/engine/structures"
	"runtime"
	"time"
)

func GetKeyboardCallbacks() []structures.Method {
	return []structures.Method{
		{
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

// TODO fix parallel execution

// This hijacks control from event loop
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
	lastTrigger := time.Now().UnixNano()
	hook.Register(hook.KeyDown, keys, func(e hook.Event) {
		newTime := time.Now().UnixNano()
		if newTime > lastTrigger+1000000000 {
			lastTrigger = newTime
			for _, child := range runContext.CallNode.MainChildren {
				child.Method.CallFunc(runContext.Executable, child)
			}
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}
