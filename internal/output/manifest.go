package output

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

const ManifestVersion = "vrp-output-manifest/v1"

var (
	ErrManifestNil       = errors.New("manifest is nil")
	ErrManifestEmptyName = errors.New("manifest name is empty")
)

// Manifest describes the public artifacts produced by the
// Engineer Evaluation Kit.
//
// It intentionally contains only public metadata and file hashes.
// It must never include protected runtime state, proprietary
// algorithms, private configuration, cryptographic secrets,
// or confidential implementation details.
type Manifest struct {
	Version     string      `json:"version"`
	Name        string      `json:"name"`
	GeneratedAt time.Time   `json:"generated_at"`
	Artifacts   []*Artifact `json:"artifacts"`
}

// Artifact describes one generated public artifact.
type Artifact struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	SHA256 string `json:"sha256"`
	Size   int64  `json:"size_bytes"`
}

// NewManifest creates an empty manifest.
func NewManifest(name string) (*Manifest, error) {
	if name == "" {
		return nil, ErrManifestEmptyName
	}

	return &Manifest{
		Version:     ManifestVersion,
		Name:        name,
		GeneratedAt: time.Now().UTC(),
		Artifacts:   []*Artifact{},
	}, nil
}

// AddFile adds a file to the manifest.
func (m *Manifest) AddFile(path string) error {
	if m == nil {
		return ErrManifestNil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read artifact: %w", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("stat artifact: %w", err)
	}

	sum := sha256.Sum256(data)

	m.Artifacts = append(m.Artifacts, &Artifact{
		Name:   filepath.Base(path),
		Path:   path,
		SHA256: hex.EncodeToString(sum[:]),
		Size:   info.Size(),
	})

	sort.Slice(m.Artifacts, func(i, j int) bool {
		return m.Artifacts[i].Path < m.Artifacts[j].Path
	})

	return nil
}

// Write saves the manifest as formatted JSON.
func (m *Manifest) Write(path string) error {
	if m == nil {
		return ErrManifestNil
	}

	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal manifest: %w", err)
	}

	if err := os.WriteFile(path, append(data, '\n'), 0o644); err != nil {
		return fmt.Errorf("write manifest: %w", err)
	}

	return nil
}
