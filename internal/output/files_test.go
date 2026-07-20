package output

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEnsureDirectory(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "reports")

	if err := EnsureDirectory(dir); err != nil {
		t.Fatalf("EnsureDirectory failed: %v", err)
	}

	info, err := os.Stat(dir)
	if err != nil {
		t.Fatalf("directory does not exist: %v", err)
	}

	if !info.IsDir() {
		t.Fatal("expected a directory")
	}
}

func TestWriteFile(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "output")

	path, err := WriteFile(
		dir,
		"result.txt",
		[]byte("public engineering evaluation"),
	)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	if string(data) != "public engineering evaluation" {
		t.Fatalf("unexpected file contents: %q", string(data))
	}
}

func TestExists(t *testing.T) {
	dir := t.TempDir()

	path, err := WriteFile(
		dir,
		"artifact.txt",
		[]byte("artifact"),
	)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	if !Exists(path) {
		t.Fatal("expected file to exist")
	}

	if Exists(filepath.Join(dir, "missing.txt")) {
		t.Fatal("unexpected existing file")
	}
}

func TestWriteFileRejectsEmptyDirectory(t *testing.T) {
	if _, err := WriteFile(
		"",
		"file.txt",
		[]byte("data"),
	); err != ErrOutputDirectoryEmpty {
		t.Fatalf(
			"unexpected error: got %v want %v",
			err,
			ErrOutputDirectoryEmpty,
		)
	}
}

func TestWriteFileRejectsEmptyFilename(t *testing.T) {
	if _, err := WriteFile(
		t.TempDir(),
		"",
		[]byte("data"),
	); err != ErrOutputFileEmpty {
		t.Fatalf(
			"unexpected error: got %v want %v",
			err,
			ErrOutputFileEmpty,
		)
	}
}
