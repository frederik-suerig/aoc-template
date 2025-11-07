package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "adventofcode",
	Aliases: []string{"aoc"},
	Short:   "A CLI for Advent of Code",
	Long:    `Advent of Code is a series of programming puzzles. This CLI helps you with creating scaffolding and helper functions for the puzzles.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
