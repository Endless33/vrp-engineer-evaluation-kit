package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/output"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Manifest Generation Example")
	fmt.Println("========================================")

	manifest := output.NewManifest()

	manifest.AddArtifact("evaluation-report.md")
	manifest.AddArtifact("evidence.json")
	manifest.AddArtifact("runtime.log")

	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal manifest: %v", err)
	}

	outDir := "./manifest-output"

	if err := output.EnsureDirectory(outDir); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	outFile := filepath.Join(outDir, "manifest.json")

	if err := os.WriteFile(outFile, data, 0644); err != nil {
		log.Fatalf("failed to write manifest: %v", err)
	}

	fmt.Println()
	fmt.Printf("Artifacts : %d\n", len(manifest.Artifacts))
	fmt.Printf("Output    : %s\n", outFile)
	fmt.Println("Manifest generated successfully.")
}
