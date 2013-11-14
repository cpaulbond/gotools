package main

import (
	".."

	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rows := flag.Int("r", 5, "Number of rows.")
	cols := flag.Int("c", 5, "Number of columns.")
	seed := flag.Int64("s", 0, "If the specified seed is less then 1 the time will be used as a seed.")
	dump := flag.Bool("d", false, "Dump maze data.")
	flag.Parse()

	if *rows < 5 {
		fmt.Printf("maze: Selected number of rows is %d. Mimumum number of rows is %d", *rows, 5)
		os.Exit(1)
	}

	if *cols < 5 {
		fmt.Printf("maze: Selected number of cols is %d. Mimumum number of cols is %d", *cols, 5)
		os.Exit(1)
	}

	if *seed < 1 {
		rand.Seed(time.Now().UTC().UnixNano())
	} else {
		rand.Seed(*seed)
	}

	m := maze.NewMaze(*rows, *cols)
	m.Carve()
	if *dump {
		m.Dump()
	}
	m.Print()
}
