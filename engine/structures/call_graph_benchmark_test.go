package structures

import "testing"

func BenchmarkCallGraph_appendChild(b *testing.B) {
	node := &CallNode{
		MainChildren: []*CallNode{},
		AltChildren:  []*CallNode{},
	}
	c := CallGraph{}
	for i := 0; i < b.N; i++ {
		c.appendChild(node, node, false)
	}
}

func BenchmarkCallGraph_GetRoots(b *testing.B) {
	c := CallGraph{}
	for i := 0; i < b.N; i++ {
		c.GetRoots()
	}
}

func BenchmarkCallGraph_AttachMethodToNode(b *testing.B) {
	node := &CallNode{
		MainChildren: []*CallNode{},
		AltChildren:  []*CallNode{},
	}
	c := CallGraph{}
	for i := 0; i < b.N; i++ {
		c.AttachMethodToNode(node, "test", Method{})
	}
}

func BenchmarkCallGraph_AttachMethod(b *testing.B) {
	c := CallGraph{}
	for i := 0; i < b.N; i++ {
		c.AttachMethod("test", Method{})
	}
}

func BenchmarkConstructCallGraph(b *testing.B) {
	definitions := []Definition{}
	for i := 0; i < 100; i++ {
		definitions = append(definitions, Definition{
			Signature: CallSignature{
				MethodIdentifier: "test",
			},
		})
	}
	for i := 0; i < b.N; i++ {
		ConstructCallGraph(definitions)
	}
}

func BenchmarkConstructCallGraph_100(b *testing.B) {
	definitions := []Definition{}
	for i := 0; i < 100; i++ {
		definitions = append(definitions, Definition{
			Signature: CallSignature{
				MethodIdentifier: "test",
			},
		})
	}
	for i := 0; i < b.N; i++ {
		ConstructCallGraph(definitions)
	}
}

func BenchmarkConstructCallGraph_1000(b *testing.B) {
	definitions := []Definition{}
	for i := 0; i < 1000; i++ {
		definitions = append(definitions, Definition{
			Signature: CallSignature{
				MethodIdentifier: "test",
			},
		})
	}
	for i := 0; i < b.N; i++ {
		ConstructCallGraph(definitions)
	}
}

func BenchmarkConstructCallGraph_B(b *testing.B) {
	definitions := []Definition{}
	for i := 0; i < b.N; i++ {
		definitions = append(definitions, Definition{
			Signature: CallSignature{
				MethodIdentifier: "test",
			},
		})
	}
	for i := 0; i < b.N; i++ {
		ConstructCallGraph(definitions)
	}
}
