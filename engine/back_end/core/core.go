package core

import "quando/engine/structures"

func RunNode(node *structures.CallNode, executable *structures.Executable) {
	if node.Method.Identifier != "" {
		switch node.Type {
		case "action":
			node.Method.CallFunc(executable, node)
		case "callback":
			node.Method.CallFunc(executable, node)
		case "conditional callback":
			// TODO Implement Conditional Callbacks
			// Logic : Perform Main Child sequentially until condition
			// Then switch to Alt Child and perform sequentially until condition resets
		}
	}
}

func Execute(essence *structures.Executable) {
	roots := essence.DependencyGraph.GetRoots()
	for _, root := range roots {
		RunNode(root, essence)
	}
}
