package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var (
	flagSecond bool
	flagMinute bool
	flagHour   bool
	flagBeep   bool
)

var rootCmd = &cobra.Command{
	Use:          "slp [time]",
	Short:        "sleep command with rich progress bar",
	Long:         "sleep command with rich progress bar.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return err
		}

		base := time.Second
		switch {
		case flagSecond:
			base = time.Second
		case flagMinute:
			base = time.Minute
		case flagHour:
			base = time.Hour
		}

		m := newModel(modelConfig{
			Duration: time.Duration(t * float64(base)),
		})
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

	rootCmd.Flags().BoolVar(&flagSecond, "second", false, "set the time unit to seconds (default)")
	rootCmd.Flags().BoolVar(&flagMinute, "minute", false, "set the time unit to minutes")
	rootCmd.Flags().BoolVar(&flagHour, "hour", false, "set the time unit to hours")

	rootCmd.Flags().BoolVarP(&flagBeep, "beep", "b", false, "beep when finished sleeping")
}
