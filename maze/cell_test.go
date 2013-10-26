package maze

import (
	"testing"
)

func test_cell_decode(v, s string, t *testing.T) {
	if v != s {
		t.Errorf("Expected cell value %s found %s", v, s)
	}
}

func Test_cell(t *testing.T) {
	c := cell{}

	test_cell_decode("......", c.decode(), t)
	c.edge = true
	test_cell_decode("e.....", c.decode(), t)
	c.mark = true
	test_cell_decode("e....X", c.decode(), t)
}
