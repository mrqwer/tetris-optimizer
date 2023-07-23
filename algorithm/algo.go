package algorithm

import "fmt"

type Board struct {
	Dim   int
	Cells [][]byte
}

type Tetromino struct {
	Letter        byte
	Shape         []string
	Width, Height int
}

func checkFit(board *Board, tetromino *Tetromino, row, col int) bool {
	for i := 0; i < tetromino.Height; i++ {
		for j := 0; j < tetromino.Width; j++ {
			if row+i >= board.Dim || col+j >= board.Dim {
				return false
			}
			if tetromino.Shape[i][j] != '.' && board.Cells[row+i][col+j] != '.' {
				return false
			}
		}
	}
	return true
}

func placeTetromino(board *Board, tetromino *Tetromino, row, col int) {
	for i := 0; i < tetromino.Height; i++ {
		for j := 0; j < tetromino.Width; j++ {
			if tetromino.Shape[i][j] != '.' {
				board.Cells[row+i][col+j] = tetromino.Letter
			}
		}
	}
}

func removeTetromino(board *Board, tetromino *Tetromino, row, col int) {
	for i := 0; i < tetromino.Height; i++ {
		for j := 0; j < tetromino.Width; j++ {
			if tetromino.Shape[i][j] != '.' {
				board.Cells[row+i][col+j] = '.'
			}
		}
	}
}
func Solve(board *Board, tetrominos []*Tetromino) bool {
	if len(tetrominos) == 0 {
		return true // Base case
	}
	for row := 0; row < board.Dim; row++ {
		for col := 0; col < board.Dim; col++ {
			for i, tetromino := range tetrominos {
				if checkFit(board, tetromino, row, col) {
					placeTetromino(board, tetromino, row, col)
					remainingTetrominos := make([]*Tetromino, len(tetrominos)-1)
					copy(remainingTetrominos, tetrominos[:i])
					copy(remainingTetrominos[i:], tetrominos[i+1:])
					if Solve(board, remainingTetrominos) {
						return true
					}
					removeTetromino(board, tetromino, row, col)
				}
			}
		}
	}
	return false
}

func printProcess(board [][]byte) {
	for _, sl := range board {
		for _, b := range sl {
			fmt.Printf("%v ", string(b))
		}
		fmt.Println()
	}
}
