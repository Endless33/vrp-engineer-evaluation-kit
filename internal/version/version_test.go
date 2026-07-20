package version

import (
	"runtime"
	"strings"
	"testing"
)

func TestCurrent(t *testing.T) {
	info := Current()

	if info.Version != Version {
		t.Fatalf(
			"unexpected version: got %q want %q",
			info.Version,
			Version,
		)
	}

	if info.ReleaseName != ReleaseName {
		t.Fatalf(
			"unexpected release name: got %q want %q",
			info.ReleaseName,
			ReleaseName,
		)
	}

	if info.GoVersion == "" {
		t.Fatal("Go version should not be empty")
	}

	expectedPlatform := runtime.GOOS + "/" + runtime.GOARCH

	if info.Platform != expectedPlatform {
		t.Fatalf(
			"unexpected platform: got %q want %q",
			info.Platform,
			expectedPlatform,
		)
	}
}

func TestString(t *testing.T) {
	value := String()

	if value == "" {
		t.Fatal("version string should not be empty")
	}

	if !strings.Contains(value, Version) {
		t.Fatalf("version string does not contain version %q", Version)
	}

	if !strings.Contains(value, ReleaseName) {
		t.Fatalf("version string does not contain release name %q", ReleaseName)
	}

	if !strings.Contains(value, runtime.GOOS+"/"+runtime.GOARCH) {
		t.Fatal("version string does not contain platform")
	}
}