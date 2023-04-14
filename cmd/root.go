package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var (
	flagBeep bool
)

var rootCmd = &cobra.Command{
	Use:          "slp [duration]",
	Short:        "sleep command with rich progress bar",
	Long:         "sleep command with rich progress bar.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
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

		if flagBeep {
			fmt.Print("\a")
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

func init() {
	rootCmd.Flags().SortFlags = false

	rootCmd.Flags().BoolVarP(&flagBeep, "beep", "b", false, "beep when finished sleeping")
}
