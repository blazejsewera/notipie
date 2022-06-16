package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping the Notipie backend",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ping called")
	},
}
