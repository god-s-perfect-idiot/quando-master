package actions

import (
	"quando/engine/structures"
	"quando/internal/server/devices/mouse"
)

type MouseClient struct{}

func GetMouseActions() []structures.Method {
	mouseClient := NewMouse()
	return []structures.Method{
		{
			Identifier: "quando.control.mouseX",
			Function:   mouseClient.MouseX,
			Type:       "action",
			Iterator:   false,
			Arbiter:    false,
		},
		{
			Identifier: "quando.control.mouseY",
			Function:   mouseClient.MouseY,
			Type:       "action",
			Iterator:   false,
			Arbiter:    false,
		},
		{
			Identifier: "quando.control.mouse",
			Function:   mouseClient.MouseXY,
			Type:       "action",
			Iterator:   false,
			Arbiter:    false,
		},
	}
}

func NewMouse() *MouseClient {
	return &MouseClient{}
}

func (m *MouseClient) MouseX(params map[string]interface{}, _ *structures.RunContext) {
	x := params["x"].(float64)
	m.mouseX(x)
}

func (m *MouseClient) mouseX(x float64) {
	xBalanced := float32(x)
	y := float32(0)
	mouse.MovePress(mouse.MouseJSON{
		X:      &xBalanced,
		Y:      &y,
		Left:   "up",
		Middle: "up",
		Right:  "up",
	})
}

func (m *MouseClient) MouseY(params map[string]interface{}, _ *structures.RunContext) {
	y := params["y"].(float64)
	m.mouseY(y)
}

func (m *MouseClient) mouseY(y float64) {
	x := float32(0)
	yBalanced := float32(y)
	mouse.MovePress(mouse.MouseJSON{
		X:      &x,
		Y:      &yBalanced,
		Left:   "up",
		Middle: "up",
		Right:  "up",
	})
}

func (m *MouseClient) MouseXY(params map[string]interface{}, _ *structures.RunContext) {
	x := params["x"].(float64)
	y := params["y"].(float64)
	m.mouseXY(x, y)
}

func (m *MouseClient) mouseXY(x float64, y float64) {
	xBalanced := float32(x)
	yBalanced := float32(y)
	mouse.MovePress(mouse.MouseJSON{
		X:      &xBalanced,
		Y:      &yBalanced,
		Left:   "up",
		Middle: "up",
		Right:  "up",
	})
}
