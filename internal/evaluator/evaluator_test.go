package evaluator

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	result, err := Run()
	if err != nil {
		t.Fatalf("Run failed: %v", err)
	}

	if result == nil {
		t.Fatal("expected result")
	}

	if !result.Passed {
		t.Fatal("expected evaluation to pass")
	}

	if result.Message != "ENGINEERING_EVALUATION_COMPLETED" {
		t.Fatalf(
			"unexpected message: got %q",
			result.Message,
		)
	}

	if result.StartTime.IsZero() {
		t.Fatal("start time should not be zero")
	}

	if result.EndTime.IsZero() {
		t.Fatal("end time should not be zero")
	}

	if result.EndTime.Before(result.StartTime) {
		t.Fatal("end time must not be before start time")
	}

	if result.Duration <= 0 {
		t.Fatal("duration should be positive")
	}
}

func TestResultDurationConsistency(t *testing.T) {
	result := &Result{
		StartTime: time.Now().UTC(),
	}

	result.EndTime = result.StartTime.Add(250 * time.Millisecond)
	result.Duration = result.EndTime.Sub(result.StartTime)

	if result.Duration != 250*time.Millisecond {
		t.Fatalf(
			"unexpected duration: got %s",
			result.Duration,
		)
	}
}

func TestVersionConstant(t *testing.T) {
	if Version == "" {
		t.Fatal("version must not be empty")
	}
}