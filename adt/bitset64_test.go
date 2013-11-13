package adt

import (
	"testing"
)

func Test_basic_testing(t *testing.T) {
	var s Bitset64

	s.Set(15)
	if s.String() != "{1000000000000000}" {
		t.Errorf("Unexpected output string %s\n", s.String())
		return
	}
	s.Init()

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
