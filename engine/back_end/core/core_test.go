package core

import (
	"context"
	"quando/engine/structures"
	"testing"
)

var definitions = []structures.Definition{
	structures.Definition{
		Line: 0,
		Type: "invocation",
		Signature: structures.CallSignature{
			MethodIdentifier: "dummy",
			Parameters: []structures.Parameter{
				structures.Parameter{
					Identifier: "dummy",
					Type:       "STRING",
					Value:      "\"dummy\"",
				},
			},
		},
	},
}
var callgraph = structures.ConstructCallGraph(definitions)
var invocations = structures.GenerateInvocationTable(definitions)
var essence = structures.Essence{
	DependencyGraph: &callgraph,
	Invocations:     &invocations,
	CallStack:       structures.NewStack(),
	Context:         context.Background(),
	Val:             0.0,
}
var mockedAPIDummy1 = structures.Method{
	Identifier: "dummy",
	Function: func(params map[string]interface{}) (float64, map[string]interface{}) {
		return 0.1111, nil
	},
	Type: "callback",
}
var mockedAPIDummy2 = structures.Method{
	Identifier: "dummy",
	Function: func(params map[string]interface{}) (float64, map[string]interface{}) {
		return 0.2222, nil
	},
	Type: "callback",
}

func TestExecute(t *testing.T) {
	essence.DependencyGraph.AttachMethod("dummy", mockedAPIDummy1)
	Execute(&essence)
	if essence.Val != 0.1111 {
		t.Error("Essence should have val 0.1111")
	}
}

func TestRunNode(t *testing.T) {
	node := essence.DependencyGraph.GetRoots()[0]
	node.Method = mockedAPIDummy2
	RunNode(node, &essence)
	if essence.Val != 0.2222 {
		t.Error("Essence should have val 0.2222")
	}
}
