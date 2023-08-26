package core

import "quando/engine/structures"

func RunNode(node *structures.CallNode, essence *structures.Essence) {
	essence.CallStack.Push(node)
	if node.Method.Identifier != "" {
		//switch node.Method.Type {
		//case "callback":
		//	//
		//case "action":
		node.Method.Call(node.Definition.Signature.Parameters, essence)
		// TODO FIX ME
		for _, child := range node.MainChildren {
			RunNode(child, essence)
		}
	}
}

func Execute(essence *structures.Essence) {
	roots := essence.DependencyGraph.GetRoots()
	for _, root := range roots {
		RunNode(root, essence)
	}
}
