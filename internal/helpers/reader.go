package helpers

import (
	"bufio"
	"fmt"
	"io"
)

// LinesFromReader reads all lines from the provided io.Reader and returns them as a slice of strings.
// Returns an error if reading fails.
func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if s.Err() != nil {
		return nil, fmt.Errorf("failed to scan reader: %w", s.Err())
	}

	return lines, nil
}
