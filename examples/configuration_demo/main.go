package main

import (
	"fmt"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/config"
)

func printConfiguration(title string, cfg config.Config) {
	fmt.Println(title)
	fmt.Println("----------------------------------------")
	fmt.Printf("Output Directory : %s\n", cfg.OutputDirectory)
	fmt.Printf("Evidence Dir     : %s\n", cfg.EvidenceDir)
	fmt.Printf("Reports Dir      : %s\n", cfg.ReportsDir)
	fmt.Printf("Manifest Name    : %s\n", cfg.ManifestName)
	fmt.Printf("Is Default       : %v\n", cfg.IsDefault())
	fmt.Println()
}

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Configuration Demo")
	fmt.Println("========================================")
	fmt.Println()

	cfg := config.Default()
	printConfiguration("Default Configuration", cfg)

	cfg.OutputDirectory = "demo-output"
	cfg.EvidenceDir = "demo-output/evidence"
	cfg.ReportsDir = "demo-output/reports"
	cfg.ManifestName = "demo-manifest.json"

	printConfiguration("Modified Configuration", cfg)

	cfg.ResetToDefaults()
	printConfiguration("Reset Configuration", cfg)
}
