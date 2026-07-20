package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/output"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/report"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Markdown Export Demo")
	fmt.Println("========================================")

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	const outDir = "markdown-output"

	if err := output.EnsureDirectory(outDir); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	reportGenerator := report.New()

	reportPath := filepath.Join(
		outDir,
		"engineering-evaluation-report.md",
	)

	if err := reportGenerator.WriteMarkdown(reportPath, result); err != nil {
		log.Fatalf("failed to generate markdown report: %v", err)
	}

	fmt.Println()
	fmt.Println("Markdown Export Summary")
	fmt.Println("----------------------------------------")
	fmt.Printf("Passed      : %v\n", result.Passed)
	fmt.Printf("Message     : %s\n", result.Message)
	fmt.Printf("Report File : %s\n", reportPath)
	fmt.Println()
	fmt.Println("Markdown report exported successfully.")
}
