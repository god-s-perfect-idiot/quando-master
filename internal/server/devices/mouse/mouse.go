//go:build full

package mouse

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-vgo/robotgo"
)

type MouseJSON struct {
	X      *float32 `json:"x,omitempty"`
	Y      *float32 `json:"y,omitempty"`
	Left   string   `json:"left"`
	Middle string   `json:"middle"`
	Right  string   `json:"right"`
}

type mouseJSON struct {
	X      *float32 `json:"x,omitempty"`
	Y      *float32 `json:"y,omitempty"`
	Left   string   `json:"left"`
	Middle string   `json:"middle"`
	Right  string   `json:"right"`
}

var (
	last_left   = ""
	last_middle = ""
	last_right  = ""
)

func check_mouse(button_action string, last_action *string, button_name string) {
	if button_action != "" {
		if button_action != *last_action {
			switch button_action {
			case "click":
				robotgo.Click(button_name)
				*last_action = "up"
			case "up":
				robotgo.Toggle(button_name, "up")
				*last_action = button_action
			case "down":
				robotgo.Toggle(button_name, "down")
				*last_action = button_action
			default:
				fmt.Println("runtime error or hacking - unexpected mouse button action:", button_action, " for ", button_name)
			}
		}
	}
}

func MovePress(mouse MouseJSON) {
	mouseVal := mouseJSON{
		X:      mouse.X,
		Y:      mouse.Y,
		Left:   mouse.Left,
		Middle: mouse.Middle,
		Right:  mouse.Right,
	}
	move_press(mouseVal)
}

func move_press(mouse mouseJSON) {
	check_mouse(mouse.Left, &last_left, "left")
	check_mouse(mouse.Right, &last_right, "right")
	check_mouse(mouse.Middle, &last_middle, "center")

	if mouse.X != nil && mouse.Y != nil {
		width, height := robotgo.GetScreenSize()
		x := int(*mouse.X * float32(width))
		y := int(*mouse.Y * float32(height))
		robotgo.Move(x, y)
	}
}

func HandleMouse(w http.ResponseWriter, req *http.Request) {
	var mouse mouseJSON
	err := json.NewDecoder(req.Body).Decode(&mouse)
	if err != nil {
		fmt.Println("Error parsing request", err)
		return
	}
	move_press(mouse)
}
