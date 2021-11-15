package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(data string) (*Matrix, error) {

	rows := strings.Split(data, "\n")
	var matrix = make(Matrix, len(rows))

	for i, r := range rows {
		nums := strings.Fields(r)
		if i > 0 && len(nums) != len(matrix[0]) {
			return nil, errors.New("rows not equal")
		}

		for _, col := range nums {
			cellData, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}

			matrix[i] = append(matrix[i], cellData)
		}
	}

	return &matrix, nil
}

func (m *Matrix) Cols() [][]int {
	var matrixData = make([][]int, len((*m)[0]))
	for i := range matrixData {
		var nCol = make([]int, len(*m))
		for j := range nCol {
			nCol[j] = (*m)[j][i]
		}
		matrixData[i] = nCol
	}

	return matrixData
}

func (m *Matrix) Rows() [][]int {
	var cols = make([][]int, len(*m))
	for i, col := range *m {
		nCol := []int{}
		nCol = append(nCol, col...)
		cols[i] = nCol
	}
	return cols
}

func (m *Matrix) Set(r, c, d int) bool {
	if r < 0 || c < 0 || r >= len(*m) || c >= len((*m)[0]) {
		return false
	}

	(*m)[r][c] = d
	return true
}
