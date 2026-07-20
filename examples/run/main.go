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
	fmt.Println("Run Example")
	fmt.Println("========================================")

	fmt.Printf("Version: %s\n", version.String())
	fmt.Println()

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	fmt.Println("Evaluation Result")
	fmt.Println("----------------------------------------")
	fmt.Printf("Passed      : %v\n", result.Passed)
	fmt.Printf("Message     : %s\n", result.Message)
	fmt.Printf("Started     : %s\n", result.StartTime.Format("2006-01-02 15:04:05 UTC"))
	fmt.Printf("Finished    : %s\n", result.EndTime.Format("2006-01-02 15:04:05 UTC"))
	fmt.Printf("Duration    : %s\n", result.Duration)

	fmt.Println()
	fmt.Println("========================================")
	if result.Passed {
		fmt.Println("FINAL VERDICT: PASS")
	} else {
		fmt.Println("FINAL VERDICT: FAIL")
	}
	fmt.Println("========================================")
}