package config

// DefaultConfiguration returns the recommended public configuration
// used by the VRP Engineer Evaluation Kit.
//
// This configuration is intended solely for public engineering
// evaluation workflows. It must never include protected runtime
// configuration, confidential deployment information, proprietary
// protocol parameters, cryptographic material, authority settings,
// or any private implementation details.
func DefaultConfiguration() Config {
	return Config{
		OutputDirectory: DefaultOutputDirectory,
		EvidenceDir:     DefaultEvidenceDir,
		ReportsDir:      DefaultReportsDir,
		ManifestName:    DefaultManifestName,
	}
}

// ResetToDefaults replaces the current configuration with the
// recommended public defaults.
func (c *Config) ResetToDefaults() {
	if c == nil {
		return
	}

	*c = DefaultConfiguration()
}

// IsDefault reports whether the configuration matches the
// recommended public defaults.
func (c Config) IsDefault() bool {
	defaults := DefaultConfiguration()

	return c.OutputDirectory == defaults.OutputDirectory &&
		c.EvidenceDir == defaults.EvidenceDir &&
		c.ReportsDir == defaults.ReportsDir &&
		c.ManifestName == defaults.ManifestName
}