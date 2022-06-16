package cli

import (
	"fmt"
	"github.com/blazejsewera/notipie/producer/cmd/nnp/internal/wire"
	"github.com/spf13/cobra"
	"os"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping the Notipie backend",
	Run: func(cmd *cobra.Command, args []string) {
		producer := wire.GetProducer()

		err := producer.Ping()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "ping:", err)
			os.Exit(3)
		}
	},
}