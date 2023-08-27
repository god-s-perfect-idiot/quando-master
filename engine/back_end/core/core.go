package core

import "quando/engine/structures"

func RunNode(node *structures.CallNode, essence *structures.Essence) {
	essence.CallStack.Push(node)
	if node.Method.Identifier != "" {
		switch node.Type {
		case "action":
			node.Method.Call(node.Definition.Signature.Parameters, essence)
		case "callback":
			count := 1
			if node.Method.IsIterator() {
				count = -1
			}
			//TODO Fix break
			for i := 0; i != count; i++ {
				if node.Method.IsArbiter() {
					if !essence.ValidateData("sequence") {
						essence.SetData("sequence", structures.GenerateRandomSequence(len(node.MainChildren)))
					}
					if !essence.ValidateData("keys") {
						essence.SetData("keys", []int{})
					}
					if !essence.ValidateData("nodeCount") {
						essence.SetData("nodeCount", len(node.MainChildren))
					}
					node.Method.Call(node.Definition.Signature.Parameters, essence)
					selection := essence.GetData("selection").(int)
					RunNode(node.MainChildren[selection], essence)
				} else {
					node.Method.Call(node.Definition.Signature.Parameters, essence)
					for _, child := range node.MainChildren {
						RunNode(child, essence)
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

func Execute(essence *structures.Essence) {
	roots := essence.DependencyGraph.GetRoots()
	for _, root := range roots {
		RunNode(root, essence)
	}
}
