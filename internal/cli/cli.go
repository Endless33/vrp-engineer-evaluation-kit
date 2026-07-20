package cli

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

var (
	ErrCommandAlreadyRegistered = errors.New("command already registered")
	ErrCommandNotFound          = errors.New("command not found")
)

// Command represents one public CLI command.
//
// Commands in this repository are limited to public engineering
// evaluation functionality only. They must never expose protected
// runtime operations, proprietary protocol logic, confidential
// configuration, or internal VRP implementation details.
type Command struct {
	Name        string
	Description string
	Run         func(io.Writer, []string) error
}

// CLI represents the public command-line interface.
type CLI struct {
	commands map[string]Command
	output   io.Writer
}

// New creates a new CLI.
func New() *CLI {
	return &CLI{
		commands: make(map[string]Command),
		output:   os.Stdout,
	}
}

// Register registers a command.
func (c *CLI) Register(cmd Command) error {
	if _, exists := c.commands[cmd.Name]; exists {
		return ErrCommandAlreadyRegistered
	}

	c.commands[cmd.Name] = cmd
	return nil
}

// Execute executes a registered command.
func (c *CLI) Execute(name string, args []string) error {
	cmd, ok := c.commands[name]
	if !ok {
		return ErrCommandNotFound
	}

	return cmd.Run(c.output, args)
}

// PrintHelp prints the available commands.
func (c *CLI) PrintHelp() {
	names := make([]string, 0, len(c.commands))

	for name := range c.commands {
		names = append(names, name)
	}

	sort.Strings(names)

	fmt.Fprintln(c.output, "Available commands:")
	fmt.Fprintln(c.output)

	for _, name := range names {
		cmd := c.commands[name]
		fmt.Fprintf(
			c.output,
			"  %-20s %s\n",
			cmd.Name,
			cmd.Description,
		)
	}
}
