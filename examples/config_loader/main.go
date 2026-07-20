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
	fmt.Println("VRP Configuration Loader")
	fmt.Println("========================================")

	cfg := config.DefaultConfiguration()

	tempDir, err := os.MkdirTemp("", "vrp-config-loader-*")
	if err != nil {
		log.Fatalf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	configPath := filepath.Join(tempDir, "evaluation-config.json")

	cfg.OutputDirectory = filepath.Join(tempDir, "output")
	cfg.EvidenceDir = filepath.Join(cfg.OutputDirectory, "evidence")
	cfg.ReportsDir = filepath.Join(cfg.OutputDirectory, "reports")
	cfg.ManifestName = "evaluation-manifest.json"

	if err := config.Save(configPath, cfg); err != nil {
		log.Fatalf("failed to save configuration: %v", err)
	}

	loaded, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	fmt.Printf("Configuration File : %s\n", configPath)
	fmt.Printf("Output Directory   : %s\n", loaded.OutputDirectory)
	fmt.Printf("Evidence Directory : %s\n", loaded.EvidenceDir)
	fmt.Printf("Reports Directory  : %s\n", loaded.ReportsDir)
	fmt.Printf("Manifest Name      : %s\n", loaded.ManifestName)
}
