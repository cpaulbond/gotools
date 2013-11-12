package bitset64

import (
	"testing"
)

func Test_basic_testing(t *testing.T) {
	var s Bitset64

	for i := 0; i < 16; i++ {
		if s.IsSet(i) {
			t.Errorf("Unexpected bit %d set\n", i)
			return
		}

		s.Set(i)

		if !s.IsSet(i) {
			t.Errorf("Unexpected bit %d not set\n", i)
			return
		}

		s.Clear(i)

		if s.IsSet(i) {
			t.Errorf("Unexpected bit %d set\n", i)
			return
		}
	}
}
