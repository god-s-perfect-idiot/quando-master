package structures

import "testing"

var callgraphDefinitions = []Definition{
	Definition{
		Line: 0,
		Signature: CallSignature{
			MethodIdentifier: "test",
			Parameters:       make([]Parameter, 0),
		},
		Type: "invocation",
	},
}

var callgraphDefinitionsEmpty = []Definition{}

var callgraphDefinitionsMulti = []Definition{
	Definition{
		Line: 0,
		Signature: CallSignature{
			MethodIdentifier: "test",
			Parameters:       make([]Parameter, 0),
		},
		Type: "invocation",
	},
	Definition{
		Line: 1,
		Signature: CallSignature{
			MethodIdentifier: "test",
			Parameters:       make([]Parameter, 0),
		},
		Type: "invocation",
	},
}

var callgraphDefinitionsWithChild = []Definition{
	Definition{
		Line: 0,
		Signature: CallSignature{
			MethodIdentifier: "test",
			Parameters:       make([]Parameter, 0),
			HasCallback:      true,
		},
		Type: "invocation",
	},
	Definition{
		Line: 1,
		Signature: CallSignature{
			MethodIdentifier: "test",
			Parameters:       make([]Parameter, 0),
		},
		Type: "invocation",
	},
	Definition{Line: 2,
		Signature: CallSignature{},
		Type:      "callback terminator",
	},
}

func TestSingleDAG(t *testing.T) {
	callgraph := ConstructCallGraph(callgraphDefinitions)
	if callgraph.Roots == nil {
		t.Error("ConstructCallGraph() should not return nil roots")
	}
	if len(callgraph.Roots) != 1 {
		t.Error("ConstructCallGraph() should return one root")
	}
	if callgraph.Roots[0].identifier != "0:action:test" {
		t.Error("ConstructCallGraph() should return root with identifier test")
	}
}

func TestSingleDAGWithChild(t *testing.T) {
	callgraph := ConstructCallGraph(callgraphDefinitionsWithChild)
	if callgraph.Roots == nil {
		t.Error("ConstructCallGraph() should not return nil roots")
	}
	if len(callgraph.Roots) != 1 {
		t.Error("ConstructCallGraph() should return one root")
	}
	if callgraph.Roots[0].identifier != "0:callback:test" {
		t.Error("ConstructCallGraph() should return root with identifier test")
	}
	if len(callgraph.Roots[0].Children) != 1 {
		t.Error("ConstructCallGraph() should return root with one child")
	}
	if callgraph.Roots[0].Children[0].identifier != "1:action:test" {
		t.Error("ConstructCallGraph() should return root with child with identifier test")
	}
}

func TestEmptyDAG(t *testing.T) {
	callgraph := ConstructCallGraph(callgraphDefinitionsEmpty)
	if len(callgraph.Roots) != 0 {
		t.Error("ConstructCallGraph() should return zero roots")
	}
}

func TestMultiDAG(t *testing.T) {
	callgraph := ConstructCallGraph(callgraphDefinitionsMulti)
	if callgraph.Roots == nil {
		t.Error("ConstructCallGraph() should not return nil roots")
	}
	if len(callgraph.Roots) != 2 {
		t.Error("ConstructCallGraph() should return two roots")
	}
	if callgraph.Roots[0].identifier != "0:action:test" {
		t.Error("ConstructCallGraph() should return root with identifier test")
	}
	if callgraph.Roots[1].identifier != "1:action:test" {
		t.Error("ConstructCallGraph() should return root with identifier test")
	}
}

func TestGetRoots(t *testing.T) {
	callgraph := ConstructCallGraph(callgraphDefinitions)
	roots := callgraph.GetRoots()
	if roots == nil {
		t.Error("GetRoots() should not return nil")
	}
	if len(roots) != 1 {
		t.Error("GetRoots() should return one root")
	}
	if roots[0].identifier != "0:action:test" {
		t.Error("GetRoots() should return root with identifier test")
	}
}

func TestAttachMethod(t *testing.T) {
	callgraph := ConstructCallGraph(callgraphDefinitions)
	method := Method{
		Identifier: "test",
		Function:   "test",
		Type:       "test",
	}
	callgraph.AttachMethod("test", method)
	if callgraph.Roots[0].Method.Identifier != "test" {
		t.Error("AttachMethod() should attach method to root")
	}
}

func TestAttachMethodToNode(t *testing.T) {
	callgraph := ConstructCallGraph(callgraphDefinitions)
	method := Method{
		Identifier: "test",
		Function:   "test",
		Type:       "test",
	}
	callgraph.AttachMethodToNode(callgraph.Roots[0], "test", method)
	if callgraph.Roots[0].Method.Identifier != "test" {
		t.Error("AttachMethodToNode() should attach method to root")
	}
}
