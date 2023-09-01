package core

import "quando/engine/structures"

//func RunAction(action *structures.CallNode, executable *structures.Executable) {
//	action.Method.Call(action.Definition.Signature.Parameters, executable)
//}
//
//func RunCallback(callback *structures.CallNode, executable *structures.Executable) {
//}

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

//func RunNode(node *structures.CallNode, executable *structures.Executable) {
//	executable.CallStack.Push(node)
//	select {
//	// TODO Fix break THIS ISNT WORKING
//	case callData := <-*executable.CallPipe:
//		if callData["crash"] == executable.Hash {
//			//return
//		}
//	default:
//		if node.Method.Identifier != "" {
//			switch node.Type {
//			case "action":
//				node.Method.Call(node.Definition.Signature.Parameters, executable)
//			case "callback":
//				//count := 1
//				//if node.Method.IsIterator() {
//				//	count = -1
//				//}
//				//TODO Fix break
//				//for i := 0; i != count; i++ {
//				if node.Method.IsArbiter() {
//					if !executable.ValidateData("sequence") {
//						executable.SetData("sequence", structures.GenerateRandomSequence(len(node.MainChildren)))
//					}
//					if !executable.ValidateData("keys") {
//						executable.SetData("keys", []int{})
//					}
//					if !executable.ValidateData("nodeCount") {
//						executable.SetData("nodeCount", len(node.MainChildren))
//					}
//					node.Method.Call(node.Definition.Signature.Parameters, executable)
//					selection := executable.GetData("selection").(int)
//					RunNode(node.MainChildren[selection], executable)
//				} else {
//					node.Method.Call(node.Definition.Signature.Parameters, executable)
//					for _, child := range node.MainChildren {
//						RunNode(child, executable)
//					}
//				}
//				if node.Method.IsIterator() {
//					RunNode(node, executable)
//				}
//				//}
//			case "conditional callback":
//				// TODO Implement Conditional Callbacks
//				// Logic : Perform Main Child sequentially until condition
//				// Then switch to Alt Child and perform sequentially until condition resets
//			}
//		}
//	}
//}

func Execute(essence *structures.Executable) {
	roots := essence.DependencyGraph.GetRoots()
	for _, root := range roots {
		RunNode(root, essence)
	}
}
