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
	fmt.Println("VRP Evidence Validation Demo")
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
		"evidence-validation-demo",
		"public-engineering-evaluation",
		verdict,
		result.Message,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to create evidence: %v", err)
	}

	data, err := json.Marshal(record)
	if err != nil {
		log.Fatalf("failed to encode evidence: %v", err)
	}

	var decoded evidence.Record

	if err := json.Unmarshal(data, &decoded); err != nil {
		log.Fatalf("failed to decode evidence: %v", err)
	}

	if err := decoded.Validate(); err != nil {
		log.Fatalf("evidence field validation failed: %v", err)
	}

	if err := decoded.VerifyDigest(); err != nil {
		log.Fatalf("evidence digest validation failed: %v", err)
	}

	fmt.Println()
	fmt.Println("Evidence Validation")
	fmt.Println("----------------------------------------")
	fmt.Printf("Format : %s\n", decoded.FormatVersion)
	fmt.Printf("Status : VALID\n")
	fmt.Printf("Size   : %d bytes\n", len(data))
}
