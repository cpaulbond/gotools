package maze

import (
	"testing"
)

func check_pointset_len(set *pointset, expected int, t *testing.T) {
	if set.len() != expected {
		t.Errorf("The expected value of set.len() was %d we found %d", expected, set.len())
	}
}

func Test_pointset(t *testing.T) {
	set := new_pointset()
	check_pointset_len(set, 0, t)
	set.add(point{r: 1, c: 1})
	check_pointset_len(set, 1, t)
	set.add(point{r: 1, c: 1})
	check_pointset_len(set, 1, t)
}
