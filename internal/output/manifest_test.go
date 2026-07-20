package output

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewManifest(t *testing.T) {
	manifest, err := NewManifest("evaluation-run")
	if err != nil {
		t.Fatalf("NewManifest failed: %v", err)
	}

	if manifest.Version != ManifestVersion {
		t.Fatalf(
			"unexpected manifest version: got %q want %q",
			manifest.Version,
			ManifestVersion,
		)
	}

	if manifest.Name != "evaluation-run" {
		t.Fatalf(
			"unexpected manifest name: got %q",
			manifest.Name,
		)
	}

	if len(manifest.Artifacts) != 0 {
		t.Fatal("new manifest should not contain artifacts")
	}
}

func TestAddFile(t *testing.T) {
	manifest, err := NewManifest("evaluation-run")
	if err != nil {
		t.Fatalf("NewManifest failed: %v", err)
	}

	dir := t.TempDir()
	file := filepath.Join(dir, "artifact.txt")

	if err := os.WriteFile(file, []byte("public artifact"), 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	if err := manifest.AddFile(file); err != nil {
		t.Fatalf("AddFile failed: %v", err)
	}

	if len(manifest.Artifacts) != 1 {
		t.Fatalf("expected one artifact, got %d", len(manifest.Artifacts))
	}

	if manifest.Artifacts[0].SHA256 == "" {
		t.Fatal("artifact hash should not be empty")
	}
}

func TestWriteManifest(t *testing.T) {
	manifest, err := NewManifest("evaluation-run")
	if err != nil {
		t.Fatalf("NewManifest failed: %v", err)
	}

	output := filepath.Join(t.TempDir(), "manifest.json")

	if err := manifest.Write(output); err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	if _, err := os.Stat(output); err != nil {
		t.Fatalf("manifest file was not created: %v", err)
	}
}

func TestNewManifestRejectsEmptyName(t *testing.T) {
	if _, err := NewManifest(""); err != ErrManifestEmptyName {
		t.Fatalf(
			"unexpected error: got %v want %v",
			err,
			ErrManifestEmptyName,
		)
	}
}