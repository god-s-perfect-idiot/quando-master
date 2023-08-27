package core

import "quando/engine/structures"

func RunNode(node *structures.CallNode, essence *structures.Essence) {
	essence.CallStack.Push(node)
	if node.Method.Identifier != "" {
		//switch node.Method.Type {
		//case "callback":
		//	//
		//case "action":
		switch node.Type {
		case "action":
			node.Method.Call(node.Definition.Signature.Parameters, essence)
		case "callback":
			node.Method.Call(node.Definition.Signature.Parameters, essence)
			for _, child := range node.MainChildren {
				RunNode(child, essence)
			}
		case "conditional callback":
			// TODO Implement Conditional Callbacks
			// Logic : Perform Main Child sequentially until condition
			// Then switch to Alt Child and perform sequentially until condition resets
		}
		node.Method.Call(node.Definition.Signature.Parameters, essence)
		// TODO FIX ME
	}
}

func Execute(essence *structures.Essence) {
	roots := essence.DependencyGraph.GetRoots()
	for _, root := range roots {
		RunNode(root, essence)
	}
}
