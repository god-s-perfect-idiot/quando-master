package generator

import (
	"context"
	"quando/engine/structures"
	"testing"
)

func TestNewGenerator(t *testing.T) {
	g := NewGenerator()
	if g == nil {
		t.Error("NewGenerator() should not return nil")
	}
}

func TestGeneratorLookups(t *testing.T) {
	g := NewGenerator()
	if g.Lookups == (Lookups{}) {
		t.Error("Generator lookups should not be empty")
	}
	if g.Lookups.Callbacks == nil {
		t.Error("Generator callbacks should not be nil")
	}
	if g.Lookups.Actions == nil {
		t.Error("Generator actions should not be nil")
	}
}

func TestGeneratorGenerateCode(t *testing.T) {
	g := NewGenerator()
	definitions := []structures.Definition{
		structures.Definition{
			Line: 0,
			Type: "invocation",
			Signature: structures.CallSignature{
				MethodIdentifier: "quando.testCallback",
				Parameters: []structures.Parameter{
					structures.Parameter{
						Identifier: "quando.test",
						Type:       "STRING",
						Value:      "\"test\"",
					},
				},
				HasCallback: true,
			},
		},
		structures.Definition{
			Line: 0,
			Type: "invocation",
			Signature: structures.CallSignature{
				MethodIdentifier: "quando.testAction",
				Parameters: []structures.Parameter{
					structures.Parameter{
						Identifier: "quando.test",
						Type:       "STRING",
						Value:      "\"test\"",
					},
				},
				HasCallback: false,
			},
		},
	}
	callGraph := structures.ConstructCallGraph(definitions)
	invocations := structures.GenerateInvocationTable(definitions)
	essence := structures.Executable{
		DependencyGraph: &callGraph,
		Invocations:     &invocations,
		CallStack:       structures.NewStack(),
		Context:         context.Background(),
		Val:             0.0,
	}
	g.Lookups = Lookups{
		Callbacks: &structures.LookupTable{
			Methods: []structures.Method{
				structures.Method{
					Identifier: "quando.testCallback",
					Function: func(params map[string]interface{}) float64 {
						return 0.0
					},
					Type: "action",
				},
			},
		},
		Actions: &structures.LookupTable{
			Methods: []structures.Method{
				structures.Method{
					Identifier: "quando.testAction",
					Function: func(params map[string]interface{}) float64 {
						return 0.0
					},
					Type: "action",
				},
			},
		},
	}
	g.GenerateCode(essence)
	if essence.DependencyGraph.GetRoots()[0].Method.Identifier != "quando.testCallback" {
		t.Error("Executable dependency graph should have quando.test as root")
	}
	if essence.DependencyGraph.GetRoots()[0].MainChildren[0].Method.Identifier != "quando.testAction" {
		t.Error("Executable dependency graph should have quando.test as root")
	}
}
