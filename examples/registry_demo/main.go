package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/scenarios"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Scenario Registry Demo")
	fmt.Println("========================================")

	registry := scenarios.NewRegistry()

	if err := scenarios.RegisterDefaultScenarios(registry); err != nil {
		log.Fatalf("failed to register default scenarios: %v", err)
	}

	list := registry.List()

	fmt.Printf("Registered Scenarios: %d\n\n", len(list))

	for i, scenario := range list {
		fmt.Printf("[%02d] %s\n", i+1, scenario.Name)
		fmt.Printf("Description : %s\n", scenario.Description)
		fmt.Printf("Category    : %s\n", scenario.Category)
		fmt.Printf("Enabled     : %v\n", scenario.Enabled)

		if scenario.Tags != nil && len(scenario.Tags) > 0 {
			fmt.Printf("Tags        : %v\n", scenario.Tags)
		}

		fmt.Println("----------------------------------------")
	}

	fmt.Println()
	fmt.Println("Scenario registry demonstration completed successfully.")
}