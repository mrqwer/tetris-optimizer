package program

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"
	"tetris-optimizer/algorithm"
	"tetris-optimizer/file"
	"tetris-optimizer/helper"
)

func newTetromino(s string, letter byte) *algorithm.Tetromino {
	lines := strings.Split(s, "\n")
	lines = shiftPiece(lines)
	fmt.Printf("Height: %v\n", len(lines))
	fmt.Printf("Width: %v\n", len(lines[0]))
	return &algorithm.Tetromino{
		Letter: letter,
		Shape:  lines,
		Width:  len(lines[0]),
		Height: len(lines),
	}
}

func generateBoard(dim int) [][]byte {
	board_form := make([][]byte, dim)

	for i := 0; i < dim; i++ {
		board_form[i] = make([]byte, dim)
		for j := 0; j < dim; j++ {
			board_form[i][j] = '.'
		}
	}
	return board_form
}

func shiftPiece(grid []string) []string {
	minRow, maxRow, minCol, maxCol := len(grid), 0, len(grid[0]), 0

	for row, rowStr := range grid {
		for col, char := range rowStr {
			if char == '#' {
				if row < minRow {
					minRow = row
				}
				if row > maxRow {
					maxRow = row
				}
				if col < minCol {
					minCol = col
				}
				if col > maxCol {
					maxCol = col
				}
			}
		}
	}

	numRows := maxRow - minRow + 1
	numCols := maxCol - minCol + 1

	shiftedGrid := make([]string, numRows)
	for i := range shiftedGrid {
		shiftedGrid[i] = strings.Repeat(".", numCols)
	}

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if grid[row][col] == '#' {
				shiftedGrid[row-minRow] = shiftedGrid[row-minRow][:col-minCol] + "#" + shiftedGrid[row-minRow][col-minCol+1:]
			}
		}
	}

	return shiftedGrid
}

func printBoard(board *algorithm.Board) {
	for _, row := range board.Cells {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func Run(filename string) {
	content, _ := file.ReadFile(filename)

	re := regexp.MustCompile(`\n\n+`)
	splitStrs := re.Split(content, -1)
	n := len(splitStrs)

	for i := 0; i < n; i++ {
		splitStrs[i] = helper.CleanData(splitStrs[i])
	}
	tetrominoes := []*algorithm.Tetromino{}
	var char byte = 'A'
	for i := 0; i < n; i++ {
		if char > 90 {
			char = 97
		}
		if char > 122 {
			log.Println("You exceeded limit alphabet letters. Please, be content with all you have already")
			return
		}

		tetrominoes = append(tetrominoes, newTetromino(splitStrs[i], char))
		char++
	}
	smallestSquareFound := false

	for dim := int(math.Sqrt(float64(n * 4))); dim < 10; dim++ {
		fmt.Printf("Square Dimension: %d\n", dim)
		cells := generateBoard(dim)
		board := &algorithm.Board{
			Dim:   dim,
			Cells: cells,
		}
		if algorithm.Solve(board, tetrominoes) {
			smallestSquareFound = true
			fmt.Println("Smallest square found!")
			printBoard(board)
			return
		}
		fmt.Println("-------------------------------------------------")
	}

	if !smallestSquareFound {
		fmt.Println("No solution found.")
	}
}
