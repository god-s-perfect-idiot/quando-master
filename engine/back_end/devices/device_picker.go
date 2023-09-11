package devices

import (
	"math/rand"
	"quando/engine/structures"
	"time"
)

func GetPickerCallbacks() []structures.Method {
	return []structures.Method{
		{
			Identifier: "quando.pick.random",
			Function:   PickRandom,
			Type:       "callback",
			Iterator:   false,
			Arbiter:    true,
		},
		{
			Identifier: "quando.pick.every",
			Function:   PickEvery,
			Type:       "callback",
			Iterator:   true,
			Arbiter:    true,
		},
		{
			Identifier: "quando.pick.one",
			Function:   PickOne,
			Type:       "callback",
			Iterator:   false,
			Arbiter:    true,
		},
		{
			Identifier: "quando.pick.val",
			Function:   PickVal,
			Type:       "callback",
			Iterator:   false,
			Arbiter:    true,
		},
	}
}

func PickEvery(params map[string]interface{}, ctx *structures.RunContext) {
	count := params["count"].(int)
	units := params["units"].(string)
	duration := timeMS(count, units)
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
	pickType := params["type"].(string)
	pickRandom(len(ctx.CallNode.MainChildren), pickType, ctx)
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
	inverted := params["inverted"].(bool)
	pickOne(len(ctx.CallNode.MainChildren), inverted, ctx)
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

func PickVal(_ map[string]interface{}, ctx *structures.RunContext) {
	pickVal(len(ctx.CallNode.MainChildren), ctx)
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
