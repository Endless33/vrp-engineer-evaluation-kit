package main

import (
	"fmt"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/config"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Configuration Demo")
	fmt.Println("========================================")

	cfg := config.Default()

	fmt.Println()
	fmt.Println("Default Configuration")
	fmt.Println("----------------------------------------")
	fmt.Printf("Output Directory : %s\n", cfg.OutputDirectory)
	fmt.Printf("Report Name      : %s\n", cfg.ReportName)
	fmt.Printf("Verbose          : %v\n", cfg.Verbose)

	fmt.Println()
	fmt.Println("Applying custom configuration...")

	cfg.OutputDirectory = "./evaluation-results"
	cfg.ReportName = "engineering-evaluation.md"
	cfg.Verbose = true

	fmt.Println()
	fmt.Println("Updated Configuration")
	fmt.Println("----------------------------------------")
	fmt.Printf("Output Directory : %s\n", cfg.OutputDirectory)
	fmt.Printf("Report Name      : %s\n", cfg.ReportName)
	fmt.Printf("Verbose          : %v\n", cfg.Verbose)

	fmt.Println()
	fmt.Println("Configuration demonstration completed successfully.")
}