package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/config"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/output"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/report"
)

func main() {
	cfg := config.Default()

	cfg.OutputDirectory = "./custom-output"
	cfg.EvidenceDir = filepath.Join(cfg.OutputDirectory, "evidence")
	cfg.ReportsDir = filepath.Join(cfg.OutputDirectory, "reports")
	cfg.ManifestName = "custom-manifest.json"

	if err := cfg.Validate(); err != nil {
		log.Fatalf("invalid configuration: %v", err)
	}

	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation")
	fmt.Println("Custom Configuration Example")
	fmt.Println("========================================")

	if err := output.EnsureDirectory(cfg.ReportsDir); err != nil {
		log.Fatalf("failed to create reports directory: %v", err)
	}

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	reportPath := filepath.Join(
		cfg.ReportsDir,
		"custom-evaluation-report.md",
	)

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
