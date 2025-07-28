package bundler

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBundler_New(t *testing.T) {
	bundler := New()
	if bundler == nil {
		t.Fatal("Expected bundler to be created")
	}

	if bundler.config.Mode != "development" {
		t.Errorf("Expected default mode to be 'development', got %s", bundler.config.Mode)
	}
}

func TestBundler_Bundle(t *testing.T) {
	// Create a temporary test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.js")
	testContent := "console.log('test');"

	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create bundler and set output to temp directory
	bundler := New()
	bundler.SetOutputDir(tmpDir)

	// Test bundling
	if err := bundler.Bundle(testFile); err != nil {
		t.Fatalf("Bundle failed: %v", err)
	}

	// Check if output file was created
	outputFile := filepath.Join(tmpDir, "bundle.js")
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Fatal("Output bundle file was not created")
	}

	// Verify content
	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	if len(content) == 0 {
		t.Fatal("Output file is empty")
	}
}
