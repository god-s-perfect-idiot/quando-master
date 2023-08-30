package core

import "quando/engine/structures"

func RunNode(node *structures.CallNode, executable *structures.Executable) {
	executable.CallStack.Push(node)
	if node.Method.Identifier != "" {
		switch node.Type {
		case "action":
			node.Method.Call(node.Definition.Signature.Parameters, executable)
		case "callback":
			count := 1
			if node.Method.IsIterator() {
				count = -1
			}
			//TODO Fix break
			for i := 0; i != count; i++ {
				if node.Method.IsArbiter() {
					if !executable.ValidateData("sequence") {
						executable.SetData("sequence", structures.GenerateRandomSequence(len(node.MainChildren)))
					}
					if !executable.ValidateData("keys") {
						executable.SetData("keys", []int{})
					}
					if !executable.ValidateData("nodeCount") {
						executable.SetData("nodeCount", len(node.MainChildren))
					}
					node.Method.Call(node.Definition.Signature.Parameters, executable)
					selection := executable.GetData("selection").(int)
					RunNode(node.MainChildren[selection], executable)
				} else {
					node.Method.Call(node.Definition.Signature.Parameters, executable)
					for _, child := range node.MainChildren {
						RunNode(child, executable)
					}
				}
			}

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