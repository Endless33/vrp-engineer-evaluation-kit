package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/report"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Report Validation Demo")
	fmt.Println("========================================")

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	const reportPath = "validation-report.md"

	generator := report.New()

	if err := generator.WriteMarkdown(reportPath, result); err != nil {
		log.Fatalf("failed to generate report: %v", err)
	}

	info, err := os.Stat(reportPath)
	if err != nil {
		log.Fatalf("failed to verify report: %v", err)
	}

	if info.Size() == 0 {
		log.Fatal("generated report is empty")
	}

	data, err := os.ReadFile(reportPath)
	if err != nil {
		log.Fatalf("failed to read report: %v", err)
	}

	fmt.Println()
	fmt.Println("Validation Results")
	fmt.Println("----------------------------------------")
	fmt.Printf("Report File : %s\n", reportPath)
	fmt.Printf("Size        : %d bytes\n", info.Size())
	fmt.Printf("Readable    : %v\n", len(data) > 0)
	fmt.Printf("Passed      : %v\n", result.Passed)
	fmt.Println()
	fmt.Println("Report validation completed successfully.")
}