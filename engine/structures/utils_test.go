package structures

import "testing"

func TestGenerateRandomSequence(t *testing.T) {
	sequence := GenerateRandomSequence(10)
	if len(sequence) != 10 {
		t.Error("GenerateRandomSequence should return sequence with length 10")
	}
	for _, v := range sequence {
		if v < 0 || v > 9 {
			t.Error("GenerateRandomSequence should return sequence with values between 0 and 9")
		}
	}
}
