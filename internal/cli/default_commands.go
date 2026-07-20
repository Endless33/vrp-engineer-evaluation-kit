package cli

import (
	"fmt"
	"io"
)

// RegisterDefaultCommands registers the default public commands
// available in the VRP Engineer Evaluation Kit.
//
// These commands are intentionally limited to public engineering
// evaluation functionality. They must never expose protected
// runtime behavior, proprietary protocol logic, confidential
// implementation details, or internal VRP algorithms.
func RegisterDefaultCommands(c *CLI) error {
	commands := []Command{
		{
			Name:        "version",
			Description: "Display version information.",
			Run: func(w io.Writer, args []string) error {
				_, err := fmt.Fprintln(w, "VRP Engineer Evaluation Kit")
				return err
			},
		},
		{
			Name:        "evaluate",
			Description: "Run the public engineering evaluation.",
			Run: func(w io.Writer, args []string) error {
				_, err := fmt.Fprintln(w, "Engineering evaluation is available in this public toolkit.")
				return err
			},
		},
		{
			Name:        "report",
			Description: "Generate a public evaluation report.",
			Run: func(w io.Writer, args []string) error {
				_, err := fmt.Fprintln(w, "Public report generation completed.")
				return err
			},
		},
		{
			Name:        "help",
			Description: "Display available commands.",
			Run: func(w io.Writer, args []string) error {
				c.PrintHelp()
				return nil
			},
		},
	}

	for _, command := range commands {
		if err := c.Register(command); err != nil {
			return err
		}
	}

	return nil
}
