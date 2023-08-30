package devices

import (
	"github.com/fstanis/screenresolution"
	hook "github.com/robotn/gohook"
	"math"
	"quando/engine/structures"
	"runtime"
)

func GetMouseCallbacks() []structures.Method {
	return []structures.Method{
		structures.Method{
			Identifier: "quando.mouse.handleX",
			Function:   MouseMoveX,
			Type:       "callback",
			Iterator:   true,
			Arbiter:    false,
		},
		structures.Method{
			Identifier: "quando.mouse.handleY",
			Function:   MouseMoveY,
			Type:       "callback",
			Iterator:   true,
			Arbiter:    false,
		},
		structures.Method{
			Identifier: "quando.mouse.handleClick",
			Function:   MouseClick,
			Type:       "callback",
			Iterator:   true,
			Arbiter:    false,
		},
	}
}

// TODO Very buggy. There's a race condition somewhere. Probably in the hook library.
func MouseMoveX(params map[string]interface{}) (float64, map[string]interface{}) {
	pressed := params["pressed"].(bool)
	inverted := params["inverted"].(bool)
	var val float64
	switch runtime.GOOS {
	case "linux":
		val = mouseMoveXLinux(pressed, inverted)
	case "windows":
		// TODO FIXME
	}
	println("val", val)
	return val, nil
}

func MouseMoveY(params map[string]interface{}) (float64, map[string]interface{}) {
	pressed := params["pressed"].(bool)
	inverted := params["inverted"].(bool)
	var val float64
	switch runtime.GOOS {
	case "linux":
		val = mouseMoveYLinux(pressed, inverted)
	case "windows":
		// TODO FIXME
	}
	println("val", val)
	return val, nil
}

func MouseClick(params map[string]interface{}) (float64, map[string]interface{}) {
	key := params["key"].(string)
	switch runtime.GOOS {
	case "linux":
		mouseClickLinux(key)
	case "windows":
		// TODO FIXME
	}
	return -1.0, nil
}

func getScreenSize() (int, int) {
	res := screenresolution.GetPrimary()
	return res.Height, res.Width
}

func mouseMoveLinux() {
	var prevX *int16
	var prevY *int16
	hook.Register(hook.MouseMove, []string{}, func(e hook.Event) {
		if prevX != nil && prevY != nil {
			xDiff := e.X - *prevX
			yDiff := e.Y - *prevY
			if xDiff > 10 || xDiff < -10 || yDiff > 10 || yDiff < -10 {
				if math.Abs(float64(xDiff)) > math.Abs(float64(yDiff)) {
					if xDiff > 0 {
						hook.End()
					} else {
						hook.End()
					}
					prevX = &e.X
				} else {
					if yDiff > 0 {
						hook.End()
					} else {
						hook.End()
					}
					prevY = &e.Y
				}
			}
		} else {
			prevY = &e.Y
			prevX = &e.X
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}

func mouseClickLinux(key string) {
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
			hook.End()
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}

func mouseMoveYLinux(pressed bool, inverted bool) float64 {
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
					defer hook.End()
				}
			} else {
				prevY = &e.Y
			}
		}
	})

	s := hook.Start()
	<-hook.Process(s)
	return diff
}

func mouseMoveXLinux(pressed bool, inverted bool) float64 {
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
					hook.End()
				}
			} else {
				prevX = &e.X
			}
		}
	})

	s := hook.Start()
	<-hook.Process(s)
	return diff
}
