package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 3

type Cell struct {
	isAlive    bool
	neighbours int
}

func main() {
	grid := newGrid()
	fmt.Println(grid)
	updatedGrid := updateAliveNeighbours(*grid)
	final := gameOfLife(updatedGrid)
	fmt.Println(final)
}

// newGrid creates a random matrix of size (size X size).
func newGrid() *[size][size]Cell {

	grid := [size][size]Cell{}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			grid[i][j].isAlive = rand.Float32() > 0.5
		}
	}

	return &grid
}

// Takes a grid and checks if the cells should live or die.
func gameOfLife(r [size][size]Cell) [size][size]Cell {

	c := r
	for i := 0; i < size; i++ {
		for j, val := range r[i] {
			total := val.neighbours
			if total < 2 || total >= 4 {
				val.isAlive = false
			} else if total == 3 {
				val.isAlive = true
			}

			c[i][j] = val
		}
	}

	return c
}

// updateAliveNeighbours updates the neighbours field of the Cell struct with the number of alive neighbours.
func updateAliveNeighbours(r [size][size]Cell) [size][size]Cell {

	c := r
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			c[i][j].neighbours = countAliveNeighbours(i, j, r)
		}
	}

	return c
}

// countAliveNeighbours gives a sum of the total alive neighbours of a cell at position (x,y).
func countAliveNeighbours(x, y int, r [size][size]Cell) int {

	count := 0
	xCordinates := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	yCordinates := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := range xCordinates {
		if getCellMortality(x+xCordinates[i], y+yCordinates[i], r) {
			count++
		}
	}

	return count
}

// getCell checks if the cell is alive or not.
func getCellMortality(x, y int, grid [size][size]Cell) bool {

	if x >= 0 && y >= 0 && x < size && y < size {
		return grid[x][y].isAlive
	}
	return false
}
