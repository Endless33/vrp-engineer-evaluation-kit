package logging

import (
	"io"
	"log"
	"os"
)

// Logger wraps the standard library logger used by the
// public engineering evaluation toolkit.
//
// This package intentionally provides only generic logging.
// It must never log protected runtime state, proprietary
// algorithms, confidential configuration, authentication
// secrets, or private protocol internals.
type Logger struct {
	logger *log.Logger
}

// New returns a logger writing to stdout.
func New() *Logger {
	return &Logger{
		logger: log.New(
			os.Stdout,
			"[vrp-evaluator] ",
			log.LstdFlags|log.LUTC,
		),
	}
}

// NewWithWriter returns a logger writing to the supplied writer.
func NewWithWriter(w io.Writer) *Logger {
	return &Logger{
		logger: log.New(
			w,
			"[vrp-evaluator] ",
			log.LstdFlags|log.LUTC,
		),
	}
}

// Info writes an informational message.
func (l *Logger) Info(format string, args ...any) {
	if l == nil || l.logger == nil {
		return
	}

	l.logger.Printf("[INFO] "+format, args...)
}

// Warning writes a warning message.
func (l *Logger) Warning(format string, args ...any) {
	if l == nil || l.logger == nil {
		return
	}

	l.logger.Printf("[WARNING] "+format, args...)
}

// Error writes an error message.
func (l *Logger) Error(format string, args ...any) {
	if l == nil || l.logger == nil {
		return
	}

	l.logger.Printf("[ERROR] "+format, args...)
}

// Section writes a section header.
func (l *Logger) Section(title string) {
	if l == nil || l.logger == nil {
		return
	}

	l.logger.Println("========================================")
	l.logger.Println(title)
	l.logger.Println("========================================")
}