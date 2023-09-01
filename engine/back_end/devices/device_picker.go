package devices

import (
	"math/rand"
	"quando/engine/structures"
	"time"
)

func GetPickerCallbacks() []structures.Method {
	return []structures.Method{
		structures.Method{
			Identifier: "quando.pick.random",
			Function:   PickRandom,
			Type:       "callback",
			Iterator:   false,
			Arbiter:    true,
		},
		structures.Method{
			Identifier: "quando.pick.every",
			Function:   PickEvery,
			Type:       "callback",
			Iterator:   true,
			Arbiter:    true,
		},
		structures.Method{
			Identifier: "quando.pick.one",
			Function:   PickOne,
			Type:       "callback",
			Iterator:   false,
			Arbiter:    true,
		},
		structures.Method{
			Identifier: "quando.pick.val",
			Function:   PickVal,
			Type:       "callback",
			Iterator:   false,
			Arbiter:    true,
		},
	}
}

func PickEvery(params map[string]interface{}, ctx *structures.RunContext) {
	//nodeCount := params["nodeCount"].(int)
	//keys := params["keys"].([]int)
	count := params["count"].(int)
	units := params["units"].(string)
	duration := timeMS(count, units)
	// Is this a good idea?
	// pioneered Inter-device communication
	//after(duration)
	//selection := pickNext(nodeCount, keys)
	//keys = []int{selection}
	//ctx.Executable.SetData("keys", keys)
	//ctx.Executable.SetData("selection", selection)
	go pickEvery(duration, len(ctx.CallNode.MainChildren), ctx)
}

func pickEvery(duration int, nodeCount int, ctx *structures.RunContext) {
	previous := 0
	for {
		wait(duration)
		selection := pickNext(nodeCount, previous)
		previous = selection
		lucky := ctx.CallNode.MainChildren[selection]
		lucky.Method.CallFunc(ctx.Executable, lucky)
	}
}

func PickRandom(params map[string]interface{}, ctx *structures.RunContext) {
	//nodeCount := params["nodeCount"].(int)
	pickType := params["type"].(string)
	//keys := params["keys"].([]int)
	pickRandom(len(ctx.CallNode.MainChildren), pickType, ctx)
	//var selection int
	//switch pickType {
	//case "random":
	//	selection = random(nodeCount)
	//case "sequence":
	//	sequenceList := params["sequence"].([]int)
	//	selection = sequence(sequenceList, keys)
	//	keys = append(keys, selection)
	//	if len(keys) == len(sequenceList) {
	//		keys = []int{}
	//	}
	//case "unique":
	//	selection = unique(nodeCount, keys)
	//	keys = append(keys, selection)
	//	if len(keys) == nodeCount {
	//		keys = []int{}
	//	}
	//}
	//data := make(map[string]interface{})
	//data["keys"] = keys
	//data["selection"] = selection
	//return -1.0, data
}

func pickRandom(nodeCount int, pickType string, ctx *structures.RunContext) {
	var selection int
	if ctx.Executable.GetData("random-sequence") == nil {
		ctx.Executable.SetData("random-sequence", structures.GenerateRandomSequence(nodeCount))
	}
	if ctx.Executable.GetData("random-keys") == nil {
		ctx.Executable.SetData("random-keys", []int{})
	}
	keys := ctx.Executable.GetData("random-keys").([]int)
	switch pickType {
	case "random":
		selection = random(nodeCount)
	case "sequence":
		sequenceList := ctx.Executable.GetData("random-sequence").([]int)
		selection = sequence(sequenceList, keys)
		keys = append(keys, selection)
		if len(keys) == nodeCount {
			keys = []int{}
		}
	case "unique":
		selection = unique(nodeCount, keys)
		keys = append(keys, selection)
		if len(keys) == nodeCount {
			keys = []int{}
		}
	}
	ctx.Executable.SetData("random-keys", keys)
	lucky := ctx.CallNode.MainChildren[selection]
	lucky.Method.CallFunc(ctx.Executable, lucky)
}

func PickOne(params map[string]interface{}, ctx *structures.RunContext) {
	//nodeCount := params["nodeCount"].(int)
	//keys := params["keys"].([]int)
	inverted := params["inverted"].(bool)
	pickOne(len(ctx.CallNode.MainChildren), inverted, ctx)
	//var selection int
	//if inverted {
	//	selection = pickPrevious(len(ctx.CallNode.MainChildren), keys)
	//} else {
	//	selection = pickNext(nodeCount, keys)
	//}
	//keys = []int{selection}
	//data := make(map[string]interface{})
	//data["keys"] = keys
	//data["selection"] = selection
}

func pickOne(nodeCount int, inverted bool, ctx *structures.RunContext) {
	var selection int
	if ctx.Executable.GetData("one-lastKey") == nil {
		if inverted {
			ctx.Executable.SetData("one-lastKey", 0)
		} else {
			ctx.Executable.SetData("one-lastKey", nodeCount-1)
		}
	}
	previous := ctx.Executable.GetData("one-lastKey").(int)
	if inverted {
		selection = pickPrevious(nodeCount, previous)
	} else {
		selection = pickNext(nodeCount, previous)
	}
	ctx.Executable.SetData("one-lastKey", selection)
	lucky := ctx.CallNode.MainChildren[selection]
	lucky.Method.CallFunc(ctx.Executable, lucky)
}

func PickVal(params map[string]interface{}, ctx *structures.RunContext) {
	//nodeCount := params["nodeCount"].(int)
	//val := params["val"].(float64)
	pickVal(len(ctx.CallNode.MainChildren), ctx)
	//selection := valNext(val, nodeCount)
	//data := make(map[string]interface{})
	//data["selection"] = selection
	//return -1.0, data
}

func pickVal(nodeCount int, ctx *structures.RunContext) {
	val := ctx.Executable.Val
	selection := valNext(val, nodeCount)
	ctx.Executable.Val = val
	lucky := ctx.CallNode.MainChildren[selection]
	lucky.Method.CallFunc(ctx.Executable, lucky)
}

func pickPrevious(nodeCount int, previous int) int {
	selection := previous - 1
	if selection < 0 {
		selection = nodeCount - 1
	}
	return selection
}

//func pickPrevious(nodeCount int, previous int) int {
//	if len(keys) == 0 {
//		return nodeCount - 1
//	}
//	index := keys[0] - 1
//	if index < 0 {
//		index = nodeCount - 1
//	}
//	return index
//}

func valNext(val float64, nodeCount int) int {
	sectionIndex := int(val * float64(nodeCount))
	if sectionIndex >= nodeCount {
		sectionIndex = nodeCount - 1
	}
	return sectionIndex
}

func pickNext(nodeCount int, previous int) int {
	return (previous + 1) % nodeCount
}

//func pickNext(nodeCount int, keys []int) int {
//	if len(keys) == 0 {
//		return 0
//	}
//	return (keys[0] + 1) % nodeCount
//}

func unique(nodeCount int, keys []int) int {
	var newKey int
	for {
		newKey = rand.Intn(nodeCount)
		unique := true
		for _, v := range keys {
			if v == newKey {
				unique = false
				break
			}
		}
		if unique {
			break
		}
	}
	return newKey
}

func sequence(sequence []int, keys []int) int {
	var keyIndex int
	if len(keys) == 0 {
		return sequence[0]
	}
	for i, v := range sequence {
		if v == keys[len(keys)-1] {
			keyIndex = i
			break
		}
	}
	// kinda redundant, but it's for the best
	return sequence[(keyIndex+1)%len(sequence)]
}

func random(nodeCount int) int {
	return rand.Intn(nodeCount)
}

func wait(count int) {
	time.Sleep(time.Duration(count) * time.Millisecond)
}
