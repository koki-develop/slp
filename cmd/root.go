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
		var d time.Duration
		if sec, err := strconv.ParseFloat(args[0], 64); err == nil {
			d = time.Duration(sec * float64(time.Second))
		} else {
			d, err = time.ParseDuration(args[0])
			if err != nil {
				return err
			}
		}

		cfg := modelConfig{Duration: d}
		if cmd.Flags().Changed("color") {
			cfg.Color = flagColor
		} else {
			if len(flagGradient) != 2 {
				return errors.New("gradient must have only two colors")
			}
			cfg.Gradient = flagGradient
		}

		m := newModel(cfg)
		p := tea.NewProgram(m, tea.WithOutput(os.Stderr))

		if _, err := p.Run(); err != nil {
			return err
		}
		if m.abort {
			os.Exit(130)
		}

		if flagBeep {
			fmt.Fprint(os.Stderr, "\a")
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

	rootCmd.Flags().BoolVarP(&flagBeep, "beep", "b", false, "beep when finished sleeping")

	rootCmd.Flags().StringVar(&flagColor, "color", "", "color of progress bar")
	rootCmd.Flags().StringSliceVar(&flagGradient, "gradient", []string{"#005B72", "#83E6FF"}, "apply a gradient between the two colors")
	rootCmd.MarkFlagsMutuallyExclusive("color", "gradient")
}
