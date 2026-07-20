package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/config"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Configuration Loader Demo")
	fmt.Println("========================================")

	cfg := config.Default()

	fmt.Println()
	fmt.Println("Default Configuration")
	fmt.Println("----------------------------------------")
	fmt.Printf("Output Directory : %s\n", cfg.OutputDirectory)
	fmt.Printf("Report Name      : %s\n", cfg.ReportName)
	fmt.Printf("Verbose          : %v\n", cfg.Verbose)

	fmt.Println()
	fmt.Println("Applying runtime overrides...")

	cfg.OutputDirectory = "./runtime-output"
	cfg.ReportName = "runtime-report.md"
	cfg.Verbose = true

	if err := cfg.Validate(); err != nil {
		log.Fatalf("configuration validation failed: %v", err)
	}

	fmt.Println()
	fmt.Println("Validated Configuration")
	fmt.Println("----------------------------------------")
	fmt.Printf("Output Directory : %s\n", cfg.OutputDirectory)
	fmt.Printf("Report Name      : %s\n", cfg.ReportName)
	fmt.Printf("Verbose          : %v\n", cfg.Verbose)

	fmt.Println()
	fmt.Println("Configuration loaded and validated successfully.")
}