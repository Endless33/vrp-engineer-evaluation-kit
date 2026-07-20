package version

import (
	"fmt"
	"runtime"
)

const (
	Version     = "v0.1.0"
	ReleaseName = "Engineer Evaluation Kit"
)

// Info contains public version metadata only.
//
// This package intentionally exposes only build identification
// suitable for public engineering evaluation. It must never
// expose protected runtime versions, private component versions,
// internal commit mappings, confidential build metadata,
// signing information, or proprietary release details.
type Info struct {
	Version     string `json:"version"`
	ReleaseName string `json:"release_name"`
	GoVersion   string `json:"go_version"`
	Platform    string `json:"platform"`
}

// Current returns the current public version information.
func Current() Info {
	return Info{
		Version:     Version,
		ReleaseName: ReleaseName,
		GoVersion:   runtime.Version(),
		Platform:    runtime.GOOS + "/" + runtime.GOARCH,
	}
}

// String returns a human-readable version string.
func String() string {
	info := Current()

	return fmt.Sprintf(
		"%s %s (%s, %s)",
		info.ReleaseName,
		info.Version,
		info.GoVersion,
		info.Platform,
	)
}
