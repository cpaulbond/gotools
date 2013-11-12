package main

import (
	".."

	"math/rand"
	"time"
)

func main() {
	// TODO: When done turn on seeding the random number generator.
	if false {
		rand.Seed(time.Now().UTC().UnixNano())
	}

	m := maze.NewMaze(5, 5)
	m.Carve()
	m.Dump()
	m.Print()
}
