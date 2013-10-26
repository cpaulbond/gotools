package maze

type cell struct {
	edge,
	top,
	right,
	visited,
	solution,
	mark bool
}

func (c *cell) decode() string {
	set := func(val []byte, i int, on bool, v byte) {
		if on {
			val[i] = v
		} else {
			val[i] = '.'
		}
	}

	val := make([]byte, 6)

	set(val, 0, c.edge, 'e')
	set(val, 1, c.top, 't')
	set(val, 2, c.right, 'r')
	set(val, 3, c.visited, 'v')
	set(val, 4, c.solution, 's')
	set(val, 5, c.mark, 'X')

	return string(val)
}
