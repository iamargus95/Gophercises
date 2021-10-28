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
	updatedGrid := updateAliveNeighbours(*grid)
	final := gameOfLife(updatedGrid)
	fmt.Println(final)

}

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

// getCell checks the value of one neighbour depending on x and y
func getCellMortality(x, y int, grid [size][size]Cell) bool {
	if x >= 0 && y >= 0 && x < size && y < size {
		return grid[x][y].isAlive
	}
	return false
}

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

func updateAliveNeighbours(r [size][size]Cell) [size][size]Cell {

	c := r
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			c[i][j].neighbours = countAliveNeighbours(i, j, r)
		}
	}

	return c
}

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
