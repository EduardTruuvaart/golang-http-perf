package main

import (
	"fmt"
	"time"
)

func main() {
	//lambda.Start(runLambda)
	runLambda()
}

func runLambda() {
	start := time.Now()

	fmt.Println("Test log")

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
