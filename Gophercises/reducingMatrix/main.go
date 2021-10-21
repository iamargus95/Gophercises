package main

import "fmt"

func main() {
	x := matrixThis(2)
	for _, v := range x {
		fmt.Printf("%v\n", v)
	}
}

func matrixThis(N int) [][]int {

	size := 2*N - 1
	begin := 0
	end := size - 1

	output := make([][]int, size)
	for i := range output {
		output[i] = make([]int, size)
	}

	for N != 0 {
		for i := begin; i <= end; i++ {
			for j := begin; j <= end; j++ {
				if i == begin || i == end || j == begin || j == end {
					output[i][j] = N
				}
			}
		}
		begin++
		end--
		N--
	}
	return output
}
