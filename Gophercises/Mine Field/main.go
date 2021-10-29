package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	var filename string
	fmt.Println("Enter the filename of the minefield: ")
	fmt.Scanln(&filename)

	content := readFile(filename)
	mineGrid := formatFileContent(content)

	hints := getAliveNeighbours(mineGrid)
	for i := 0; i < len(hints); i++ {
		fmt.Println(hints[i])
	}

}

func readFile(filename string) string {

	content, _ := os.ReadFile(filename)
	fileData := string(content)

	return fileData
}

func formatFileContent(content string) [][]string {

	content = strings.ReplaceAll(content, "\n", "")
	content = strings.ReplaceAll(content, " ", "")
	contentArray := strings.Fields(content)
	content = strings.Join(contentArray, "")
	contentArray = strings.Split(content, "")

	mineSize := int(math.Sqrt(float64(len(content))))

	mineFeild := make([][]string, mineSize)
	for i := range mineFeild {
		mineFeild[i] = make([]string, mineSize)
	}

	for i := 0; i < mineSize; i++ {
		for j := 0; j < mineSize; j++ {
			mineFeild[j][i] = contentArray[(j*mineSize)+i]
		}
	}

	return mineFeild
}

func getAliveNeighbours(grid [][]string) [][]string {

	neighbours := make([][]string, len(grid))
	for i := range neighbours {
		neighbours[i] = make([]string, len(grid))
	}

	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid[0]); j++ {
			neighbours[i][j] = countAliveNeighbours(i, j, grid)
		}
	}

	return neighbours
}

func countAliveNeighbours(x, y int, grid [][]string) string {

	count := 0
	xCordinates := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	yCordinates := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := range xCordinates {
		if getCellMortality(x+xCordinates[i], y+yCordinates[i], grid) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func getCellMortality(x, y int, grid [][]string) bool {

	if x >= 0 && y >= 0 && x < len(grid) && y < len(grid) {
		if grid[x][y] == "*" {
			return true
		}
	}
	return false
}
