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

	builder := evidence.New()

	bundle, err := builder.Build(result)
	if err != nil {
		log.Fatalf("failed to build evidence: %v", err)
	}

	data, err := json.Marshal(bundle)
	if err != nil {
		log.Fatalf("failed to encode evidence: %v", err)
	}

	var decoded evidence.Bundle

	if err := json.Unmarshal(data, &decoded); err != nil {
		log.Fatalf("failed to decode evidence: %v", err)
	}

	if err := builder.Validate(&decoded); err != nil {
		log.Fatalf("evidence validation failed: %v", err)
	}

	fmt.Println()
	fmt.Println("Evidence Validation")
	fmt.Println("----------------------------------------")
	fmt.Printf("Version : %s\n", decoded.Version)
	fmt.Printf("Status  : VALID\n")
	fmt.Printf("Size    : %d bytes\n", len(data))
	fmt.Println()
	fmt.Println("Evidence successfully validated.")
}
