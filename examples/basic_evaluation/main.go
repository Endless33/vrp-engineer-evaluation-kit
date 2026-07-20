package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/report"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation Example")
	fmt.Println("========================================")

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	fmt.Printf("Status   : %v\n", result.Passed)
	fmt.Printf("Message  : %s\n", result.Message)
	fmt.Printf("Duration : %s\n", result.Duration)

	generator := report.New()

	if err := generator.WriteMarkdown("evaluation-report.md", result); err != nil {
		log.Fatalf("unable to write report: %v", err)
	}

	fmt.Println()
	fmt.Println("Evaluation completed successfully.")
	fmt.Println("Report written to evaluation-report.md")
}