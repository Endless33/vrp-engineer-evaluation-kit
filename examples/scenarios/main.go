package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/scenarios"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("Scenarios Example")
	fmt.Println("========================================")

	registry := scenarios.NewRegistry()

	if err := scenarios.RegisterDefaultScenarios(registry); err != nil {
		log.Fatalf("failed to register scenarios: %v", err)
	}

	list := registry.List()

	fmt.Printf("Registered Scenarios: %d\n\n", len(list))

	for i, scenario := range list {
		fmt.Printf("[%02d] %s\n", i+1, scenario.Name)
		fmt.Printf("ID          : %s\n", scenario.ID)
		fmt.Printf("Description : %s\n", scenario.Description)
		fmt.Printf("Timeout     : %s\n", scenario.Timeout)

		result := scenario.Execute(context.Background())

		fmt.Printf("Status      : %s\n", result.Status)
		fmt.Printf("Summary     : %s\n", result.Summary)
		fmt.Printf("Duration    : %s\n", result.Duration)

		if result.Err != nil {
			fmt.Printf("Error       : %v\n", result.Err)
		}

		fmt.Println("----------------------------------------")
	}

	fmt.Println("Scenario execution completed.")
}
