package main

import (
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

	const outDir = "manifest-output"

	if err := output.EnsureDirectory(outDir); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	artifactPaths := []string{
		filepath.Join(outDir, "evaluation-report.md"),
		filepath.Join(outDir, "evidence.json"),
		filepath.Join(outDir, "runtime.log"),
	}

	for _, path := range artifactPaths {
		if err := os.WriteFile(path, []byte("public evaluation artifact\n"), 0o644); err != nil {
			log.Fatalf("failed to create artifact %s: %v", path, err)
		}
	}

	manifest, err := output.NewManifest("vrp-evaluation-artifacts")
	if err != nil {
		log.Fatalf("failed to create manifest: %v", err)
	}

	for _, path := range artifactPaths {
		if err := manifest.AddFile(path); err != nil {
			log.Fatalf("failed to add artifact %s: %v", path, err)
		}
	}

	outFile := filepath.Join(outDir, "manifest.json")

	if err := manifest.Write(outFile); err != nil {
		log.Fatalf("failed to write manifest: %v", err)
	}

	fmt.Println()
	fmt.Printf("Artifacts : %d\n", len(manifest.Artifacts))
	fmt.Printf("Output    : %s\n", outFile)
}
