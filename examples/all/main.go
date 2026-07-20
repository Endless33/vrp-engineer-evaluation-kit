package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/cli"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/config"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evidence"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/output"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/report"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/scenarios"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/version"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("Complete Public API Example")
	fmt.Println("========================================")
	fmt.Printf("Version: %s\n", version.String())

	cfg := config.Default()

	tempDir, err := os.MkdirTemp("", "vrp-all-example-*")
	if err != nil {
		log.Fatalf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cfg.OutputDirectory = tempDir
	cfg.ReportsDir = filepath.Join(tempDir, "reports")
	cfg.EvidenceDir = filepath.Join(tempDir, "evidence")
	cfg.ManifestName = "manifest.json"

	if err := output.EnsureDirectory(cfg.ReportsDir); err != nil {
		log.Fatalf("failed to create reports directory: %v", err)
	}

	if err := output.EnsureDirectory(cfg.EvidenceDir); err != nil {
		log.Fatalf("failed to create evidence directory: %v", err)
	}

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	reportPath := filepath.Join(cfg.ReportsDir, "evaluation-report.md")

	if err := report.New().WriteMarkdown(reportPath, result); err != nil {
		log.Fatalf("failed to write report: %v", err)
	}

	verdict := "FAILED"
	if result.Passed {
		verdict = "PASSED"
	}

	record, err := evidence.NewRecord(
		"complete-example",
		"public-engineering-evaluation",
		verdict,
		result.Message,
		map[string]string{
			"duration": result.Duration.String(),
		},
	)
	if err != nil {
		log.Fatalf("failed to create evidence: %v", err)
	}

	evidencePath := filepath.Join(cfg.EvidenceDir, "evidence.json")

	if err := evidence.WriteJSON(evidencePath, record); err != nil {
		log.Fatalf("failed to write evidence: %v", err)
	}

	manifest, err := output.NewManifest("complete-public-api-example")
	if err != nil {
		log.Fatalf("failed to create manifest: %v", err)
	}

	for _, path := range []string{reportPath, evidencePath} {
		if err := manifest.AddFile(path); err != nil {
			log.Fatalf("failed to add artifact: %v", err)
		}
	}

	manifestPath := filepath.Join(cfg.OutputDirectory, cfg.ManifestName)

	if err := manifest.Write(manifestPath); err != nil {
		log.Fatalf("failed to write manifest: %v", err)
	}

	registry := scenarios.NewRegistry()

	if err := scenarios.RegisterDefaultScenarios(registry); err != nil {
		log.Fatalf("failed to register scenarios: %v", err)
	}

	passedScenarios := 0

	for _, scenario := range registry.List() {
		scenarioResult := scenario.Execute(context.Background())

		if scenarioResult.Status == scenarios.StatusPassed {
			passedScenarios++
		}
	}

	application := cli.New()

	if err := cli.RegisterDefaultCommands(application); err != nil {
		log.Fatalf("failed to register CLI commands: %v", err)
	}

	fmt.Println()
	fmt.Println("Summary")
	fmt.Println("----------------------------------------")
	fmt.Printf("Evaluation Passed : %v\n", result.Passed)
	fmt.Printf("Scenarios Passed  : %d/%d\n", passedScenarios, registry.Count())
	fmt.Printf("Report            : %s\n", reportPath)
	fmt.Printf("Evidence          : %s\n", evidencePath)
	fmt.Printf("Manifest          : %s\n", manifestPath)

	fmt.Println()
	application.PrintHelp()
}
