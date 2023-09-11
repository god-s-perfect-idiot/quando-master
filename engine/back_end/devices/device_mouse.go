package devices

import (
	"github.com/fstanis/screenresolution"
	hook "github.com/robotn/gohook"
	"quando/engine/structures"
	"runtime"
)

func GetMouseCallbacks() []structures.Method {
	return []structures.Method{
		{
			Identifier: "quando.mouse.handleX",
			Function:   MouseMoveX,
			Type:       "callback",
			Iterator:   true,
			Arbiter:    false,
		},
		{
			Identifier: "quando.mouse.handleY",
			Function:   MouseMoveY,
			Type:       "callback",
			Iterator:   true,
			Arbiter:    false,
		},
		{
			Identifier: "quando.mouse.handleClick",
			Function:   MouseClick,
			Type:       "callback",
			Iterator:   true,
			Arbiter:    false,
		},
	}
}

func MouseMoveX(params map[string]interface{}, runContext *structures.RunContext) {
	pressed := params["pressed"].(bool)
	inverted := params["inverted"].(bool)
	switch runtime.GOOS {
	case "linux":
		mouseMoveXLinux(pressed, inverted, runContext)
	case "windows":
		// TODO FIXME
	}
}

func MouseMoveY(params map[string]interface{}, runContext *structures.RunContext) {
	pressed := params["pressed"].(bool)
	inverted := params["inverted"].(bool)
	switch runtime.GOOS {
	case "linux":
		mouseMoveYLinux(pressed, inverted, runContext)
	case "windows":
		// TODO FIXME
	}
}

func MouseClick(params map[string]interface{}, runContext *structures.RunContext) {
	key := params["key"].(string)
	switch runtime.GOOS {
	case "linux":
		mouseClickLinux(key, runContext)
	case "windows":
		// TODO FIXME
	}
}

func getScreenSize() (int, int) {
	res := screenresolution.GetPrimary()
	return res.Height, res.Width
}

func mouseClickLinux(key string, runContext *structures.RunContext) {
	var keyButton uint16
	switch key {
	case "left":
		keyButton = 1
	case "right":
		keyButton = 3
	case "middle":
		keyButton = 2
	case "acc1":
		keyButton = 4
	case "acc2":
		keyButton = 5
	}
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		if e.Button == keyButton {
			for _, child := range runContext.CallNode.MainChildren {
				child.Method.CallFunc(runContext.Executable, child)
			}
			//hook.End()
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}

func mouseMoveYLinux(pressed bool, inverted bool, runContext *structures.RunContext) {
	yMax, _ := getScreenSize()
	var prevY *int16
	var diff float64
	hook.Register(hook.MouseMove, []string{}, func(e hook.Event) {
		var newY int16
		if inverted {
			newY = int16(yMax) - e.Y
		} else {
			newY = e.Y
		}
		if (pressed && e.Kind == 8) || (!pressed && e.Kind == 9) {
			if prevY != nil {
				yDiff := newY - *prevY
				//confidence
				if yDiff > 5 || yDiff < -5 {
					diff = float64(newY+1) / float64(yMax)
					runContext.Executable.Val = diff
					for _, child := range runContext.CallNode.MainChildren {
						child.Method.CallFunc(runContext.Executable, child)
					}
					//defer hook.End()
				}
			} else {
				prevY = &e.Y
			}
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}

func mouseMoveXLinux(pressed bool, inverted bool, runContext *structures.RunContext) {
	_, xMax := getScreenSize()
	var prevX *int16
	var diff float64
	hook.Register(hook.MouseMove, []string{}, func(e hook.Event) {
		var newX int16
		if inverted {
			newX = int16(xMax) - e.X
		} else {
			newX = e.X
		}
		if (pressed && e.Kind == 8) || (!pressed && e.Kind == 9) {
			if prevX != nil {
				xDiff := newX - *prevX
				//confidence
				if xDiff > 5 || xDiff < -5 {
					diff = float64(newX+1) / float64(xMax)
					runContext.Executable.Val = diff
					for _, child := range runContext.CallNode.MainChildren {
						child.Method.CallFunc(runContext.Executable, child)
					}
					//hook.End()
				}
			} else {
				prevX = &e.X
			}
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}
