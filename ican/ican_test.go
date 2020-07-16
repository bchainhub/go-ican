package ican

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Test fake icans
func TestFake(t *testing.T) {
	fakeICANs := []string{
		"VG96VPVG00000L2345678901",
		"1234567890",
		"12345678901234567890",
		"NL30ABNA0123456789",
		"NL30ABNA0517552264",
		"NL30ABNA05175522AB",
	}

	// Loop through fake icans, they should all raise an error
	for _, fake := range fakeICANs {
		ican, err := NewICAN(fake)
		if err == nil {
			// Fake ican did not raise an error,
			t.Errorf("ICAN fake test error: %v", ican.Code)
		}
	}
}

// Test real icans
func TestICANS(t *testing.T) {
	err := filepath.Walk("example-icans", func(path string, f os.FileInfo, err error) error {
		if f.IsDir() || strings.HasPrefix(f.Name(), ".") {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		t.Logf("Start ICAN file test %v\n", path)

		for scanner.Scan() {
			line := scanner.Text()
			t.Logf("ICAN code input: %v\n", line)
			ican, icanErr := NewICAN(line)
			if icanErr != nil {
				return icanErr
			}
			t.Logf("ICAN code validated: %v\n", ican.PrintCode)
		}

		if err := scanner.Err(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		t.Errorf("ICAN test error: %v", err)
	}
}
