package analyser

import "testing"

func TestIsBoolean(t *testing.T) {
	if !IsBooleanParameter("true") {
		t.Error("IsBooleanParameter() should return true")
	}
	if !IsBooleanParameter("false") {
		t.Error("IsBooleanParameter() should return true")
	}
	if IsBooleanParameter("1") {
		t.Error("IsBoolean() should return false")
	}
	if IsBooleanParameter("0") {
		t.Error("IsBooleanParameter() should return false")
	}
	if IsBooleanParameter("true1") {
		t.Error("IsBooleanParameter() should return false")
	}
	if IsBooleanParameter("false1") {
		t.Error("IsBooleanParameter() should return false")
	}
}

func TestIsInteger(t *testing.T) {
	if !IsIntegerParameter("1") {
		t.Error("IsIntegerParameter() should return true")
	}
	if !IsIntegerParameter("0") {
		t.Error("IsIntegerParameter() should return true")
	}
	if IsIntegerParameter("true") {
		t.Error("IsIntegerParameter() should return false")
	}
	if IsIntegerParameter("false") {
		t.Error("IsIntegerParameter() should return false")
	}
	if IsIntegerParameter("1true") {
		t.Error("IsIntegerParameter() should return false")
	}
	if IsIntegerParameter("0false") {
		t.Error("IsIntegerParameter() should return false")
	}
}

func TestIsFloat(t *testing.T) {
	if !IsFloatParameter("1.0") {
		t.Error("IsFloatParameter() should return true")
	}
	if !IsFloatParameter("0.0") {
		t.Error("IsFloatParameter() should return true")
	}
	if IsFloatParameter("true") {
		t.Error("IsFloatParameter() should return false")
	}
	if IsFloatParameter("false") {
		t.Error("IsFloatParameter() should return false")
	}
	if IsFloatParameter("1.0true") {
		t.Error("IsFloatParameter() should return false")
	}
	if IsFloatParameter("0.0false") {
		t.Error("IsFloatParameter() should return false")
	}
}

func TestIsString(t *testing.T) {
	if !IsStringParameter(`"string"`) {
		t.Error("IsStringParameter() should return true")
	}
	if IsStringParameter("true") {
		t.Error("IsStringParameter() should return false")
	}
	if IsStringParameter("false") {
		t.Error("IsStringParameter() should return false")
	}
	if IsStringParameter(`"string"true`) {
		t.Error("IsStringParameter() should return false")
	}
	if IsStringParameter(`"string"false`) {
		t.Error("IsStringParameter() should return false")
	}
	if IsStringParameter(`string`) {
		t.Error("IsStringParameter() should return false")
	}
	if IsStringParameter(`'string'`) {
		t.Error("IsStringParameter() should return false")
	}
}

func TestIsCallback(t *testing.T) {
	if !IsCallbackParameter(`{`) {
		t.Error("IsCallbackParameter() should return true")
	}
	if IsCallbackParameter("true") {
		t.Error("IsCallbackParameter() should return false")
	}
	if IsCallbackParameter("false") {
		t.Error("IsCallbackParameter() should return false")
	}
}

func TestIsVal(t *testing.T) {
	if !IsVal("val") {
		t.Error("IsVal() should return true")
	}
	if IsVal("true") {
		t.Error("IsVal() should return false")
	}
	if IsVal("false") {
		t.Error("IsVal() should return false")
	}
	if IsVal("valval") {
		t.Error("IsVal() should return false")
	}
}
