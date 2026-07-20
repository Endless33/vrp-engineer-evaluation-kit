package scenarios

import (
	"context"
	"time"
)

// RegisterDefaultScenarios registers the public engineering
// evaluation scenarios shipped with the evaluation kit.
//
// These scenarios demonstrate the evaluation framework only.
// They intentionally do not expose any protected VRP runtime,
// proprietary algorithms, confidential implementation,
// or internal protocol behavior.
func RegisterDefaultScenarios(registry *Registry) error {
	defaults := []*Scenario{
		{
			ID:          "framework-health-check",
			Name:        "Framework Health Check",
			Description: "Verifies that the evaluation framework is operational.",
			Timeout:     5 * time.Second,
			Run: func(ctx context.Context) Result {
				start := time.Now().UTC()

				select {
				case <-ctx.Done():
					end := time.Now().UTC()

					return Result{
						Status:     StatusFailed,
						StartedAt:  start,
						FinishedAt: end,
						Duration:   end.Sub(start),
						Summary:    "FRAMEWORK_HEALTH_CHECK_CANCELLED",
						Err:        ctx.Err(),
					}

				default:
				}

				end := time.Now().UTC()

				return Result{
					Status:     StatusPassed,
					StartedAt:  start,
					FinishedAt: end,
					Duration:   end.Sub(start),
					Summary:    "FRAMEWORK_HEALTH_CHECK_PASSED",
					Details: map[string]string{
						"framework": "ready",
						"profile":   "public-evaluation",
					},
				}
			},
		},
	}

	for _, scenario := range defaults {
		if err := registry.Register(scenario); err != nil {
			return err
		}
	}

	return nil
}