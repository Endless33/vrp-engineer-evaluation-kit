package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/logging"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/scenarios"
)

func main() {
	logger := logging.New()

	logger.Section("VRP Scenario Runner")

	registry := scenarios.NewRegistry()

	if err := scenarios.RegisterDefaultScenarios(registry); err != nil {
		log.Fatalf("failed to register scenarios: %v", err)
	}

	list := registry.List()

	if len(list) == 0 {
		logger.Warning("No scenarios registered.")
		return
	}

	logger.Info("Registered scenarios: %d", len(list))

	fmt.Println()
	fmt.Println("Available Scenarios")
	fmt.Println("===================")

	for i, scenario := range list {
		fmt.Printf("%2d. %s\n", i+1, scenario.Name)
		fmt.Printf("    Description : %s\n", scenario.Description)
		fmt.Printf("    Category    : %s\n", scenario.Category)
		fmt.Printf("    Enabled     : %v\n", scenario.Enabled)
		fmt.Println()

		if !scenario.Enabled {
			logger.Warning("Scenario %q is disabled.", scenario.Name)
			continue
		}

		logger.Info("Executing %q...", scenario.Name)

		if err := scenario.Run(); err != nil {
			logger.Error("Scenario %q failed: %v", scenario.Name, err)
			continue
		}

		logger.Info("Scenario %q completed successfully.", scenario.Name)
	}

	fmt.Println("========================================")
	fmt.Println("All runnable scenarios processed.")
	fmt.Println("========================================")
}