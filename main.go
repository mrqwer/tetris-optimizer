package main

import (
	"log"
	"os"
	"tetris-optimizer/check"
	"tetris-optimizer/program"
)

var e = "ERROR"

func main() {
	params := os.Args[1:]

	if !check.Valid(params) {
		log.Println(e)
		return
	}

	program.Run(params[0])
}
