package devices

import "testing"

func TestGetMouseCallbacks(t *testing.T) {
	callbacks := GetMouseCallbacks()
	if len(callbacks) != 2 {
		t.Error("GetMouseCallbacks should return 2 callback")
	}
	if callbacks[0].Identifier != "quando.mouse.handleX" {
		t.Error("GetMouseCallbacks should return a callback with identifier 'quando.mouse.move'")
	}
	if callbacks[1].Identifier != "quando.mouse.handleY" {
		t.Error("GetMouseCallbacks should return a callback with identifier 'quando.mouse.move'")
	}
}
