// Package testy implements some helpers for writing tests.
package testy

import (
	"fmt"
	"os"
)

// CopyFile copies file.
func CopyFile(src string, dest string) (err error) {
	input, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	if err = os.WriteFile(dest, input, 0644); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

// FileSize returns the file size.
func FileSize(name string) (size int64, err error) {
	file, err := os.Open(name)
	if err != nil {
		return size, fmt.Errorf("open file: %w", err)
	}

	stats, err := file.Stat()
	if err != nil {
		return size, fmt.Errorf("get file stats: %w", err)
	}

	size = stats.Size()

	if err = file.Close(); err != nil {
		return size, fmt.Errorf("close file: %w", err)
	}

	return size, nil
}
