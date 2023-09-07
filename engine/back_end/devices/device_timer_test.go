package devices

import "testing"

func TestGetTimerCallbacks(t *testing.T) {
	methods := GetTimerCallbacks()
	if methods == nil {
		t.Error("GetTimerCallbacks() should not return nil")
	}
	if len(methods) == 0 {
		t.Error("GetTimerCallbacks() should not return empty")
	}
	if len(methods) != 4 {
		t.Error("GetTimerCallbacks() should return 3 methods")
	}
	if methods[0].Identifier != "quando.time.after" {
		t.Error("GetTimerCallbacks()[0] should have identifier quando.time.after")
	}
	if methods[1].Identifier != "quando.time.every" {
		t.Error("GetTimerCallbacks()[1] should have identifier quando.time.every")
	}
	if methods[2].Identifier != "quando.time.per" {
		t.Error("GetTimerCallbacks()[2] should have identifier quando.time.per")
	}
	if methods[3].Identifier != "quando.time.vary" {
		t.Error("GetTimerCallbacks()[3] should have identifier quando.time.vary")
	}
}
