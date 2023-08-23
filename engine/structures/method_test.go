package structures

import "testing"

func TestCleanStringParam(t *testing.T) {
	essence := Essence{}
	param := Parameter{
		Identifier: "test",
		Type:       "STRING",
		Value:      "\"test\"",
	}
	m := Method{}
	param = m.cleanParam(param, &essence)
	if param.Value != "test" {
		t.Error("Method cleanParam should return param with value test")
	}
}

func TestCleanBooleanParam(t *testing.T) {
	essence := Essence{}
	param := Parameter{
		Identifier: "test",
		Type:       "BOOLEAN",
		Value:      "true",
	}
	m := Method{}
	param = m.cleanParam(param, &essence)
	if param.Value != true {
		t.Error("Method cleanParam should return param with value true")
	}
}

func TestCleanNumberParam(t *testing.T) {
	essence := Essence{}
	param := Parameter{
		Identifier: "test",
		Type:       "NUMBER",
		Value:      "1",
	}
	m := Method{}
	param = m.cleanParam(param, &essence)
	if param.Value != 1 {
		t.Error("Method cleanParam should return param with value 1")
	}
}

func TestCleanValParam(t *testing.T) {
	essence := Essence{
		Val: 0.5,
	}
	param := Parameter{
		Identifier: "test",
		Type:       "VAL",
		Value:      "val",
	}
	m := Method{}
	param = m.cleanParam(param, &essence)
	if param.Value != 0.5 {
		t.Error("Method cleanParam should return param with value 0.5")
	}
}

func TestCall(t *testing.T) {
	essence := Essence{}
	param := Parameter{
		Identifier: "test",
		Type:       "NUMBER",
		Value:      "1",
	}
	m := Method{
		Identifier: "test",
		Function: func(params map[string]interface{}) float64 {
			return 1.1
		},
		Type: "callback",
	}
	m.Call([]Parameter{param}, &essence)
	if essence.Val != 1.1 {
		t.Error("Method Call should return essence with val 2")
	}
}

func TestCallWithVal(t *testing.T) {
	essence := Essence{
		Val: 0.5,
	}
	param := Parameter{
		Identifier: "test",
		Type:       "VAL",
		Value:      "val",
	}
	m := Method{
		Identifier: "test",
		Function: func(params map[string]interface{}) float64 {
			if params["test"] == 0.5 {
				return 0.1
			} else {
				return 0.0
			}
		},
		Type: "callback",
	}
	m.Call([]Parameter{param}, &essence)
	if essence.Val != 0.1 {
		t.Error("Method Call should return essence with val 0.5")
	}
}

func TestCallWithSkippedCallback(t *testing.T) {
	essence := Essence{}
	param := Parameter{
		Identifier: "callback",
		Type:       "CALLBACK",
		Value:      "{",
	}
	m := Method{
		Identifier: "test",
		Function: func(params map[string]interface{}) float64 {
			if len(params) == 0 {
				return 0.0
			} else {
				return 0.1
			}
		},
		Type: "callback",
	}
	m.Call([]Parameter{param}, &essence)
	if essence.Val != 0.0 {
		t.Error("Method Call should return essence with val 0.0")
	}
}
