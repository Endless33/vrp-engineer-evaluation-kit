package scenarios

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrScenarioNil       = errors.New("scenario is nil")
	ErrScenarioIDEmpty   = errors.New("scenario id is empty")
	ErrScenarioNameEmpty = errors.New("scenario name is empty")
	ErrRunnerNil         = errors.New("scenario runner is nil")
)

// Status identifies the public outcome of a scenario execution.
type Status string

const (
	StatusPassed       Status = "PASSED"
	StatusFailed       Status = "FAILED"
	StatusSkipped      Status = "SKIPPED"
	StatusInconclusive Status = "INCONCLUSIVE"
)

// Scenario defines one public engineering evaluation scenario.
//
// A scenario in this repository may describe observable inputs,
// expected public outcomes, timing, and generated evidence.
//
// It must never contain protected VRP runtime implementation,
// proprietary decision logic, private authority rules, secret
// configuration, cryptographic key material, or internal algorithms.
type Scenario struct {
	ID          string
	Name        string
	Description string
	Timeout     time.Duration
	Run         Runner
}

// Runner executes a public evaluation scenario.
type Runner func(context.Context) Result

// Result contains only public scenario execution data.
type Result struct {
	ScenarioID string
	Status     Status
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   time.Duration
	Summary    string
	Details    map[string]string
	Err        error
}

// Validate checks whether a scenario is safe and executable.
func (s *Scenario) Validate() error {
	if s == nil {
		return ErrScenarioNil
	}

	if s.ID == "" {
		return ErrScenarioIDEmpty
	}

	if s.Name == "" {
		return ErrScenarioNameEmpty
	}

	if s.Run == nil {
		return ErrRunnerNil
	}

	if s.Timeout < 0 {
		return fmt.Errorf("scenario timeout must not be negative: %s", s.Timeout)
	}

	return nil
}

// Execute validates and runs the scenario.
//
// Timeout handling belongs to the evaluation harness only. It does not
// represent or expose protected runtime timeout behavior.
func (s *Scenario) Execute(parent context.Context) Result {
	if parent == nil {
		parent = context.Background()
	}

	startedAt := time.Now().UTC()

	if err := s.Validate(); err != nil {
		return Result{
			ScenarioID: scenarioID(s),
			Status:     StatusFailed,
			StartedAt:  startedAt,
			FinishedAt: time.Now().UTC(),
			Summary:    "SCENARIO_VALIDATION_FAILED",
			Err:        err,
		}
	}

	ctx := parent
	cancel := func() {}

	if s.Timeout > 0 {
		ctx, cancel = context.WithTimeout(parent, s.Timeout)
	}

	defer cancel()

	resultChannel := make(chan Result, 1)

	go func() {
		resultChannel <- s.Run(ctx)
	}()

	var result Result

	select {
	case result = <-resultChannel:
		if result.Status == "" {
			result.Status = StatusInconclusive
		}

	case <-ctx.Done():
		result = Result{
			Status:  StatusFailed,
			Summary: "SCENARIO_EXECUTION_TIMEOUT",
			Err:     ctx.Err(),
		}
	}

	finishedAt := time.Now().UTC()

	result.ScenarioID = s.ID
	result.StartedAt = startedAt
	result.FinishedAt = finishedAt
	result.Duration = finishedAt.Sub(startedAt)
	result.Details = cloneDetails(result.Details)

	return result
}

func scenarioID(s *Scenario) string {
	if s == nil {
		return ""
	}

	return s.ID
}

func cloneDetails(source map[string]string) map[string]string {
	if len(source) == 0 {
		return nil
	}

	cloned := make(map[string]string, len(source))

	for key, value := range source {
		cloned[key] = value
	}

	return cloned
}
