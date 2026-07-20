package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/config"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/report"
)

func main() {
	cfg := config.Default()

	cfg.OutputDirectory = "./custom-output"
	cfg.ReportName = "custom-evaluation-report.md"
	cfg.Verbose = true

	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation")
	fmt.Println("Custom Configuration Example")
	fmt.Println("========================================")

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	reportPath := cfg.OutputDirectory + "/" + cfg.ReportName

	generator := report.New()

	if err := generator.WriteMarkdown(reportPath, result); err != nil {
		log.Fatalf("failed to write report: %v", err)
	}

	fmt.Println()
	fmt.Printf("Evaluation Passed : %v\n", result.Passed)
	fmt.Printf("Result            : %s\n", result.Message)
	fmt.Printf("Duration          : %s\n", result.Duration)
	fmt.Printf("Report            : %s\n", reportPath)
}