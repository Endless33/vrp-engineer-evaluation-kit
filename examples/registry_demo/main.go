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
		fmt.Printf("ID          : %s\n", scenario.ID)
		fmt.Printf("Description : %s\n", scenario.Description)
		fmt.Printf("Timeout     : %s\n", scenario.Timeout)
		fmt.Println("----------------------------------------")
	}
}
