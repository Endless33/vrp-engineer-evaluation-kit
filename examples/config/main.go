package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/config"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("Configuration Example")
	fmt.Println("========================================")

	cfg := config.Default()

	fmt.Println()
	fmt.Println("Default Configuration")
	fmt.Println("----------------------------------------")
	fmt.Printf("Output Directory : %s\n", cfg.OutputDirectory)
	fmt.Printf("Report Name      : %s\n", cfg.ReportName)
	fmt.Printf("Verbose          : %v\n", cfg.Verbose)

	cfg.OutputDirectory = "./evaluation-output"
	cfg.ReportName = "evaluation-report.md"
	cfg.Verbose = true

	fmt.Println()
	fmt.Println("Modified Configuration")
	fmt.Println("----------------------------------------")
	fmt.Printf("Output Directory : %s\n", cfg.OutputDirectory)
	fmt.Printf("Report Name      : %s\n", cfg.ReportName)
	fmt.Printf("Verbose          : %v\n", cfg.Verbose)

	if err := cfg.Validate(); err != nil {
		log.Fatalf("configuration validation failed: %v", err)
	}

	fmt.Println()
	fmt.Println("Configuration successfully validated.")
}
