package algorithm

import (
	"testing"
)

func TestSolve(t *testing.T) {
	bigMap := make(map[*Board][]*Tetromino)
	board1 := &Board{
		Dim: 4,
		Cells: [][]byte{
			[]byte("...."),
			[]byte("...."),
			[]byte("...."),
			[]byte("...."),
		},
	}

	tetrominos1 := []*Tetromino{
		&Tetromino{
			Letter: 'A',
			Shape: []string{
				"..#.",
				"..#.",
				"..#.",
				"..#.",
			},
			Width:  1,
			Height: 4,
		},
		&Tetromino{
			Letter: 'B',
			Shape: []string{
				"###.",
				"..#.",
				"....",
				"....",
			},
			Width:  2,
			Height: 2,
		},
	}
	bigMap[board1] = tetrominos1

	for board, tetrominos := range bigMap {
		if !Solve(board, tetrominos, 0) {
			t.Error("Failed to solve the puzzle with the provided Tetrominos.")
		}
	}
}
