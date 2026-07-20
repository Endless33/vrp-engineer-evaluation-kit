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
	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("Manifest Example")
	fmt.Println("========================================")

	tempDir, err := os.MkdirTemp("", "vrp-manifest-example-*")
	if err != nil {
		log.Fatalf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	paths := []string{
		filepath.Join(tempDir, "vrp-report.md"),
		filepath.Join(tempDir, "evidence.json"),
		filepath.Join(tempDir, "evaluation.log"),
		filepath.Join(tempDir, "summary.json"),
	}

	for _, path := range paths {
		if err := os.WriteFile(path, []byte(filepath.Base(path)+"\n"), 0o644); err != nil {
			log.Fatalf("failed to create artifact: %v", err)
		}
	}

	manifest, err := output.NewManifest("evaluation-manifest")
	if err != nil {
		log.Fatalf("failed to create manifest: %v", err)
	}

	for _, path := range paths {
		if err := manifest.AddFile(path); err != nil {
			log.Fatalf("failed to add artifact: %v", err)
		}
	}

	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal manifest: %v", err)
	}

	fmt.Println()
	fmt.Println(string(data))
	fmt.Printf("\nArtifacts : %d\n", len(manifest.Artifacts))
}
