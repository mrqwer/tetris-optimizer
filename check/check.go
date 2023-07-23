package check

import (
	"path/filepath"
	"strings"
	"tetris-optimizer/file"
)

func Valid(params []string) bool {
	if len(params) != 1 {
		return false
	}

	if !checkFileExt(params[0]) {
		return false
	}
	if !file.FileExists(params[0]) {
		return false
	}

	content, err := file.ReadFile(params[0])
	if err != nil {
		return false
	}

	if !checkFileData(content) {
		return false
	}

	blocks := strings.Split(content, "\n\n")

	for _, block := range blocks {
		if !validateTetromino(block) {
			return false
		}
	}

	return true
}

func checkFileExt(s string) bool {
	ext := filepath.Ext(s)
	if ext != ".txt" {
		return false
	}
	return true
}

func checkFileData(content string) bool {
	blocks := strings.Split(content, "\n\n")
	for _, block := range blocks {
		lines := strings.Split(strings.Trim(block, "\n\t "), "\n")
		numRows := len(lines)
		numCols := len(lines[0])
		if numRows != 4 || numCols != 4 || strings.Count(block, "#") != 4 || strings.Count(block, ".") != 12 {
			return false
		}
	}
	return true
}

func validateTetromino(input string) bool {
	lines := strings.Split(strings.Trim(input, "\n\t "), "\n")
	numRows := len(lines)
	numCols := len(lines[0])

	for i := 0; i < numRows; i++ {
		lines[i] = "." + lines[i] + "."
	}
	lines = append([]string{strings.Repeat(".", numCols+2)}, lines...)
	lines = append(lines, strings.Repeat(".", numCols+2))

	adjacentBlocks := 0
	for i := 1; i <= numRows; i++ {
		for j := 1; j <= numCols; j++ {
			if lines[i][j] == '#' {
				if lines[i-1][j] == '#' { // above
					adjacentBlocks++
				}
				if lines[i+1][j] == '#' { // below
					adjacentBlocks++
				}
				if lines[i][j-1] == '#' { // left
					adjacentBlocks++
				}
				if lines[i][j+1] == '#' { // right
					adjacentBlocks++
				}
			}
		}
	}

	return adjacentBlocks == 6 || adjacentBlocks == 8
}
