package output

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrOutputDirectoryEmpty = errors.New("output directory is empty")
	ErrOutputFileEmpty      = errors.New("output file is empty")
)

// EnsureDirectory creates the output directory if it does not already exist.
func EnsureDirectory(path string) error {
	if path == "" {
		return ErrOutputDirectoryEmpty
	}

	return os.MkdirAll(path, 0o755)
}

// WriteFile writes a public evaluation artifact to disk.
//
// This helper is intended only for public engineering artifacts such as
// reports, evidence, logs, and evaluation metadata.
//
// It must never be used to write protected runtime state, confidential
// implementation details, proprietary algorithms, private keys,
// authentication material, or internal VRP runtime data.
func WriteFile(directory, filename string, data []byte) (string, error) {
	if directory == "" {
		return "", ErrOutputDirectoryEmpty
	}

	if filename == "" {
		return "", ErrOutputFileEmpty
	}

	if err := EnsureDirectory(directory); err != nil {
		return "", fmt.Errorf("create output directory: %w", err)
	}

	fullPath := filepath.Join(directory, filename)

	if err := os.WriteFile(fullPath, data, 0o644); err != nil {
		return "", fmt.Errorf("write output file: %w", err)
	}

	return fullPath, nil
}

// Exists reports whether the specified output file exists.
func Exists(path string) bool {
	if path == "" {
		return false
	}

	_, err := os.Stat(path)
	return err == nil
}
