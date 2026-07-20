package main

import (
	"fmt"
	"log"

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

	verdict := "FAILED"
	if result.Passed {
		verdict = "PASSED"
	}

	record, err := evidence.NewRecord(
		"saved-evidence",
		"public-engineering-evaluation",
		verdict,
		result.Message,
		map[string]string{
			"duration": result.Duration.String(),
		},
	)
	if err != nil {
		log.Fatalf("failed to create evidence: %v", err)
	}

	const outputPath = "evidence.json"

	if err := evidence.WriteJSON(outputPath, record); err != nil {
		log.Fatalf("failed to write evidence: %v", err)
	}

	fmt.Println()
	fmt.Printf("Evaluation Passed : %v\n", result.Passed)
	fmt.Printf("Evidence File     : %s\n", outputPath)
	fmt.Printf("Evidence Digest   : %s\n", record.DigestSHA256)
}
