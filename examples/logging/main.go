package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/evaluator"
	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/logging"
)

func main() {
	logger := logging.New()

	logger.Section("VRP Engineer Evaluation Kit")
	logger.Info("Logging example started.")

	result, err := evaluator.Run()
	if err != nil {
		logger.Error("Evaluation failed: %v", err)
		log.Fatal(err)
	}

	logger.Info("Evaluation completed successfully.")
	logger.Info("Passed: %v", result.Passed)
	logger.Info("Message: %s", result.Message)
	logger.Info("Duration: %s", result.Duration)

	fmt.Println()
	fmt.Println("Summary")
	fmt.Println("----------------------------------------")
	fmt.Printf("Passed   : %v\n", result.Passed)
	fmt.Printf("Message  : %s\n", result.Message)
	fmt.Printf("Duration : %s\n", result.Duration)

	logger.Section("Logging example finished")
}