package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/output"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Manifest Demo")
	fmt.Println("========================================")

	manifest := output.NewManifest()

	manifest.AddArtifact("engineering-report.md")
	manifest.AddArtifact("evaluation-results.json")
	manifest.AddArtifact("evidence.json")
	manifest.AddArtifact("runtime.log")
	manifest.AddArtifact("manifest.json")

	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal manifest: %v", err)
	}

	fmt.Println()
	fmt.Println("Generated Manifest")
	fmt.Println("----------------------------------------")
	fmt.Println(string(data))

	fmt.Println()
	fmt.Printf("Artifacts: %d\n", len(manifest.Artifacts))
	fmt.Println("Manifest demonstration completed successfully.")
}