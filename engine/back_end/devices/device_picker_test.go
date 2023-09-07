package devices

import (
	"testing"
)

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

func TestPickNextMethod(t *testing.T) {
	nodeCount := 5
	previous := 0
	for i := 0; i < 100; i++ {
		index := pickNext(nodeCount, previous)
		if index != (previous+1)%nodeCount {
			t.Error("pickNext() should return next index in sequence")
		}
		previous = index
	}
}

func TestPickPreviousMethod(t *testing.T) {
	nodeCount := 100
	previous := 100
	for i := 0; i < 100; i++ {
		index := pickPrevious(nodeCount, previous)
		if index != (previous-1)%nodeCount {
			t.Error("pickPrevious() should return previous index in sequence")
		}
		previous = index
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
