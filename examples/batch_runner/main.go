package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Batch Evaluation Demo")
	fmt.Println("========================================")

	const runs = 5

	var passed int

	start := time.Now()

	for i := 1; i <= runs; i++ {
		fmt.Printf("\nRun %d/%d\n", i, runs)
		fmt.Println("----------------------------------------")

		result, err := evaluator.Run()
		if err != nil {
			log.Fatalf("evaluation %d failed: %v", i, err)
		}

		fmt.Printf("Passed   : %v\n", result.Passed)
		fmt.Printf("Message  : %s\n", result.Message)
		fmt.Printf("Duration : %s\n", result.Duration)

		if result.Passed {
			passed++
		}
	}

	total := time.Since(start)

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("Batch Summary")
	fmt.Println("========================================")
	fmt.Printf("Runs           : %d\n", runs)
	fmt.Printf("Successful     : %d\n", passed)
	fmt.Printf("Failed         : %d\n", runs-passed)
	fmt.Printf("Success Rate   : %.2f%%\n", float64(passed)/float64(runs)*100)
	fmt.Printf("Total Time     : %s\n", total)
	fmt.Println()
	fmt.Println("Batch evaluation completed successfully.")
}
