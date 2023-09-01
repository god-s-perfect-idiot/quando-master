package actions

import (
	"testing"
)

func TestNewKeyboard(t *testing.T) {
	k := NewKeyboard()
	if k == nil {
		t.Error("NewKeyboard() should not return nil")
	}
}

// TODO: Mock key function and test that it is called

func TestKeyboardGetKeyboardActions(t *testing.T) {
	keyboardActions := GetKeyboardActions()
	if len(keyboardActions) == 0 {
		t.Error("No keyboard actions found")
	}
}

func TestControlKeyAction(t *testing.T) {
	keyboardActions := GetKeyboardActions()
	controlKey := keyboardActions[0]
	if controlKey.Identifier != "quando.control.key" {
		t.Error("Identifier should be quando.control.key")
	}
	if controlKey.Type != "action" {
		t.Error("Type should be action")
	}
}
