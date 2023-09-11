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

func TestCallFunc(t *testing.T) {
	essence := Executable{
		DependencyGraph: &CallGraph{
			Roots: []*CallNode{
				{
					Method: Method{
						Identifier: "test",
						Function: func(params map[string]interface{}, ctx *RunContext) {
							ctx.Executable.Val = 1.1
						},
					},
				},
			},
		},
	}
	m := essence.DependencyGraph.Roots[0].Method
	m.CallFunc(&essence, essence.DependencyGraph.Roots[0])
	if essence.Val != 1.1 {
		t.Error("Method CallFunc should return essence with val 2")
	}
}

func TestCallFuncWithVal(t *testing.T) {
	essence := Executable{
		Val: 0.5,
		DependencyGraph: &CallGraph{
			Roots: []*CallNode{
				{
					Method: Method{
						Identifier: "test",
						Function: func(params map[string]interface{}, ctx *RunContext) {
							if params["test"] == 0.5 {
								ctx.Executable.Val = 0.1
							} else {
								ctx.Executable.Val = 0.0
							}
						},
					},
					Definition: Definition{
						Signature: CallSignature{
							Parameters: []Parameter{
								{
									Identifier: "test",
									Type:       "VAL",
									Value:      "val",
								},
							},
						},
					},
				},
			},
		},
	}
	m := essence.DependencyGraph.Roots[0].Method
	m.CallFunc(&essence, essence.DependencyGraph.Roots[0])
	if essence.Val != 0.1 {
		t.Error("Method CallFunc should return essence with val 0.5")
	}
}

func TestCallFuncWithValAndData(t *testing.T) {
	essence := Executable{
		Val:  0.5,
		Data: make(map[string]interface{}),
		DependencyGraph: &CallGraph{
			Roots: []*CallNode{
				{
					Method: Method{
						Identifier: "test",
						Function: func(params map[string]interface{}, ctx *RunContext) {
							if params["test"] == 0.5 {
								ctx.Executable.Val = 0.1
								ctx.Executable.Data["test"] = 0.2
							} else {
								ctx.Executable.Val = 0.0
							}
						},
					},
					Definition: Definition{
						Signature: CallSignature{
							Parameters: []Parameter{
								{
									Identifier: "test",
									Type:       "VAL",
									Value:      "val",
								},
							},
						},
					},
				},
			},
		},
	}
	m := essence.DependencyGraph.Roots[0].Method
	m.CallFunc(&essence, essence.DependencyGraph.Roots[0])
	if essence.Val != 0.1 {
		t.Error("Method CallFunc should return essence with val 0.5")
	}
	if essence.Data["test"] != 0.2 {
		t.Error("Method CallFunc should return essence with data test 0.2")
	}
}
