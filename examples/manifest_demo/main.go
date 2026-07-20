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
	fmt.Println("VRP Manifest Demo")
	fmt.Println("========================================")

	tempDir, err := os.MkdirTemp("", "vrp-manifest-demo-*")
	if err != nil {
		log.Fatalf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	artifactNames := []string{
		"engineering-report.md",
		"evaluation-results.json",
		"evidence.json",
		"runtime.log",
	}

	manifest, err := output.NewManifest("manifest-demo")
	if err != nil {
		log.Fatalf("failed to create manifest: %v", err)
	}

	for _, name := range artifactNames {
		path := filepath.Join(tempDir, name)

		if err := os.WriteFile(path, []byte("public artifact: "+name+"\n"), 0o644); err != nil {
			log.Fatalf("failed to create artifact: %v", err)
		}

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
	fmt.Printf("\nArtifacts: %d\n", len(manifest.Artifacts))
}
