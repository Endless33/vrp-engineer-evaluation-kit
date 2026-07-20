package main

import (
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
		fmt.Printf("Category    : %s\n", scenario.Category)
		fmt.Printf("Description : %s\n", scenario.Description)
		fmt.Printf("Enabled     : %v\n", scenario.Enabled)

		if len(scenario.Tags) > 0 {
			fmt.Printf("Tags        : %v\n", scenario.Tags)
		}

		fmt.Println("----------------------------------------")

		if !scenario.Enabled {
			continue
		}

		fmt.Printf("Running %q...\n", scenario.Name)

		if err := scenario.Run(); err != nil {
			fmt.Printf("Result: FAILED (%v)\n\n", err)
			continue
		}

		fmt.Println("Result: PASSED")
		fmt.Println()
	}

	fmt.Println("========================================")
	fmt.Println("Scenario execution completed.")
	fmt.Println("========================================")
}
