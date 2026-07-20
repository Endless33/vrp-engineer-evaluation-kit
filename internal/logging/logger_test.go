package logging

import (
	"bytes"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	logger := New()

	if logger == nil {
		t.Fatal("expected logger")
	}

	if logger.logger == nil {
		t.Fatal("expected underlying logger")
	}
}

func TestNewWithWriter(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	logger := NewWithWriter(buffer)

	logger.Info("hello %s", "world")

	output := buffer.String()

	if output == "" {
		t.Fatal("expected log output")
	}

	if !strings.Contains(output, "[INFO]") {
		t.Fatal("expected INFO prefix")
	}

	if !strings.Contains(output, "hello world") {
		t.Fatal("expected formatted message")
	}
}

func TestWarning(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	logger := NewWithWriter(buffer)

	logger.Warning("warning message")

	if !strings.Contains(buffer.String(), "[WARNING]") {
		t.Fatal("expected WARNING prefix")
	}
}

func TestError(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	logger := NewWithWriter(buffer)

	logger.Error("error message")

	if !strings.Contains(buffer.String(), "[ERROR]") {
		t.Fatal("expected ERROR prefix")
	}
}

func TestSection(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	logger := NewWithWriter(buffer)

	logger.Section("Evaluation")

	output := buffer.String()

	if !strings.Contains(output, "Evaluation") {
		t.Fatal("expected section title")
	}

	if !strings.Contains(output, "========================================") {
		t.Fatal("expected section separator")
	}
}

func TestNilLoggerDoesNotPanic(t *testing.T) {
	var logger *Logger

	logger.Info("info")
	logger.Warning("warning")
	logger.Error("error")
	logger.Section("section")
}