package cli

import (
	"fmt"
	"github.com/blazejsewera/notipie/producer/cmd/nnp/internal/wire"
	"github.com/spf13/cobra"
	"os"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push a notification to the Notipie backend",
	Run: func(cmd *cobra.Command, args []string) {
		producer := wire.GetProducer()
		notification, err := wire.GetAppNotification()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "retrieve notification:", err)
			os.Exit(2)
		}

		err = producer.Push(notification)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "ping:", err)
			os.Exit(2)
		}
	},
}
