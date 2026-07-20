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
		fmt.Printf("Category    : %s\n", scenario.Category)
		fmt.Printf("Description : %s\n", scenario.Description)
		fmt.Printf("Enabled     : %v\n", scenario.Enabled)

		if len(scenario.Tags) > 0 {
			fmt.Printf("Tags        : ")
			for i, tag := range scenario.Tags {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Print(tag)
			}
			fmt.Println()
		}

		fmt.Println("----------------------------------------")
	}

	fmt.Println()
	fmt.Println("Scenario listing completed successfully.")
}