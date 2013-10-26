package maze

import (
	"fmt"
)

type Maze struct {
	rows int
	cols int
	data [][]cell
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
			fmt.Print(m.data[r][c].decode())
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
			if m.data[r][c].top {
				fmt.Print("--+")
			} else {
				fmt.Print("  +")
			}
		}
		fmt.Print("\n")

		// Check right walls.
		for c := 0; c < m.cols-1; c++ {
			if m.data[r][c].right {
				fmt.Print("  |")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Print("\n")
	}
}
