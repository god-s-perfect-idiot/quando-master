package devices

import (
	"math/rand"
	"quando/engine/structures"
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

func PickEvery(params map[string]interface{}) (float64, map[string]interface{}) {
	nodeCount := params["nodeCount"].(int)
	keys := params["keys"].([]int)
	count := params["count"].(int)
	units := params["units"].(string)
	duration := timeMS(count, units)
	// Is this a good idea?
	// pioneered Inter-device communication
	after(duration)
	selection := pickNext(nodeCount, keys)
	keys = []int{selection}
	data := make(map[string]interface{})
	data["keys"] = keys
	data["selection"] = selection
	return -1.0, data
}

func PickRandom(params map[string]interface{}) (float64, map[string]interface{}) {
	nodeCount := params["nodeCount"].(int)
	pickType := params["type"].(string)
	keys := params["keys"].([]int)
	var selection int
	switch pickType {
	case "random":
		selection = random(nodeCount)
	case "sequence":
		sequenceList := params["sequence"].([]int)
		selection = sequence(sequenceList, keys)
		keys = append(keys, selection)
		if len(keys) == len(sequenceList) {
			keys = []int{}
		}
	case "unique":
		selection = unique(nodeCount, keys)
		keys = append(keys, selection)
		if len(keys) == nodeCount {
			keys = []int{}
		}
	}
	data := make(map[string]interface{})
	data["keys"] = keys
	data["selection"] = selection
	return -1.0, data
}

func PickOne(params map[string]interface{}) (float64, map[string]interface{}) {
	nodeCount := params["nodeCount"].(int)
	keys := params["keys"].([]int)
	inverted := params["inverted"].(bool)
	var selection int
	if inverted {
		selection = pickPrevious(nodeCount, keys)
	} else {
		selection = pickNext(nodeCount, keys)
	}
	keys = []int{selection}
	data := make(map[string]interface{})
	data["keys"] = keys
	data["selection"] = selection
	return -1.0, data
}

func PickVal(params map[string]interface{}) (float64, map[string]interface{}) {
	nodeCount := params["nodeCount"].(int)
	val := params["val"].(float64)
	selection := valNext(val, nodeCount)
	data := make(map[string]interface{})
	data["selection"] = selection
	return -1.0, data
}

func pickPrevious(nodeCount int, keys []int) int {
	if len(keys) == 0 {
		return nodeCount - 1
	}
	index := keys[0] - 1
	if index < 0 {
		index = nodeCount - 1
	}
	return index
}

func valNext(val float64, nodeCount int) int {
	sectionIndex := int(val * float64(nodeCount))
	if sectionIndex >= nodeCount {
		sectionIndex = nodeCount - 1
	}
	return sectionIndex
}

func pickNext(nodeCount int, keys []int) int {
	if len(keys) == 0 {
		return 0
	}
	return (keys[0] + 1) % nodeCount
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
