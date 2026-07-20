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
	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("Evidence Example")
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

	data, err := json.MarshalIndent(bundle, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal evidence: %v", err)
	}

	fmt.Println()
	fmt.Println("Evidence Bundle")
	fmt.Println("----------------------------------------")
	fmt.Println(string(data))

	fmt.Println()
	fmt.Printf("Version : %s\n", bundle.Version)
	fmt.Printf("Status  : %v\n", result.Passed)
	fmt.Printf("Size    : %d bytes\n", len(data))

	fmt.Println()
	fmt.Println("Evidence generated successfully.")
}
