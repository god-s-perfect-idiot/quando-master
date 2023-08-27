package devices

import "testing"

func TestGetKeyboardCallbacks(t *testing.T) {
	methods := GetKeyboardCallbacks()
	if methods == nil {
		t.Error("GetKeyboardCallbacks() should not return nil")
	}
	if len(methods) == 0 {
		t.Error("GetKeyboardCallbacks() should not return empty")
	}
	if len(methods) != 1 {
		t.Error("GetKeyboardCallbacks() should return 1 methods")
	}
	if methods[0].Identifier != "quando.key.handleKey" {
		t.Error("GetKeyboardCallbacks()[0] should have identifier quando.key.handleKey")
	}
}
