package scenarios

import (
	"context"
	"errors"
	"testing"
)

func TestNewRegistry(t *testing.T) {
	registry := NewRegistry()

	if registry == nil {
		t.Fatal("expected registry")
	}

	if registry.Count() != 0 {
		t.Fatalf("expected empty registry, got %d", registry.Count())
	}
}

func TestRegisterScenario(t *testing.T) {
	registry := NewRegistry()

	scenario := &Scenario{
		ID:   "health-check",
		Name: "Health Check",
		Run: func(ctx context.Context) Result {
			return Result{
				Status: StatusPassed,
			}
		},
	}

	if err := registry.Register(scenario); err != nil {
		t.Fatalf("Register failed: %v", err)
	}

	if registry.Count() != 1 {
		t.Fatalf("expected 1 scenario, got %d", registry.Count())
	}
}

func TestRegisterDuplicateScenario(t *testing.T) {
	registry := NewRegistry()

	scenario := &Scenario{
		ID:   "duplicate",
		Name: "Duplicate",
		Run: func(ctx context.Context) Result {
			return Result{
				Status: StatusPassed,
			}
		},
	}

	if err := registry.Register(scenario); err != nil {
		t.Fatalf("Register failed: %v", err)
	}

	if err := registry.Register(scenario); !errors.Is(err, ErrScenarioAlreadyRegistered) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetScenario(t *testing.T) {
	registry := NewRegistry()

	scenario := &Scenario{
		ID:   "lookup",
		Name: "Lookup",
		Run: func(ctx context.Context) Result {
			return Result{
				Status: StatusPassed,
			}
		},
	}

	if err := registry.Register(scenario); err != nil {
		t.Fatalf("Register failed: %v", err)
	}

	got, err := registry.Get("lookup")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if got.ID != scenario.ID {
		t.Fatalf("unexpected scenario id: %s", got.ID)
	}
}

func TestGetMissingScenario(t *testing.T) {
	registry := NewRegistry()

	_, err := registry.Get("missing")

	if !errors.Is(err, ErrScenarioNotFound) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestListScenariosSorted(t *testing.T) {
	registry := NewRegistry()

	for _, id := range []string{"c", "a", "b"} {
		err := registry.Register(&Scenario{
			ID:   id,
			Name: id,
			Run: func(ctx context.Context) Result {
				return Result{
					Status: StatusPassed,
				}
			},
		})

		if err != nil {
			t.Fatalf("Register failed: %v", err)
		}
	}

	list := registry.List()

	if len(list) != 3 {
		t.Fatalf("expected 3 scenarios, got %d", len(list))
	}

	if list[0].ID != "a" || list[1].ID != "b" || list[2].ID != "c" {
		t.Fatal("registry returned scenarios in unexpected order")
	}
}