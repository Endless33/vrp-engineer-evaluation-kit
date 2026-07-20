package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/scenarios"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Scenario Listing Demo")
	fmt.Println("========================================")

	registry := scenarios.NewRegistry()

	if err := scenarios.RegisterDefaultScenarios(registry); err != nil {
		log.Fatalf("failed to register scenarios: %v", err)
	}

	list := registry.List()

	fmt.Printf("\nTotal Scenarios: %d\n\n", len(list))

	for index, scenario := range list {
		fmt.Printf("[%02d] %s\n", index+1, scenario.Name)
		fmt.Printf("ID          : %s\n", scenario.ID)
		fmt.Printf("Description : %s\n", scenario.Description)
		fmt.Printf("Timeout     : %s\n", scenario.Timeout)
		fmt.Println("----------------------------------------")
	}

	fmt.Println()
	fmt.Println("Scenario listing completed successfully.")
}
