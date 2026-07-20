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
	fmt.Println("VRP Multi Report Demo")
	fmt.Println("========================================")

	const outputDir = "multi-report-output"

	if err := output.EnsureDirectory(outputDir); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	reportGenerator := report.New()

	for i := 1; i <= 3; i++ {
		result, err := evaluator.Run()
		if err != nil {
			log.Fatalf("evaluation %d failed: %v", i, err)
		}

		reportPath := filepath.Join(
			outputDir,
			fmt.Sprintf("evaluation-report-%d.md", i),
		)

		if err := reportGenerator.WriteMarkdown(reportPath, result); err != nil {
			log.Fatalf("failed to write report %d: %v", i, err)
		}

		fmt.Printf(
			"Generated report %d -> %s\n",
			i,
			reportPath,
		)
	}

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("Summary")
	fmt.Println("========================================")
	fmt.Println("Generated reports: 3")
	fmt.Printf("Output directory : %s\n", outputDir)
	fmt.Println("Multi-report demonstration completed successfully.")
}