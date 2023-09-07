package structures

import "testing"

func TestNewCallSignature(t *testing.T) {
	c := NewCallSignature("test", []Parameter{}, false)
	if c == nil {
		t.Error("NewCallSignature() should not return nil")
	}
}

func TestNewCallSignatureControls(t *testing.T) {
	c := NewCallSignature("test", []Parameter{}, false)
	if c.MethodIdentifier != "test" {
		t.Error("CallSignature method identifier should be test")
	}
	if c.Parameters == nil {
		t.Error("CallSignature parameters should not be nil")
	}
	if c.HasCallback {
		t.Error("CallSignature has callback should be false")
	}
}

func TestNewCallSignatureWithCallbackControls(t *testing.T) {
	c := NewCallSignature("test", []Parameter{}, true)
	if c.MethodIdentifier != "test" {
		t.Error("CallSignature method identifier should be test")
	}
	if c.Parameters == nil {
		t.Error("CallSignature parameters should not be nil")
	}
	if !c.HasCallback {
		t.Error("CallSignature has callback should be true")
	}
}

func TestGetMethodIdentifier(t *testing.T) {
	c := NewCallSignature("test", []Parameter{}, false)
	if c.GetMethodIdentifier() != "test" {
		t.Error("CallSignature GetMethodIdentifier() should return test")
	}
}

func TestGetParameters(t *testing.T) {
	c := NewCallSignature("test", []Parameter{}, false)
	if c.GetParameters() == nil {
		t.Error("CallSignature GetParameters() should not return nil")
	}
}

func TestGetParameterCount(t *testing.T) {
	c := NewCallSignature("test", []Parameter{}, false)
	if c.GetParameterCount() != 0 {
		t.Error("CallSignature GetParameterCount() should return 0")
	}
}

func TestGetParameter(t *testing.T) {
	c := NewCallSignature("test", []Parameter{
		Parameter{
			Identifier: "test",
			Type:       "test",
			Value:      "test",
		},
	}, false)
	if c.GetParameter(0).Identifier != "test" {
		t.Error("CallSignature GetParameter(0) should return test")
	}
	if c.GetParameter(0).Type != "test" {
		t.Error("CallSignature GetParameter(0) should return test")
	}
	if c.GetParameter(0).Value != "test" {
		t.Error("CallSignature GetParameter(0) should return test")
	}
}

func TestGetParameterIdentifier(t *testing.T) {
	c := NewCallSignature("test", []Parameter{
		Parameter{
			Identifier: "test",
			Type:       "test",
			Value:      "test",
		},
	}, false)
	if c.GetParameterIdentifier(0) != "test" {
		t.Error("CallSignature GetParameterIdentifier(0) should return test")
	}
}

func TestGetParameterType(t *testing.T) {
	c := NewCallSignature("test", []Parameter{
		Parameter{
			Identifier: "test",
			Type:       "test",
			Value:      "test",
		},
	}, false)
	if c.GetParameterType(0) != "test" {
		t.Error("CallSignature GetParameterType(0) should return test")
	}
}

func TestGetParameterValue(t *testing.T) {
	c := NewCallSignature("test", []Parameter{
		Parameter{
			Identifier: "test",
			Type:       "test",
			Value:      "test",
		},
	}, false)
	if c.GetParameterValue(0) != "test" {
		t.Error("CallSignature GetParameterValue(0) should return test")
	}
}

func TestCallbackExists(t *testing.T) {
	c := NewCallSignature("test", []Parameter{}, false)
	if c.CallbackExists() {
		t.Error("CallSignature CallbackExists() should return false")
	}
}

func TestCallbackExistsWithCallback(t *testing.T) {
	c := NewCallSignature("test", []Parameter{}, true)
	if !c.CallbackExists() {
		t.Error("CallSignature CallbackExists() should return true")
	}
}
