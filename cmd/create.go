package cmd

import (
	"fmt"
	"time"

	"github.com/frederik-suerig/advent-of-code/internal/create"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Generate code for a new day",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		g, err := create.NewGenerator(
			viper.GetInt("day"),
			viper.GetInt("year"),
			viper.GetString("workdir"),
			viper.GetString("cookie"))
		if err != nil {
			return fmt.Errorf("could not create generator: %w", err)
		}

		if err := g.Run(); err != nil {
			return fmt.Errorf("could not run generator: %w", err)
		}

		fmt.Println("Files created!")

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

	viper.BindPFlags(createCmd.Flags())
}
