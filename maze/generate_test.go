package maze

import (
	"testing"
)

func Test_get_moves(t *testing.T) {
	m := NewMaze(5, 5)

	check_pointset_len(m.get_moves(point{r: 1, c: 1}), 2, t)
	check_pointset_len(m.get_moves(point{r: 2, c: 2}), 4, t)

	m.data[1][2].visited = true
	check_pointset_len(m.get_moves(point{r: 1, c: 1}), 1, t)

	m.data[2][1].visited = true
	check_pointset_len(m.get_moves(point{r: 1, c: 1}), 0, t)
}
