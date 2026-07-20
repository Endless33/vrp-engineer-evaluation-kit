package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evidence"
)

func verdict(passed bool) string {
	if passed {
		return "PASSED"
	}

	return "FAILED"
}

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("Evidence Example")
	fmt.Println("========================================")

	result, err := evaluator.Run()
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}

	record, err := evidence.NewRecord(
		"evaluation-example",
		"public-engineering-evaluation",
		verdict(result.Passed),
		result.Message,
		map[string]string{
			"duration": result.Duration.String(),
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
	fmt.Println("Evidence Record")
	fmt.Println("----------------------------------------")
	fmt.Println(string(data))

	fmt.Println()
	fmt.Printf("Format  : %s\n", record.FormatVersion)
	fmt.Printf("Verdict : %s\n", record.Verdict)
	fmt.Printf("Size    : %d bytes\n", len(data))
}
