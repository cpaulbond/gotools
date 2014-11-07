package maze

import (
	"fmt"

	"github.com/cpaulbond/gotools/adt"
)

type Maze struct {
	rows int
	cols int
	data [][]adt.Bitset64
}

const (
	edge = iota
	top
	right
	visited
	solution
	mark
)

func (m *Maze) set(r, c int, bit int) {
	m.data[r][c].Set(bit)
}

func (m *Maze) clear(r, c int, bit int) {
	m.data[r][c].Clear(bit)
}

func (m *Maze) is_set(r, c int, bit int) bool {
	return m.data[r][c].IsSet(bit)
}

func (m *Maze) decode(r, c int) string {
	chr := []byte{'e', 't', 'r', 'v', 's', 'X'}

	val := make([]byte, len(chr))

	for i := 0; i <= mark; i++ {
		if m.data[r][c].IsSet(i) {
			val[i] = chr[i]
		} else {
			val[i] = '.'
		}
	}

	return string(val)
}

func NewMaze(rows, cols int) (rtn *Maze) {
	rtn = &Maze{
		rows: rows + 2,
		cols: cols + 2,
		data: nil,
	}

	rtn.Init()
	return rtn
}

func (m *Maze) Init() {
	m.data = make([][]adt.Bitset64, m.rows)
	for r := 0; r < m.rows; r++ {
		m.data[r] = make([]adt.Bitset64, m.cols)
	}

	// Mark top and bottom edges.
	r := m.rows - 1
	for c := 0; c < m.cols; c++ {
		m.set(0, c, edge)
		m.set(r, c, edge)
	}

	// Mark left and right edges.
	c := m.cols - 1
	for r := 0; r < m.rows; r++ {
		m.set(r, 0, edge)
		m.set(r, c, edge)
	}

	// Mark top walls.
	for r := 1; r < m.rows; r++ {
		for c := 1; c < (m.cols - 1); c++ {
			m.set(r, c, top)
		}
	}

	// Mark right walls.
	for r := 1; r < (m.rows - 1); r++ {
		for c := 0; c < (m.cols - 1); c++ {
			m.set(r, c, right)
		}
	}
}

func (m *Maze) Dump() {
	fmt.Printf("Raw rows: %d, cols: %d\n", m.rows, m.cols)

	fmt.Print("     ")
	for c := 0; c < m.cols; c++ {
		fmt.Printf("%6d ", c)
	}
	fmt.Print("\n")

	for r := 0; r < m.cols; r++ {
		fmt.Printf("%4d ", r)

		for c := 0; c < m.cols; c++ {
			fmt.Print(m.decode(r, c))
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func (m *Maze) Print() {
	fmt.Printf("SIZE %d x %d\n", m.rows-2, m.cols-2)

	for r := 1; r < m.rows; r++ {
		// Check for top walls.
		for c := 0; c < m.cols-1; c++ {
			if m.is_set(r, c, top) {
				fmt.Print("--+")
			} else {
				fmt.Print("  +")
			}
		}
		fmt.Print("\n")

		// Check right walls.
		for c := 0; c < m.cols-1; c++ {
			if m.is_set(r, c, right) {
				fmt.Print("  |")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Print("\n")
	}
}

func (m *Maze) Carve() {
	m.clear(1, 1, top)
	m.set(1, 1, visited)
	m.clear(m.rows-1, m.cols-2, top)

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

			pt = backup.get_random()
			moves = m.get_moves(pt)
		}

		move := moves.get_random()
		pt = m.carve_passage(pt, move)
		m.set(pt.r, pt.c, visited)
	}
}

func (m *Maze) carve_passage(pt, move point) (new_pt point) {
	new_pt = point{r: pt.r + move.r, c: pt.c + move.c}
	if move.r == -1 && move.c == 0 {
		m.clear(pt.r, pt.c, top)
	} else if move.r == 1 && move.c == 0 {
		m.clear(new_pt.r, new_pt.c, top)
	} else if move.r == 0 && move.c == -1 {
		m.clear(new_pt.r, new_pt.c, right)
	} else if move.r == 0 && move.c == 1 {
		m.clear(pt.r, pt.c, right)
	}
	return
}

var moves = [4]point{
	{r: -1, c: 0},
	{r: 1, c: 0},
	{r: 0, c: -1},
	{r: 0, c: 1},
}

func (m *Maze) get_moves(pt point) (rtn *pointset) {
	rtn = new_pointset()

	for i := range moves {
		if !m.is_ignored(pt.move(moves[i])) {
			rtn.add(moves[i])
		}
	}

	return
}

func (m *Maze) is_ignored(pt point) bool {
	return m.is_set(pt.r, pt.c, edge) || m.is_set(pt.r, pt.c, visited)
}
