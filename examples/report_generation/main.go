package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/report"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Report Generation Example")
	fmt.Println("========================================")

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	generator := report.New()

	const reportPath = "vrp-engineering-report.md"

	if err := generator.WriteMarkdown(reportPath, result); err != nil {
		log.Fatalf("failed to generate report: %v", err)
	}

	fmt.Println()
	fmt.Println("Evaluation Summary")
	fmt.Println("------------------")
	fmt.Printf("Passed   : %v\n", result.Passed)
	fmt.Printf("Message  : %s\n", result.Message)
	fmt.Printf("Duration : %s\n", result.Duration)
	fmt.Printf("Report   : %s\n", reportPath)
	fmt.Println()
	fmt.Println("Report generation completed successfully.")
}
