package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/config"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Configuration Example")
	fmt.Println("========================================")

	cfg := config.Default()

	fmt.Printf("Output Directory : %s\n", cfg.OutputDirectory)
	fmt.Printf("Evidence Dir     : %s\n", cfg.EvidenceDir)
	fmt.Printf("Reports Dir      : %s\n", cfg.ReportsDir)
	fmt.Printf("Manifest Name    : %s\n", cfg.ManifestName)
	fmt.Printf("Is Default       : %v\n", cfg.IsDefault())

	cfg.OutputDirectory = "custom-output"
	cfg.EvidenceDir = filepath.Join(cfg.OutputDirectory, "evidence")
	cfg.ReportsDir = filepath.Join(cfg.OutputDirectory, "reports")
	cfg.ManifestName = "custom-manifest.json"

	tempDir, err := os.MkdirTemp("", "vrp-config-example-*")
	if err != nil {
		log.Fatalf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	path := filepath.Join(tempDir, "config.json")

	if err := config.Save(path, cfg); err != nil {
		log.Fatalf("failed to save configuration: %v", err)
	}

	loaded, err := config.Load(path)
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	fmt.Println()
	fmt.Println("Loaded Configuration")
	fmt.Println("----------------------------------------")
	fmt.Printf("Output Directory : %s\n", loaded.OutputDirectory)
	fmt.Printf("Evidence Dir     : %s\n", loaded.EvidenceDir)
	fmt.Printf("Reports Dir      : %s\n", loaded.ReportsDir)
	fmt.Printf("Manifest Name    : %s\n", loaded.ManifestName)
}
