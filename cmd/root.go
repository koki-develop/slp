package cmd

import (
	"errors"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "slp",
	Short: "sleep command with rich progress bar",
	Long:  "sleep command with rich progress bar.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		d, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return err
		}

		m := newModel(d)
		p := tea.NewProgram(m)

		if _, err = p.Run(); err != nil {
			return err
		}
		if m.abort {
			return errors.New("abort")
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
