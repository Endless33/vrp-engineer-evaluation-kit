package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/output"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP JSON Export Demo")
	fmt.Println("========================================")

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal evaluation result: %v", err)
	}

	const outDir = "json-output"

	if err := output.EnsureDirectory(outDir); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	jsonPath, err := output.WriteFile(
		outDir,
		"evaluation-result.json",
		data,
	)
	if err != nil {
		log.Fatalf("failed to write JSON: %v", err)
	}

	fmt.Println()
	fmt.Println("Export Summary")
	fmt.Println("----------------------------------------")
	fmt.Printf("Passed      : %v\n", result.Passed)
	fmt.Printf("Message     : %s\n", result.Message)
	fmt.Printf("Output File : %s\n", filepath.Clean(jsonPath))
	fmt.Printf("JSON Size   : %d bytes\n", len(data))
	fmt.Println()
	fmt.Println("JSON export completed successfully.")
}