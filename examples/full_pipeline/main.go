package main

import (
	"encoding/json"
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

	outDir := "pipeline-output"

	if err := output.EnsureDirectory(outDir); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	reportGenerator := report.New()

	reportPath := filepath.Join(outDir, "evaluation-report.md")

	if err := reportGenerator.WriteMarkdown(reportPath, result); err != nil {
		log.Fatalf("failed to write report: %v", err)
	}

	evidenceBuilder := evidence.New()

	bundle, err := evidenceBuilder.Build(result)
	if err != nil {
		log.Fatalf("failed to build evidence: %v", err)
	}

	data, err := json.MarshalIndent(bundle, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal evidence: %v", err)
	}

	evidencePath, err := output.WriteFile(
		outDir,
		"evidence.json",
		data,
	)
	if err != nil {
		log.Fatalf("failed to write evidence: %v", err)
	}

	manifest := output.NewManifest()
	manifest.AddArtifact(reportPath)
	manifest.AddArtifact(evidencePath)

	manifestData, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal manifest: %v", err)
	}

	manifestPath, err := output.WriteFile(
		outDir,
		"manifest.json",
		manifestData,
	)
	if err != nil {
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
