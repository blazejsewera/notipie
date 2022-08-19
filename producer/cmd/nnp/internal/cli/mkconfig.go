package cli

import (
	"fmt"
	"github.com/blazejsewera/notipie/producer/cmd/nnp/internal/wire"
	"github.com/blazejsewera/notipie/producer/pkg/lib/config"
	"github.com/spf13/cobra"
	"os"
)

const shortDescription = "Make a configuration file. If no custom config path is specified, the default one is used."

var longDescription = fmt.Sprintf(`Make a configuration file.
If no custom config path is specified, the default one is used.
Default path: %s`, wire.DefaultProducerConfigFilePath)

var mkconfigCmd = &cobra.Command{
	Use:   "mkconfig",
	Short: shortDescription,
	Long:  longDescription,
	Run: func(cmd *cobra.Command, args []string) {
		path := wire.DefaultProducerConfigFilePath
		if *configPathArg != "" {
			path = *configPathArg
		}
		wire.SetConfigPath(path)
		cfg := wire.ConfigOf(*addressArg, *portArg, *appIdArg)
		checkConfig(cfg)

		err := wire.MkConfigDirIfDoesNotExist()
		if err != nil {
			handleMkconfigError(err)
		}
		err = wire.SaveConfig(cfg)
		if err != nil {
			handleMkconfigError(err)
		}
	},
}

func checkConfig(cfg config.Config) {
	if cfg.Address == "" {
		fmt.Fprintln(os.Stderr, "address cannot be unset, set with -a <address> flag")
		os.Exit(1)
	}
	if cfg.Port == 0 {
		fmt.Fprintln(os.Stderr, "port cannot be unset, set with -p <port> flag")
		os.Exit(1)
	}
}

func handleMkconfigError(err error) {
	fmt.Fprintln(os.Stderr, "mkconfig:", err)
	os.Exit(1)
}
