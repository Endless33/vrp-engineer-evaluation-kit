package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	DefaultOutputDirectory = "output"
	DefaultEvidenceDir     = "output/evidence"
	DefaultReportsDir      = "output/reports"
	DefaultManifestName    = "manifest.json"
)

var (
	ErrConfigurationNil = errors.New("configuration is nil")
)

// Config contains only public configuration for the
// VRP Engineer Evaluation Kit.
//
// This configuration is intentionally limited to the public
// evaluation framework. It must never contain:
//
//   - protected runtime configuration;
//   - private protocol parameters;
//   - proprietary algorithms;
//   - cryptographic secrets;
//   - signing keys;
//   - authority configuration;
//   - confidential deployment information.
type Config struct {
	OutputDirectory string `json:"output_directory"`
	EvidenceDir     string `json:"evidence_directory"`
	ReportsDir      string `json:"reports_directory"`
	ManifestName    string `json:"manifest_name"`
}

// Default returns the default public configuration.
func Default() Config {
	return Config{
		OutputDirectory: DefaultOutputDirectory,
		EvidenceDir:     DefaultEvidenceDir,
		ReportsDir:      DefaultReportsDir,
		ManifestName:    DefaultManifestName,
	}
}

// Validate validates the configuration.
func (c *Config) Validate() error {
	if c == nil {
		return ErrConfigurationNil
	}

	if c.OutputDirectory == "" {
		c.OutputDirectory = DefaultOutputDirectory
	}

	if c.EvidenceDir == "" {
		c.EvidenceDir = DefaultEvidenceDir
	}

	if c.ReportsDir == "" {
		c.ReportsDir = DefaultReportsDir
	}

	if c.ManifestName == "" {
		c.ManifestName = DefaultManifestName
	}

	return nil
}

// Load loads a public configuration from disk.
func Load(path string) (Config, error) {
	cfg := Default()

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("read config: %w", err)
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("parse config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return cfg, err
	}

	return cfg, nil
}

// Save saves the configuration to disk.
func Save(path string, cfg Config) error {
	if err := cfg.Validate(); err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	data = append(data, '\n')

	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	return nil
}
