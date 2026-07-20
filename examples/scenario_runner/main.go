package main

import (
	"context"
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

	for i, scenario := range list {
		fmt.Printf("%2d. %s\n", i+1, scenario.Name)
		fmt.Printf("    ID          : %s\n", scenario.ID)
		fmt.Printf("    Description : %s\n", scenario.Description)
		fmt.Printf("    Timeout     : %s\n", scenario.Timeout)

		result := scenario.Execute(context.Background())

		fmt.Printf("    Status      : %s\n", result.Status)
		fmt.Printf("    Summary     : %s\n", result.Summary)
		fmt.Printf("    Duration    : %s\n", result.Duration)

		if result.Err != nil {
			logger.Error("Scenario %q failed: %v", scenario.Name, result.Err)
		} else {
			logger.Info(
				"Scenario %q completed with status %s.",
				scenario.Name,
				result.Status,
			)
		}

		fmt.Println()
	}

	fmt.Println("========================================")
	fmt.Println("All registered scenarios processed.")
	fmt.Println("========================================")
}
