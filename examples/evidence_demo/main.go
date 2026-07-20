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

	builder := evidence.New()

	ev, err := builder.Build(result)
	if err != nil {
		log.Fatalf("failed to build evidence: %v", err)
	}

	data, err := json.MarshalIndent(ev, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal evidence: %v", err)
	}

	fmt.Println()
	fmt.Println("Generated Evidence")
	fmt.Println("----------------------------------------")
	fmt.Println(string(data))

	fmt.Println()
	fmt.Printf("Evaluation Passed : %v\n", result.Passed)
	fmt.Printf("Evidence Version  : %s\n", ev.Version)
	fmt.Printf("Evidence Size     : %d bytes\n", len(data))
	fmt.Println("Evidence demonstration completed successfully.")
}
