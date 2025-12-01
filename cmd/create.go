package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/frederik-suerig/advent-of-code/internal/create"
	"github.com/frederik-suerig/advent-of-code/internal/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Generate code for a new day",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := create.Config{
			Year:    viper.GetInt("year"),
			Day:     viper.GetInt("day"),
			WorkDir: viper.GetString("workdir"),
			Cookie:  viper.GetString("cookie"),
		}

		g, err := create.NewGenerator(cfg)
		if err != nil {
			return formatError(err)
		}

		if err := g.Run(); err != nil {
			return formatError(err)
		}

		ui.Success("All files created successfully!")
		ui.HighlightInfo("You can now start solving the puzzle in: ./y%04d/d%02d", cfg.Year, cfg.Day)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Year defaults to latest Advent of Code.
	year, month, day := time.Now().Date()
	if month < time.December {
		year--
	}

	createCmd.Flags().IntP("day", "d", day, "The day to build scaffolding for")
	createCmd.Flags().IntP("year", "y", year, "The year of Advent of Code you are working on")

	createCmd.Flags().StringP("workdir", "w", "", "Your Advent of Code working directory")
	createCmd.Flags().StringP("cookie", "c", "", "Your session cookie for adventofcode.com")

	if err := viper.BindPFlags(createCmd.Flags()); err != nil {
		panic(fmt.Errorf("failed to bind flags: %w", err))
	}
}

// formatError formats errors for user-friendly display
func formatError(err error) error {
	if err == nil {
		return nil
	}

	// Handle cancellation message - don't format it, just return as-is
	errMsg := err.Error()
	if errMsg == "operation cancelled by user" {
		return err
	}

	// Handle custom error types
	var fileExistsErr *create.FileExistsError
	if errors.As(err, &fileExistsErr) {
		return fmt.Errorf("%s", fileExistsErr.UserMessage())
	}

	var downloadErr *create.DownloadError
	if errors.As(err, &downloadErr) {
		// For download errors, return the reason without status code for cleaner output
		return fmt.Errorf("%s", downloadErr.Reason)
	}

	// For other errors, return as-is (they're already user-friendly)
	return err
}
