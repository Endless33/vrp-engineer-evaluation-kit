package config

import "testing"

func TestDefaultConfiguration(t *testing.T) {
	cfg := DefaultConfiguration()

	if cfg.OutputDirectory != DefaultOutputDirectory {
		t.Fatalf(
			"unexpected output directory: got %q want %q",
			cfg.OutputDirectory,
			DefaultOutputDirectory,
		)
	}

	if cfg.EvidenceDir != DefaultEvidenceDir {
		t.Fatalf(
			"unexpected evidence directory: got %q want %q",
			cfg.EvidenceDir,
			DefaultEvidenceDir,
		)
	}

	if cfg.ReportsDir != DefaultReportsDir {
		t.Fatalf(
			"unexpected reports directory: got %q want %q",
			cfg.ReportsDir,
			DefaultReportsDir,
		)
	}

	if cfg.ManifestName != DefaultManifestName {
		t.Fatalf(
			"unexpected manifest name: got %q want %q",
			cfg.ManifestName,
			DefaultManifestName,
		)
	}
}

func TestResetToDefaults(t *testing.T) {
	cfg := Config{
		OutputDirectory: "tmp",
		EvidenceDir:     "tmp/evidence",
		ReportsDir:      "tmp/reports",
		ManifestName:    "custom.json",
	}

	cfg.ResetToDefaults()

	if !cfg.IsDefault() {
		t.Fatal("configuration should match defaults")
	}
}

func TestValidatePopulatesDefaults(t *testing.T) {
	cfg := Config{}

	if err := cfg.Validate(); err != nil {
		t.Fatalf("validate failed: %v", err)
	}

	if !cfg.IsDefault() {
		t.Fatal("configuration should contain default values after validation")
	}
}

func TestValidateNilConfiguration(t *testing.T) {
	var cfg *Config

	if err := cfg.Validate(); err != ErrConfigurationNil {
		t.Fatalf(
			"unexpected error: got %v want %v",
			err,
			ErrConfigurationNil,
		)
	}
}

func TestIsDefault(t *testing.T) {
	cfg := DefaultConfiguration()

	if !cfg.IsDefault() {
		t.Fatal("expected configuration to be default")
	}

	cfg.ManifestName = "another.json"

	if cfg.IsDefault() {
		t.Fatal("configuration should no longer be default")
	}
}
