package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evidence"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/logging"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/output"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/report"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/scenarios"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/version"
)

func main() {
	logger := logging.New()

	logger.Section("VRP Engineer Evaluation Kit")
	logger.Info("Running complete demonstration.")

	fmt.Printf("Version: %s\n\n", version.String())

	registry := scenarios.NewRegistry()

	if err := scenarios.RegisterDefaultScenarios(registry); err != nil {
		log.Fatalf("failed to register scenarios: %v", err)
	}

	logger.Info("Registered scenarios: %d", len(registry.List()))

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	builder := evidence.New()

	bundle, err := builder.Build(result)
	if err != nil {
		log.Fatalf("failed to build evidence: %v", err)
	}

	reporter := report.New()

	const reportFile = "all-demo-report.md"

	if err := reporter.WriteMarkdown(reportFile, result); err != nil {
		log.Fatalf("failed to generate report: %v", err)
	}

	manifest := output.NewManifest()
	manifest.AddArtifact(reportFile)
	manifest.AddArtifact("evidence.json")

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("Complete Demonstration Summary")
	fmt.Println("========================================")
	fmt.Printf("Evaluation Passed : %v\n", result.Passed)
	fmt.Printf("Evidence Version  : %s\n", bundle.Version)
	fmt.Printf("Artifacts         : %d\n", len(manifest.Artifacts))
	fmt.Printf("Report            : %s\n", reportFile)
	fmt.Println()
	fmt.Println("Complete demonstration finished successfully.")
}