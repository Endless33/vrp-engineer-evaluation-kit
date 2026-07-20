package cli

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestRegisterCommand(t *testing.T) {
	c := New()

	err := c.Register(Command{
		Name:        "test",
		Description: "test command",
		Run: func(w io.Writer, args []string) error {
			return nil
		},
	})
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}

	if len(c.commands) != 1 {
		t.Fatalf("expected one command, got %d", len(c.commands))
	}
}

func TestRegisterDuplicateCommand(t *testing.T) {
	c := New()

	cmd := Command{
		Name:        "duplicate",
		Description: "duplicate command",
		Run: func(w io.Writer, args []string) error {
			return nil
		},
	}

	if err := c.Register(cmd); err != nil {
		t.Fatalf("Register failed: %v", err)
	}

	if err := c.Register(cmd); !errors.Is(err, ErrCommandAlreadyRegistered) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestExecuteCommand(t *testing.T) {
	c := New()

	var executed bool

	err := c.Register(Command{
		Name:        "run",
		Description: "run command",
		Run: func(w io.Writer, args []string) error {
			executed = true
			return nil
		},
	})
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}

	if err := c.Execute("run", nil); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if !executed {
		t.Fatal("command was not executed")
	}
}

func TestExecuteUnknownCommand(t *testing.T) {
	c := New()

	if err := c.Execute("unknown", nil); !errors.Is(err, ErrCommandNotFound) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPrintHelp(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	c := &CLI{
		commands: make(map[string]Command),
		output:   buffer,
	}

	if err := c.Register(Command{
		Name:        "alpha",
		Description: "alpha command",
		Run: func(w io.Writer, args []string) error {
			return nil
		},
	}); err != nil {
		t.Fatalf("Register failed: %v", err)
	}

	c.PrintHelp()

	if buffer.Len() == 0 {
		t.Fatal("expected help output")
	}
}
