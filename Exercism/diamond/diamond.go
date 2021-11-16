package diamond

import (
	"bytes"
	"errors"
)

func Gen(b byte) (string, error) {
	if b < 'A' || b > 'Z' {
		return "", errors.New("out of range")
	}
	x := int(b - 'A')
	rows := make([][]byte, 2*x+1)
	for i, j := x, 0; i >= 0; i, j = i-1, j+1 {
		line := bytes.Repeat([]byte{' '}, 2*x+2)
		line[2*x+1] = '\n'
		char := 'A' + byte(j)
		line[i], line[2*x-i] = char, char
		rows[j], rows[2*x-j] = line, line
	}
	return string(bytes.Join(rows, nil)), nil
}
