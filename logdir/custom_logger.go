package logdir

import (
	"log"
	"time"
)

var CustomLogger *log.Logger

func LogFunctionTime(functionName string, start time.Time) {
	elapsed := time.Since(start)
	CustomLogger.Printf("Function [%s] took: %v minutes\n", functionName, elapsed.Minutes())
}

func LogFunc(s string) {
	CustomLogger.Printf("%s", s)
}
