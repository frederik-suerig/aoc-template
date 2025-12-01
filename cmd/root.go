package cmd

import (
	"os"

	"github.com/frederik-suerig/advent-of-code/internal/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "adventofcode",
	Aliases:       []string{"aoc"},
	Short:         "A CLI for Advent of Code",
	Long:          `Advent of Code is a series of programming puzzles. This CLI helps you with creating scaffolding and helper functions for the puzzles.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// Extract the underlying error message for cleaner output
		ui.Error("%s", err)
		os.Exit(1)
	}
}
