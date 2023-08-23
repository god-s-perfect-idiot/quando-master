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
}
