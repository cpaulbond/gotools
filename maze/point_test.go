package maze

import (
	"math/rand"
	"testing"
	"time"
)

func check_pointset_len(set *pointset, expected int, t *testing.T) {
	if set.len() != expected {
		t.Errorf("The expected value of set.len() was %d we found %d", expected, set.len())
	}
}

func Test_pointset_len(t *testing.T) {
	set := new_pointset()
	check_pointset_len(set, 0, t)
	set.add(point{r: 1, c: 1})
	check_pointset_len(set, 1, t)
	set.add(point{r: 1, c: 1})
	check_pointset_len(set, 1, t)
}

func Test_pointset_get(t *testing.T) {
	set := new_pointset()
	for i := 0; i < 10; i++ {
		set.add(point{r: i, c: i})
	}

	check_pointset_len(set, 10, t)

	for i := 0; i < 10; i++ {
		set.get()
		//t.Log(pt)
	}

	check_pointset_len(set, 0, t)
}

func Test_pointset_get_random(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	set := new_pointset()

	for i := 0; i < 10; i++ {
		set.add(point{r: i, c: i})
	}

	check_pointset_len(set, 10, t)

	for i := 0; i < 10; i++ {
		pt := set.get_random()
		if pt.r == -1 {
			t.Error("Found invalid point") // Just for testing...
		}
		//t.Log(pt)
	}

	check_pointset_len(set, 0, t)
}
