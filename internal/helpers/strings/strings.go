package strings

import (
	"strconv"
	"strings"
	"unicode"
)

// Reverse returns a reversed copy of the string.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ExtractInts extracts all integers from a string and returns them as a slice.
// Handles negative numbers and ignores non-numeric characters.
func ExtractInts(s string) []int {
	var result []int
	var current strings.Builder
	var isNegative bool

	for _, r := range s {
		if r == '-' && current.Len() == 0 {
			isNegative = true
		} else if unicode.IsDigit(r) {
			current.WriteRune(r)
		} else {
			if current.Len() > 0 {
				val, _ := strconv.Atoi(current.String())
				if isNegative {
					val = -val
				}
				result = append(result, val)
				current.Reset()
				isNegative = false
			}
		}
	}

	// Handle number at end of string
	if current.Len() > 0 {
		val, _ := strconv.Atoi(current.String())
		if isNegative {
			val = -val
		}
		result = append(result, val)
	}

	return result
}

// RemoveAll removes all occurrences of the specified substring from the string.
func RemoveAll(s, substr string) string {
	return strings.ReplaceAll(s, substr, "")
}

// SplitByLength splits a string into chunks of the specified length.
func SplitByLength(s string, length int) []string {
	if length <= 0 {
		return []string{s}
	}

	var result []string
	for i := 0; i < len(s); i += length {
		end := i + length
		if end > len(s) {
			end = len(s)
		}
		result = append(result, s[i:end])
	}
	return result
}

// IsNumeric checks if the string contains only numeric characters.
func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// SplitByEmptyLine splits input into groups separated by empty lines.
// Useful for AoC problems with multi-line groups.
func SplitByEmptyLine(lines []string) [][]string {
	var groups [][]string
	var current []string

	for _, line := range lines {
		if line == "" {
			if len(current) > 0 {
				groups = append(groups, current)
				current = nil
			}
		} else {
			current = append(current, line)
		}
	}
	if len(current) > 0 {
		groups = append(groups, current)
	}

	return groups
}

