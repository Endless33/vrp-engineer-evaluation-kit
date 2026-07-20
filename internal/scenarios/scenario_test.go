package scenarios

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestScenarioValidate(t *testing.T) {
	s := &Scenario{
		ID:      "health-check",
		Name:    "Health Check",
		Timeout: time.Second,
		Run: func(ctx context.Context) Result {
			return Result{
				Status:  StatusPassed,
				Summary: "OK",
			}
		},
	}

	if err := s.Validate(); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}
}

func TestScenarioValidateNil(t *testing.T) {
	var s *Scenario

	if err := s.Validate(); !errors.Is(err, ErrScenarioNil) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestScenarioExecuteSuccess(t *testing.T) {
	s := &Scenario{
		ID:      "success",
		Name:    "Success Scenario",
		Timeout: time.Second,
		Run: func(ctx context.Context) Result {
			return Result{
				Status:  StatusPassed,
				Summary: "SUCCESS",
			}
		},
	}

	result := s.Execute(context.Background())

	if result.Status != StatusPassed {
		t.Fatalf("unexpected status: %s", result.Status)
	}

	if result.ScenarioID != "success" {
		t.Fatalf("unexpected scenario id: %s", result.ScenarioID)
	}
}

func TestScenarioExecuteValidationFailure(t *testing.T) {
	s := &Scenario{}

	result := s.Execute(context.Background())

	if result.Status != StatusFailed {
		t.Fatalf("unexpected status: %s", result.Status)
	}

	if result.Err == nil {
		t.Fatal("expected validation error")
	}
}

func TestScenarioExecuteTimeout(t *testing.T) {
	s := &Scenario{
		ID:      "timeout",
		Name:    "Timeout Scenario",
		Timeout: 20 * time.Millisecond,
		Run: func(ctx context.Context) Result {
			<-ctx.Done()

			return Result{
				Status: StatusFailed,
				Err:    ctx.Err(),
			}
		},
	}

	result := s.Execute(context.Background())

	if result.Err == nil {
		t.Fatal("expected timeout error")
	}
}

func TestScenarioExecuteWithoutTimeout(t *testing.T) {
	s := &Scenario{
		ID:   "no-timeout",
		Name: "No Timeout",
		Run: func(ctx context.Context) Result {
			return Result{
				Status: StatusPassed,
			}
		},
	}

	result := s.Execute(context.Background())

	if result.Status != StatusPassed {
		t.Fatalf("unexpected status: %s", result.Status)
	}
}
