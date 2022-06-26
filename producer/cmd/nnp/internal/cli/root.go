package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	configPathArg *string
	appIdArg      *string
	addressArg    *string
	portArg       *int
)

var rootCmd = &cobra.Command{
	Use:   "nnp",
	Short: "Notipie Notification Producer",
	Long: `A notification producer (and library) that can send
notifications specified in JSON, YAML, or through CLI arguments.`,
}

func Setup() {
	setupRootFlags()

	rootCmd.AddCommand(mkconfigCmd)
	rootCmd.AddCommand(pingCmd)
	rootCmd.AddCommand(pushCmd)
	SetupPush()
}

func setupRootFlags() {
	configPathArg = rootCmd.PersistentFlags().StringP("config", "c", "", "set a custom config file path")
	appIdArg = rootCmd.PersistentFlags().StringP("app-id", "i", "", "set a custom app id")
	addressArg = rootCmd.PersistentFlags().StringP("address", "a", "", "set a custom address (hostname) for Notipie backend")
	portArg = rootCmd.PersistentFlags().IntP("port", "p", 0, "set a custom port for Notipie backend")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
