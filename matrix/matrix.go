package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix represent a matrix
type Matrix [][]int

// Rows return a copy of the rows of the matrix
func (m Matrix) Rows() [][]int {
	out := make(Matrix, len(m))
	for i, c := range m {
		out[i] = make([]int, len(c))
		copy(out[i], c)
	}
	return out
}

// Cols return a copy of the rows of the matrix
func (m Matrix) Cols() [][]int {
	xl := len(m[0])
	yl := len(m)
	result := make([][]int, xl)
	for i := range result {
		result[i] = make([]int, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = m[j][i]
		}
	}
	return result
}

// Set will modify a value of the matrix
func (m *Matrix) Set(i int, j int, val int) bool {
	if i < 0 || j < 0 || i >= len(*m) || j >= len((*m)[i]) {
		return false
	}
	(*m)[i][j] = val

	return true
}

// New return a matrix for a given input string
func New(input string) (*Matrix, error) {
	var col Matrix
	rowLen := -1
	for _, j := range strings.Split(input, "\n") {
		var row []int
		for _, x := range strings.Split(strings.TrimSpace(j), " ") {
			el, err := strconv.Atoi(x)
			if err != nil {
				return nil, err
			}
			row = append(row, el)
		}
		if rowLen == -1 {
			rowLen = len(row)
		} else {
			if rowLen != len(row) {
				return nil, errors.New("invalid row len")
			}
		}
		col = append(col, row)
	}
	return &col, nil
}
