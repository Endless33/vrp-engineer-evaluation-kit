package scenarios

import (
	"errors"
	"sort"
	"sync"
)

var (
	ErrScenarioAlreadyRegistered = errors.New("scenario already registered")
	ErrScenarioNotFound          = errors.New("scenario not found")
)

// Registry stores public engineering evaluation scenarios.
//
// The registry contains only public evaluation definitions.
// It must never register protected runtime components,
// proprietary protocol logic, confidential algorithms,
// or internal VRP implementation details.
type Registry struct {
	mu        sync.RWMutex
	scenarios map[string]*Scenario
}

// NewRegistry creates an empty scenario registry.
func NewRegistry() *Registry {
	return &Registry{
		scenarios: make(map[string]*Scenario),
	}
}

// Register adds a scenario to the registry.
func (r *Registry) Register(s *Scenario) error {
	if err := s.Validate(); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.scenarios[s.ID]; exists {
		return ErrScenarioAlreadyRegistered
	}

	r.scenarios[s.ID] = s

	return nil
}

// Get returns a registered scenario.
func (r *Registry) Get(id string) (*Scenario, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	s, ok := r.scenarios[id]
	if !ok {
		return nil, ErrScenarioNotFound
	}

	return s, nil
}

// List returns all registered scenarios ordered by ID.
func (r *Registry) List() []*Scenario {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]*Scenario, 0, len(r.scenarios))

	for _, scenario := range r.scenarios {
		list = append(list, scenario)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})

	return list
}

// Count returns the number of registered scenarios.
func (r *Registry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.scenarios)
}