package maze

func (m *Maze) clear() {
	m.data = make([][]cell, m.rows)
	for r := 0; r < m.rows; r++ {
		m.data[r] = make([]cell, m.cols)
	}

	// Mark top and bottom edges.
	r := m.rows - 1
	for c := 0; c < m.cols; c++ {
		m.data[0][c].edge = true
		m.data[r][c].edge = true
	}

	// Mark left and right edges.
	c := m.cols - 1
	for r := 0; r < m.rows; r++ {
		m.data[r][0].edge = true
		m.data[r][c].edge = true
	}

	// Mark top walls.
	for r := 1; r < m.rows; r++ {
		for c := 1; c < (m.cols - 1); c++ {
			m.data[r][c].top = true
		}
	}

	// Mark right walls.
	for r := 1; r < (m.rows - 1); r++ {
		for c := 0; c < (m.cols - 1); c++ {
			m.data[r][c].right = true
		}
	}
}

func (m *Maze) Carve() {

	m.data[1][1].top = false
	m.data[1][1].visited = true
	m.data[m.rows-1][m.cols-2].top = false

	backup := new_pointset()
	pt := point{r: 1, c: 1}
	run_limit := ((m.rows + m.cols) / 3) + 4
	run := 0

	for i := 0; i < ((m.rows-2)*(m.cols-2))-1; i++ {
		moves := m.get_moves(pt)
		if moves.len() > 1 {
			backup.add(pt)
		}

		run += 1
		for (run > run_limit) || (moves.len() == 0) {
			if backup.len() == 0 {
				panic("No backup!?")
			}
			run = 0

			pt = backup.random()
			moves = m.get_moves(pt)
		}

		//       (let ((move (maze-select-random moves)))
		//         (setq pt (maze-open-passage pt move maze))
		//         (maze-set-flag pt :visited maze))
	}
}

func (m *Maze) get_moves(pt point) (rtn *pointset) {
	moves := [4]point{
		{r: -1, c: 0},
		{r: 1, c: 0},
		{r: 0, c: -1},
		{r: 0, c: 1},
	}

	rtn = new_pointset()

	for i := range moves {
		if !m.is_ignored(pt.move(moves[i])) {
			rtn.add(moves[i])
		}
	}

	return
}

func (m *Maze) is_ignored(pt point) bool {
	return m.data[pt.r][pt.c].edge || m.data[pt.r][pt.c].visited
}
