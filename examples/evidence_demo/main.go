package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evidence"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Evidence Demo")
	fmt.Println("========================================")

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	verdict := "FAILED"
	if result.Passed {
		verdict = "PASSED"
	}

	record, err := evidence.NewRecord(
		"evidence-demo",
		"evaluation-run",
		verdict,
		result.Message,
		map[string]string{
			"duration": result.Duration.String(),
			"source":   "public-evaluation-kit",
		},
	)
	if err != nil {
		log.Fatalf("failed to create evidence: %v", err)
	}

	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal evidence: %v", err)
	}

	fmt.Println()
	fmt.Println("Generated Evidence")
	fmt.Println("----------------------------------------")
	fmt.Println(string(data))

	fmt.Println()
	fmt.Printf("Evaluation Passed : %v\n", result.Passed)
	fmt.Printf("Evidence Format   : %s\n", record.FormatVersion)
	fmt.Printf("Evidence Size     : %d bytes\n", len(data))
}
