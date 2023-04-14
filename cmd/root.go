package cmd

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var (
	version string

	flagSecond   bool
	flagMinute   bool
	flagHour     bool
	flagBeep     bool
	flagColor    string
	flagGradient []string
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

		cfg := modelConfig{Duration: time.Duration(t * float64(base))}
		if cmd.Flags().Changed("color") {
			cfg.Color = flagColor
		} else {
			if len(flagGradient) != 2 {
				return errors.New("gradient must have only two colors")
			}
			cfg.Gradient = flagGradient
		}

		m := newModel(cfg)
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
	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
	}
	rootCmd.Version = version

	rootCmd.Flags().SortFlags = false

	rootCmd.Flags().BoolVar(&flagSecond, "second", false, "set the time unit to seconds (default)")
	rootCmd.Flags().BoolVar(&flagMinute, "minute", false, "set the time unit to minutes")
	rootCmd.Flags().BoolVar(&flagHour, "hour", false, "set the time unit to hours")
	rootCmd.MarkFlagsMutuallyExclusive("second", "minute", "hour")

	rootCmd.Flags().BoolVarP(&flagBeep, "beep", "b", false, "beep when finished sleeping")

	rootCmd.Flags().StringVar(&flagColor, "color", "", "color of progress bar")
	rootCmd.Flags().StringSliceVar(&flagGradient, "gradient", []string{"#005B72", "#1DD2FF"}, "apply a gradient between the two colors")
	rootCmd.MarkFlagsMutuallyExclusive("color", "gradient")
}
