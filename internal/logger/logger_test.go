package internallogger_test

import (
	"context"
	"testing"

	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	internallogger "github.com/0ffsideCompass/footballmodules/internal/logger"
)

func TestInfoLog(t *testing.T) {
	logger, err := internallogger.NewLogger()
	if err != nil {
		t.Fatalf("Error creating logger: %v", err)
	}

	// Create a memory logger to capture log entries.
	core, recorded := observer.New(zapcore.InfoLevel)
	logger.SetCore(core)

	// Log a message.
	logger.Info(context.Background(), "Info Message", map[string]interface{}{"key": "value", "number": 1})

	// Check if the log entry was recorded.
	entries := recorded.All()
	if len(entries) != 1 {
		t.Fatalf("Expected 1 log entry, got %d", len(entries))
	}

	if entries[0].Message != "Info Message" {
		t.Errorf("Unexpected message: %s", entries[0].Message)
	}

	if len(entries[0].Context) != 2 {
		t.Fatalf("Expected 2 fields, got %d", len(entries[0].Context))
	}
}

func TestErrorLog(t *testing.T) {
	logger, err := internallogger.NewLogger()
	if err != nil {
		t.Fatalf("Error creating logger: %v", err)
	}

	// Create a memory logger to capture log entries.
	core, recorded := observer.New(zapcore.ErrorLevel)
	logger.SetCore(core)

	// Log an error.
	logger.Error(context.Background(), "Error Message", map[string]interface{}{"error_key": "error_value", "error_number": 2})

	// Check if the log entry was recorded.
	entries := recorded.All()
	if len(entries) != 1 {
		t.Fatalf("Expected 1 log entry, got %d", len(entries))
	}

	if entries[0].Message != "Error Message" {
		t.Errorf("Unexpected message: %s", entries[0].Message)
	}

	if len(entries[0].Context) != 2 {
		t.Fatalf("Expected 2 fields, got %d", len(entries[0].Context))
	}
}
