package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestVersionConstant(t *testing.T) {
	if Version == "" {
		t.Fatal("version must not be empty")
	}
}

func TestVersionCommand(t *testing.T) {
	cmd := exec.Command(os.Args[0], "-test.run=TestHelperProcessVersion")

	cmd.Env = append(
		os.Environ(),
		"GO_WANT_HELPER_PROCESS=1",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("helper process failed: %v", err)
	}

	if !strings.Contains(string(output), Version) {
		t.Fatalf(
			"expected output to contain version %q, got:\n%s",
			Version,
			string(output),
		)
	}
}

func TestHelperProcessVersion(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	os.Args = []string{
		"vrp-evaluator",
		"version",
	}

	main()

	os.Exit(0)
}

func TestMainWithoutArguments(t *testing.T) {
	cmd := exec.Command(os.Args[0], "-test.run=TestHelperProcessNoArguments")

	cmd.Env = append(
		os.Environ(),
		"GO_WANT_HELPER_PROCESS=1",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("helper process failed: %v", err)
	}

	if !strings.Contains(string(output), "No command specified.") {
		t.Fatalf(
			"unexpected output:\n%s",
			string(output),
		)
	}
}

func TestHelperProcessNoArguments(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	os.Args = []string{
		"vrp-evaluator",
	}

	main()

	os.Exit(0)
}
