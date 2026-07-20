package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evidence"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/output"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/report"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Full Evaluation Pipeline")
	fmt.Println("========================================")

	const outDir = "pipeline-output"

	if err := output.EnsureDirectory(outDir); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	reportPath := filepath.Join(outDir, "evaluation-report.md")

	if err := report.New().WriteMarkdown(reportPath, result); err != nil {
		log.Fatalf("failed to write report: %v", err)
	}

	verdict := "FAILED"
	if result.Passed {
		verdict = "PASSED"
	}

	record, err := evidence.NewRecord(
		"full-pipeline-evidence",
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

	evidencePath := filepath.Join(outDir, "evidence.json")

	if err := evidence.WriteJSON(evidencePath, record); err != nil {
		log.Fatalf("failed to write evidence: %v", err)
	}

	manifest, err := output.NewManifest("full-evaluation-pipeline")
	if err != nil {
		log.Fatalf("failed to create manifest: %v", err)
	}

	for _, path := range []string{reportPath, evidencePath} {
		if err := manifest.AddFile(path); err != nil {
			log.Fatalf("failed to add artifact %s: %v", path, err)
		}
	}

	manifestPath := filepath.Join(outDir, "manifest.json")

	if err := manifest.Write(manifestPath); err != nil {
		log.Fatalf("failed to write manifest: %v", err)
	}

	fmt.Println()
	fmt.Println("Pipeline completed successfully.")
	fmt.Println("----------------------------------------")
	fmt.Printf("Evaluation : %v\n", result.Passed)
	fmt.Printf("Report     : %s\n", reportPath)
	fmt.Printf("Evidence   : %s\n", evidencePath)
	fmt.Printf("Manifest   : %s\n", manifestPath)
}
