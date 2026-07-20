package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evidence"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Evidence Export Example")
	fmt.Println("========================================")

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	ev := evidence.New()

	bundle, err := ev.Build(result)
	if err != nil {
		log.Fatalf("failed to build evidence: %v", err)
	}

	data, err := json.MarshalIndent(bundle, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal evidence: %v", err)
	}

	const output = "evidence.json"

	if err := os.WriteFile(output, data, 0644); err != nil {
		log.Fatalf("failed to write evidence: %v", err)
	}

	fmt.Println()
	fmt.Printf("Evaluation Passed : %v\n", result.Passed)
	fmt.Printf("Evidence File     : %s\n", output)
	fmt.Printf("Evidence Size     : %d bytes\n", len(data))
	fmt.Println("Evidence export completed successfully.")
}