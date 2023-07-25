package main

import (
	"fmt"
	"log"
	"os"
	"tetris-optimizer/check"
	"tetris-optimizer/logdir"
	"tetris-optimizer/program"
	"time"
)

var e = "ERROR"

func main() {
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	defer logFile.Close()

	logdir.CustomLogger = log.New(logFile, "", log.LstdFlags)

	startTime := time.Now()
	logdir.LogFunc("-------------------------------------------------------")
	logdir.LogFunc("Program has started")
	params := os.Args[1:]

	if !check.Valid(params) {
		log.Println(e)
		return
	}

	program.Run(params[0])
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime.Minutes())
}
