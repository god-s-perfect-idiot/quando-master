package actions

import "testing"

func TestGetMouseActions(t *testing.T) {
	actions := GetMouseActions()
	if len(actions) != 3 {
		t.Errorf("Expected 3 actions, got %d", len(actions))
	}
}

func TestNewMouse(t *testing.T) {
	mouse := NewMouse()
	if mouse == nil {
		t.Errorf("Expected mouse to be initialized")
	}
}

func TestMouseXAction(t *testing.T) {
	actions := GetMouseActions()
	mouseX := actions[0]
	if mouseX.Identifier != "quando.control.mouseX" {
		t.Errorf("Expected identifier to be quando.control.mouseX, got %s", mouseX.Identifier)
	}
	if mouseX.Type != "action" {
		t.Errorf("Expected type to be action, got %s", mouseX.Type)
	}
}

func TestMouseYAction(t *testing.T) {
	actions := GetMouseActions()
	mouseY := actions[1]
	if mouseY.Identifier != "quando.control.mouseY" {
		t.Errorf("Expected identifier to be quando.control.mouseY, got %s", mouseY.Identifier)
	}
	if mouseY.Type != "action" {
		t.Errorf("Expected type to be action, got %s", mouseY.Type)
	}
}

func TestMouseXYAction(t *testing.T) {
	actions := GetMouseActions()
	mouseXY := actions[2]
	if mouseXY.Identifier != "quando.control.mouse" {
		t.Errorf("Expected identifier to be quando.control.mouse, got %s", mouseXY.Identifier)
	}
	if mouseXY.Type != "action" {
		t.Errorf("Expected type to be action, got %s", mouseXY.Type)
	}
}
