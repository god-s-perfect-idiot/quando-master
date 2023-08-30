package devices

import "testing"

func TestRandom(t *testing.T) {
	params := make(map[string]interface{})
	params["type"] = "random"
	params["sequence"] = []int{1, 2, 3, 4, 5}
	params["keys"] = []int{1}
	params["nodeCount"] = 5
	val, data := PickRandom(params)
	if val != 0.0 {
		t.Error("random() should return 0.0")
	}
	if data == nil {
		t.Error("random() should return data")
	}
	if data["selection"] == nil {
		t.Error("random() should return selection")
	}
	if data["keys"] == nil {
		t.Error("random() should return keys")
	}
	if data["selection"].(int) < 0 || data["selection"].(int) > 4 {
		t.Error("random() should return an index between 0 and 4")
	}
}

func TestRandomPickRandom(t *testing.T) {
	nodeCount := 5
	for i := 0; i < 100; i++ {
		index := random(nodeCount)
		if index < 0 || index > 4 {
			t.Error("random() should return an index between 0 and 4")
		}
	}
}

func TestRandomPickSequence(t *testing.T) {
	sequenceList := []int{1, 2, 3, 4, 5}
	keys := []int{1}
	for i := 0; i < 100; i++ {
		index := sequence(sequenceList, keys)
		var currentIndex int
		for i, v := range sequenceList {
			if v == keys[len(keys)-1] {
				currentIndex = i
				break
			}
		}
		if index != sequenceList[(currentIndex+1)%len(sequenceList)] {
			t.Error("sequence() should return next index in sequence")
		}
		keys = append(keys, index)
	}
}

func TestRandomPickUnique(t *testing.T) {
	nodeCount := 5
	keys := []int{1}
	for i := 0; i < 100; i++ {
		index := unique(nodeCount, keys)
		unique := true
		for _, v := range keys {
			if v == index {
				unique = false
				break
			}
		}
		if !unique {
			t.Error("unique() should return an index not in keys")
		}
		keys = append(keys, index)
		if len(keys) == nodeCount {
			keys = []int{}
		}
	}
}

func TestPickEvery(t *testing.T) {
	params := make(map[string]interface{})
	params["keys"] = []int{1}
	params["nodeCount"] = 5
	params["units"] = "seconds"
	params["count"] = 1
	val, data := PickEvery(params)
	if val != 0.0 {
		t.Error("pickEvery() should return 0.0")
	}
	if data == nil {
		t.Error("pickEvery() should return data")
	}
	if data["selection"] == nil {
		t.Error("pickEvery() should return selection")
	}
	if data["keys"] == nil {
		t.Error("pickEvery() should return keys")
	}
}

func TestPickNextMethod(t *testing.T) {
	nodeCount := 5
	keys := []int{1}
	for i := 0; i < 100; i++ {
		index := pickNext(nodeCount, keys)
		if index != (keys[0]+1)%nodeCount {
			t.Error("pickEvery() should return next index in sequence")
		}
		keys = append(keys, index)
	}
}

func TestOneNext(t *testing.T) {
	params := make(map[string]interface{})
	params["keys"] = []int{1}
	params["nodeCount"] = 5
	params["inverted"] = false
	val, data := PickOne(params)
	if val != 0.0 {
		t.Error("next() should return 0.0")
	}
	if data == nil {
		t.Error("next() should return data")
	}
	if data["selection"] == nil {
		t.Error("next() should return selection")
	}
	if data["keys"] == nil {
		t.Error("next() should return keys")
	}
	if data["selection"].(int) < 0 || data["selection"].(int) > 4 {
		t.Error("next() should return an index between 0 and 4")
	}
}

func TestPickOneInverted(t *testing.T) {
	params := make(map[string]interface{})
	params["keys"] = []int{1}
	params["nodeCount"] = 5
	params["inverted"] = true
	val, data := PickOne(params)
	if val != 0.0 {
		t.Error("next() should return 0.0")
	}
	if data == nil {
		t.Error("next() should return data")
	}
	if data["selection"] == nil {
		t.Error("next() should return selection")
	}
	if data["keys"] == nil {
		t.Error("next() should return keys")
	}
	if data["selection"].(int) < 0 || data["selection"].(int) > 4 {
		t.Error("next() should return an index between 0 and 4")
	}
}

func TestPickPreviousMethod(t *testing.T) {
	nodeCount := 5
	keys := []int{1}
	for i := 0; i < 100; i++ {
		index := pickPrevious(nodeCount, keys)
		if index != keys[0]-1 {
			t.Error("pickPrevious() should return previous index in sequence")
		}
		keys = append(keys, index)
	}
}

func TestPickVal(t *testing.T) {
	params := make(map[string]interface{})
	params["nodeCount"] = 5
	params["val"] = 0.0
	val, data := PickVal(params)
	if val != 0.0 {
		t.Error("pickVal() should return 0.0")
	}
	if data == nil {
		t.Error("pickVal() should return data")
	}
	if data["selection"] == nil {
		t.Error("pickVal() should return selection")
	}
	if data["selection"].(int) < 0 || data["selection"].(int) > 4 {
		t.Error("pickVal() should return an index between 0 and 4")
	}
}

func TestValNext(t *testing.T) {
	nodeCount := 5
	for i := 0; i < 100; i++ {
		index := valNext(float64(i)/100.0, nodeCount)
		if i < 20 && index != 0 {
			t.Error("valNext() should return 0 for values < 0.2")
		}
		if i >= 20 && i < 40 && index != 1 {
			t.Error("valNext() should return 1 for values >= 0.2 and < 0.4")
		}
		if i >= 40 && i < 60 && index != 2 {
			t.Error("valNext() should return 2 for values >= 0.4 and < 0.6")
		}
		if i >= 60 && i < 80 && index != 3 {
			t.Error("valNext() should return 3 for values >= 0.6 and < 0.8")
		}
		if i >= 80 && index != 4 {
			t.Error("valNext() should return 4 for values >= 0.8")
		}
	}
}
