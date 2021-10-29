package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	newGrid := getGrid()
	fmt.Println(newGrid)

	neighbours := getAliveNeighbours(newGrid)
	final := gameOfLife(newGrid, neighbours)

	fmt.Println(final)
}

// gameOfLife takes a grid and it's neighbours and checks if the cells should live or die.
func gameOfLife(grid, neighbours [][]int) [][]int {

	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if neighbours[i][j] < 2 || neighbours[i][j] >= 4 {
				grid[i][j] = 0
			} else if neighbours[i][j] == 3 {
				grid[i][j] = 1
			}
		}
	}
	return grid
}

// getAliveNeighbours make a new 2D array with the number of alive neighbours for each corresponding cell in grid.
func getAliveNeighbours(grid [][]int) [][]int {

	neighbours := make([][]int, len(grid[0]))
	for i := range neighbours {
		neighbours[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid[0]); j++ {
			neighbours[i][j] = countAliveNeighbours(i, j, grid)
		}
	}

	return neighbours
}

// countAliveNeighbours gives a sum of the total alive neighbours of a cell at position (x,y).
func countAliveNeighbours(x, y int, grid [][]int) int {

	count := 0
	xCordinates := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	yCordinates := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := range xCordinates {
		if getCellMortality(x+xCordinates[i], y+yCordinates[i], grid) {
			count++
		}
	}

	return count
}

// getCell checks if the cell is alive or not.
func getCellMortality(x, y int, grid [][]int) bool {

	if x >= 0 && y >= 0 && x < len(grid[0]) && y < len(grid[0]) {
		if grid[x][y] > 0 {
			return true
		}
	}
	return false
}

func getGrid() [][]int {

	var sizeOfInputMatrix string
	fmt.Println("The input array is a 2D array of size M X N where M is equal to N")
	fmt.Println("Enter the value of M: ")
	fmt.Scanln(&sizeOfInputMatrix)

	M, _ := strconv.Atoi(sizeOfInputMatrix)

	newGrid := make([][]int, M)
	for i := range newGrid {
		newGrid[i] = make([]int, M)
	}

	var input string
	fmt.Printf("Enter the value of the %d cells separated by a comma: ", M*M)
	fmt.Scanln(&input)

	input = strings.ReplaceAll(input, " ", "")
	cellData := strings.Split(input, ",")

	for i := 0; i < M; i++ {
		for j := 0; j < M; j++ {
			newGrid[j][i], _ = strconv.Atoi(cellData[(j*M)+i])
		}
	}

	return newGrid
}
