package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
)

func main() {
	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("===========================")
	fmt.Printf("Passed   : %v\n", result.Passed)
	fmt.Printf("Message  : %s\n", result.Message)
	fmt.Printf("Duration : %s\n", result.Duration)

	if result.Passed {
		fmt.Println()
		fmt.Println("VERDICT: ENGINEERING_EVALUATION_COMPLETED")
	} else {
		fmt.Println()
		fmt.Println("VERDICT: ENGINEERING_EVALUATION_FAILED")
	}
}