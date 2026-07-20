package evidence

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	FormatVersion = "vrp-evaluation-evidence/v1"
)

var (
	ErrEvidenceNil      = errors.New("evidence is nil")
	ErrOutputPathEmpty  = errors.New("evidence output path is empty")
	ErrEvidenceIDEmpty  = errors.New("evidence id is empty")
	ErrScenarioIDEmpty  = errors.New("scenario id is empty")
	ErrVerdictEmpty     = errors.New("evidence verdict is empty")
	ErrInvalidTimestamp = errors.New("evidence timestamp is invalid")
)

// Record contains public engineering-evaluation evidence only.
//
// This structure must never contain protected runtime state, private keys,
// proprietary algorithms, confidential configuration, or internal protocol
// decision data.
type Record struct {
	FormatVersion string            `json:"format_version"`
	EvidenceID    string            `json:"evidence_id"`
	ScenarioID    string            `json:"scenario_id"`
	GeneratedAt   time.Time         `json:"generated_at"`
	Verdict       string            `json:"verdict"`
	Summary       string            `json:"summary,omitempty"`
	Metadata      map[string]string `json:"metadata,omitempty"`
	DigestSHA256  string            `json:"digest_sha256"`
}

// NewRecord creates a public evidence record.
func NewRecord(
	evidenceID string,
	scenarioID string,
	verdict string,
	summary string,
	metadata map[string]string,
) (*Record, error) {
	record := &Record{
		FormatVersion: FormatVersion,
		EvidenceID:    evidenceID,
		ScenarioID:    scenarioID,
		GeneratedAt:   time.Now().UTC(),
		Verdict:       verdict,
		Summary:       summary,
		Metadata:      cloneMetadata(metadata),
	}

	if err := record.Validate(); err != nil {
		return nil, err
	}

	digest, err := record.computeDigest()
	if err != nil {
		return nil, fmt.Errorf("compute evidence digest: %w", err)
	}

	record.DigestSHA256 = digest

	return record, nil
}

// Validate checks whether the evidence record contains the required public
// evaluation fields.
func (r *Record) Validate() error {
	if r == nil {
		return ErrEvidenceNil
	}

	if r.FormatVersion == "" {
		r.FormatVersion = FormatVersion
	}

	if r.EvidenceID == "" {
		return ErrEvidenceIDEmpty
	}

	if r.ScenarioID == "" {
		return ErrScenarioIDEmpty
	}

	if r.Verdict == "" {
		return ErrVerdictEmpty
	}

	if r.GeneratedAt.IsZero() {
		return ErrInvalidTimestamp
	}

	return nil
}

// VerifyDigest recomputes and verifies the public evidence digest.
func (r *Record) VerifyDigest() error {
	if err := r.Validate(); err != nil {
		return err
	}

	if r.DigestSHA256 == "" {
		return errors.New("evidence digest is empty")
	}

	expected, err := r.computeDigest()
	if err != nil {
		return fmt.Errorf("compute expected evidence digest: %w", err)
	}

	if expected != r.DigestSHA256 {
		return fmt.Errorf(
			"evidence digest mismatch: expected %s, received %s",
			expected,
			r.DigestSHA256,
		)
	}

	return nil
}

// WriteJSON writes the evidence record as indented JSON.
func WriteJSON(path string, record *Record) error {
	if path == "" {
		return ErrOutputPathEmpty
	}

	if record == nil {
		return ErrEvidenceNil
	}

	if err := record.Validate(); err != nil {
		return err
	}

	if record.DigestSHA256 == "" {
		digest, err := record.computeDigest()
		if err != nil {
			return fmt.Errorf("compute evidence digest: %w", err)
		}

		record.DigestSHA256 = digest
	}

	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal evidence: %w", err)
	}

	parent := filepath.Dir(path)
	if parent != "." {
		if err := os.MkdirAll(parent, 0o755); err != nil {
			return fmt.Errorf("create evidence directory: %w", err)
		}
	}

	if err := os.WriteFile(path, append(data, '\n'), 0o644); err != nil {
		return fmt.Errorf("write evidence file: %w", err)
	}

	return nil
}

func (r *Record) computeDigest() (string, error) {
	type digestPayload struct {
		FormatVersion string            `json:"format_version"`
		EvidenceID    string            `json:"evidence_id"`
		ScenarioID    string            `json:"scenario_id"`
		GeneratedAt   time.Time         `json:"generated_at"`
		Verdict       string            `json:"verdict"`
		Summary       string            `json:"summary,omitempty"`
		Metadata      map[string]string `json:"metadata,omitempty"`
	}

	payload := digestPayload{
		FormatVersion: r.FormatVersion,
		EvidenceID:    r.EvidenceID,
		ScenarioID:    r.ScenarioID,
		GeneratedAt:   r.GeneratedAt.UTC(),
		Verdict:       r.Verdict,
		Summary:       r.Summary,
		Metadata:      cloneMetadata(r.Metadata),
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(data)

	return hex.EncodeToString(sum[:]), nil
}

func cloneMetadata(source map[string]string) map[string]string {
	if len(source) == 0 {
		return nil
	}

	cloned := make(map[string]string, len(source))

	for key, value := range source {
		cloned[key] = value
	}

	return cloned
}