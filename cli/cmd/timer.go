package cmd

import (
	"github.com/devbyP/pt/cli/timer"

	"github.com/spf13/cobra"
)

// timerCmd represents the timer command
var timerCmd = &cobra.Command{
	Use:   "timer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		timer.StartTimer()
	},
}

func init() {
	tFlag := timerCmd.Flags()
	_ = tFlag.BoolP("countdown", "c", false, "set timer to countdown")
	_ = tFlag.StringP("time", "t", "5m", "countdown timer in second(s) or minute(m)")
	rootCmd.AddCommand(timerCmd)
}
