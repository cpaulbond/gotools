package main

import (
	".."
)

func main() {
	m := maze.NewMaze(5, 5)
	m.Dump()
	m.Print()
}
