package str

import (
	"testing"
)

func TestRandom(t *testing.T) {
	for length := 0; length < 20; length++ {
		str := []rune(Random(length))
		if len(str) == length {
			continue
		}
		t.Fatalf("Random(length=%v)", length)
	}
}
