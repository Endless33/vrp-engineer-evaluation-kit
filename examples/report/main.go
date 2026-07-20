package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/report"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("Report Example")
	fmt.Println("========================================")

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	const reportFile = "vrp-report.md"

	generator := report.New()

	if err := generator.WriteMarkdown(reportFile, result); err != nil {
		log.Fatalf("failed to write report: %v", err)
	}

	fmt.Println()
	fmt.Println("Report Summary")
	fmt.Println("----------------------------------------")
	fmt.Printf("Passed      : %v\n", result.Passed)
	fmt.Printf("Message     : %s\n", result.Message)
	fmt.Printf("Duration    : %s\n", result.Duration)
	fmt.Printf("Report File : %s\n", reportFile)

	fmt.Println()
	fmt.Println("Report successfully generated.")
}