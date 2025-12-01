package create

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
)

// Domain-specific errors
var (
	ErrFileExists      = errors.New("file already exists")
	ErrInvalidDay      = errors.New("invalid day")
	ErrInvalidYear     = errors.New("invalid year")
	ErrWorkdirRequired = errors.New("workdir is required")
	ErrCookieRequired  = errors.New("cookie is required")
)

// FileExistsError represents a file that already exists
type FileExistsError struct {
	Path string
}

func (e *FileExistsError) Error() string {
	return fmt.Sprintf("%s: %s", ErrFileExists, e.Path)
}

// UserMessage returns a user-friendly error message
func (e *FileExistsError) UserMessage() string {
	year, day := extractYearDayFromPath(e.Path)
	if year > 0 && day > 0 {
		return fmt.Sprintf("day %d of %d already exists", day, year)
	}
	// Fallback if we can't parse the path
	return fmt.Sprintf("file already exists: %s", filepath.Base(e.Path))
}

// extractYearDayFromPath extracts year and day from a path like .../y2024/d10/...
func extractYearDayFromPath(path string) (year, day int) {
	// Match patterns like y2024 or y24, and d10 or d01
	// Handle both absolute paths (/path/y2024/d10) and relative paths (y2024/d10 or ./y2024/d10)
	yearRegex := regexp.MustCompile(`(?:^|[\\/])y(\d{4}|\d{2})(?:[\\/]|$)`)
	dayRegex := regexp.MustCompile(`(?:^|[\\/])d(\d{1,2})(?:[\\/]|$)`)

	yearMatch := yearRegex.FindStringSubmatch(path)
	if len(yearMatch) > 1 {
		yearStr := yearMatch[1]
		if parsedYear, err := strconv.Atoi(yearStr); err == nil {
			// If 2 digits, prepend 20; if 4 digits, use as-is
			if len(yearStr) == 2 {
				year = 2000 + parsedYear
			} else {
				year = parsedYear
			}
		}
	}

	dayMatch := dayRegex.FindStringSubmatch(path)
	if len(dayMatch) > 1 {
		if parsedDay, err := strconv.Atoi(dayMatch[1]); err == nil {
			day = parsedDay
		}
	}

	return year, day
}

// DownloadError represents an error downloading input
type DownloadError struct {
	Reason string
	Status int
}

func (e *DownloadError) Error() string {
	if e.Status > 0 {
		return fmt.Sprintf("failed to download input: %s (status %d)", e.Reason, e.Status)
	}
	return fmt.Sprintf("failed to download input: %s", e.Reason)
}

// NewFileExistsError creates a new FileExistsError
func NewFileExistsError(path string) *FileExistsError {
	return &FileExistsError{Path: path}
}

// NewDownloadError creates a new DownloadError
func NewDownloadError(reason string, status int) *DownloadError {
	return &DownloadError{Reason: reason, Status: status}
}
