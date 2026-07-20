package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/version"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("Quick Start")
	fmt.Println("========================================")

	fmt.Printf("Version: %s\n\n", version.String())

	fmt.Println("Running engineering evaluation...")

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	fmt.Println()
	fmt.Println("Result")
	fmt.Println("----------------------------------------")
	fmt.Printf("Passed   : %v\n", result.Passed)
	fmt.Printf("Message  : %s\n", result.Message)
	fmt.Printf("Duration : %s\n", result.Duration)

	fmt.Println()

	if result.Passed {
		fmt.Println("SUCCESS")
		fmt.Println("The public engineering evaluation completed successfully.")
	} else {
		fmt.Println("FAILED")
		fmt.Println("The engineering evaluation reported a failure.")
	}

	fmt.Println()
	fmt.Println("Quick Start example completed.")
}