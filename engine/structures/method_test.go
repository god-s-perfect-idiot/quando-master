package structures

import "testing"

func TestCleanStringParam(t *testing.T) {
	essence := Executable{}
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
	essence := Executable{}
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
	essence := Executable{}
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
	essence := Executable{
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

func TestIsIterator(t *testing.T) {
	m := Method{
		Iterator: true,
	}
	if !m.IsIterator() {
		t.Error("Method IsIterator should return true")
	}
}

func TestIsArbiter(t *testing.T) {
	m := Method{
		Arbiter: true,
	}
	if !m.IsArbiter() {
		t.Error("Method IsArbiter should return true")
	}
}

func TestCall(t *testing.T) {
	essence := Executable{}
	param := Parameter{
		Identifier: "test",
		Type:       "NUMBER",
		Value:      "1",
	}
	m := Method{
		Identifier: "test",
		Function: func(params map[string]interface{}) (float64, map[string]interface{}) {
			return 1.1, nil
		},
		Type: "callback",
	}
	m.Call([]Parameter{param}, &essence)
	if essence.Val != 1.1 {
		t.Error("Method Call should return essence with val 2")
	}
}

func TestCallWithVal(t *testing.T) {
	essence := Executable{
		Val: 0.5,
	}
	param := Parameter{
		Identifier: "test",
		Type:       "VAL",
		Value:      "val",
	}
	m := Method{
		Identifier: "test",
		Function: func(params map[string]interface{}) (float64, map[string]interface{}) {
			if params["test"] == 0.5 {
				return 0.1, nil
			} else {
				return 0.0, nil
			}
		},
		Type: "callback",
	}
	m.Call([]Parameter{param}, &essence)
	if essence.Val != 0.1 {
		t.Error("Method Call should return essence with val 0.5")
	}
}

func TestCallWithValAndData(t *testing.T) {
	essence := Executable{
		Val:  0.5,
		Data: make(map[string]interface{}),
	}
	param := Parameter{
		Identifier: "test",
		Type:       "VAL",
		Value:      "val",
	}
	m := Method{
		Identifier: "test",
		Function: func(params map[string]interface{}) (float64, map[string]interface{}) {
			if params["test"] == 0.5 {
				return 0.1, map[string]interface{}{
					"test": 0.2,
				}
			} else {
				return 0.0, nil
			}
		},
		Type: "callback",
	}
	m.Call([]Parameter{param}, &essence)
	if essence.Val != 0.1 {
		t.Error("Method Call should return essence with val 0.5")
	}
	if essence.Data["test"] != 0.2 {
		t.Error("Method Call should return essence with data test 0.2")
	}
}
