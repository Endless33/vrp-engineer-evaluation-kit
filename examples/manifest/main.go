package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/output"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("Manifest Example")
	fmt.Println("========================================")

	manifest := output.NewManifest()

	manifest.AddArtifact("vrp-report.md")
	manifest.AddArtifact("evidence.json")
	manifest.AddArtifact("evaluation.log")
	manifest.AddArtifact("summary.json")

	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal manifest: %v", err)
	}

	fmt.Println()
	fmt.Println("Generated Manifest")
	fmt.Println("----------------------------------------")
	fmt.Println(string(data))

	fmt.Println()
	fmt.Printf("Artifacts : %d\n", len(manifest.Artifacts))

	for i, artifact := range manifest.Artifacts {
		fmt.Printf("%2d. %s\n", i+1, artifact)
	}

	fmt.Println()
	fmt.Println("Manifest generated successfully.")
}