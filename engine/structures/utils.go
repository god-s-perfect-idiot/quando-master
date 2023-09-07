package structures

import (
	"github.com/shomali11/util/xhashes"
	"math/rand"
)

func GenerateRandomSequence(count int) []int {
	var sequence []int
	for {
		random := rand.Intn(count)
		unique := true
		for _, v := range sequence {
			if v == random {
				unique = false
				break
			}
		}
		if unique {
			sequence = append(sequence, random)
		}
		if len(sequence) == count {
			break
		}
	}
	return sequence
}

func GetHash(sourceCode string) string {
	return xhashes.SHA256(sourceCode)
}
