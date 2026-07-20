package evidence

import (
	"path/filepath"
	"testing"
)

func TestNewRecord(t *testing.T) {
	record, err := NewRecord(
		"evidence-001",
		"framework-health-check",
		"PASSED",
		"Public engineering evaluation completed.",
		map[string]string{
			"profile": "public",
		},
	)
	if err != nil {
		t.Fatalf("NewRecord failed: %v", err)
	}

	if record.FormatVersion != FormatVersion {
		t.Fatalf(
			"unexpected format version: got %q want %q",
			record.FormatVersion,
			FormatVersion,
		)
	}

	if record.DigestSHA256 == "" {
		t.Fatal("expected digest to be generated")
	}
}

func TestVerifyDigest(t *testing.T) {
	record, err := NewRecord(
		"evidence-002",
		"framework-health-check",
		"PASSED",
		"",
		nil,
	)
	if err != nil {
		t.Fatalf("NewRecord failed: %v", err)
	}

	if err := record.VerifyDigest(); err != nil {
		t.Fatalf("VerifyDigest failed: %v", err)
	}
}

func TestVerifyDigestDetectsModification(t *testing.T) {
	record, err := NewRecord(
		"evidence-003",
		"framework-health-check",
		"PASSED",
		"",
		nil,
	)
	if err != nil {
		t.Fatalf("NewRecord failed: %v", err)
	}

	record.Verdict = "FAILED"

	if err := record.VerifyDigest(); err == nil {
		t.Fatal("expected digest verification to fail")
	}
}

func TestWriteJSON(t *testing.T) {
	record, err := NewRecord(
		"evidence-004",
		"framework-health-check",
		"PASSED",
		"",
		nil,
	)
	if err != nil {
		t.Fatalf("NewRecord failed: %v", err)
	}

	path := filepath.Join(t.TempDir(), "evidence.json")

	if err := WriteJSON(path, record); err != nil {
		t.Fatalf("WriteJSON failed: %v", err)
	}
}

func TestValidateRejectsMissingFields(t *testing.T) {
	record := &Record{}

	if err := record.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}
