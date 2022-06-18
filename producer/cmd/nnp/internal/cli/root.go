package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nnp",
	Short: "Notipie Notification Producer",
	Long: `A notification producer (and library) that can send
notifications specified in JSON, YAML, or through CLI arguments.`,
}

func Setup() {
	rootCmd.AddCommand(pingCmd)
	rootCmd.AddCommand(pushCmd)
	SetupPush()
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
