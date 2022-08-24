package cli

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/blazejsewera/notipie/producer/cmd/nnp/internal/wire"
	"github.com/spf13/cobra"
	"os"
)

var (
	defaultFileFlag         *bool
	notificationFilePathArg *string
)

var (
	appNameArg        *string
	appIdDebugArg     *string
	appImgURIArg      *string
	titleArg          *string
	subtitleArg       *string
	bodyArg           *string
	extURIArg         *string
	readURIArg        *string
	archiveURIArg     *string
	timestampDebugArg *string
	apiKeyArg         *string
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push a notification to the Notipie backend",
	Run: func(cmd *cobra.Command, args []string) {
		producer := wire.GetProducer(*configPathArg, wire.ConfigOf(*addressArg, *portArg, *appIdArg))
		notification, err := wire.AppNotificationFrom(wire.AppNotificationConfig{
			UseDefaultNotificationFile: *defaultFileFlag,
			NotificationFilePath:       *notificationFilePathArg,
			PartialNotification: model.AppNotification{
				HashableNetNotification: model.HashableNetNotification{
					Timestamp:  *timestampDebugArg,
					AppName:    *appNameArg,
					AppID:      *appIdDebugArg,
					AppImgURI:  *appImgURIArg,
					Title:      *titleArg,
					Subtitle:   *subtitleArg,
					Body:       *bodyArg,
					ExtURI:     *extURIArg,
					ReadURI:    *readURIArg,
					ArchiveURI: *archiveURIArg,
				},
				ApiKey: *apiKeyArg,
			},
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, "retrieve notification:", err)
			os.Exit(1)
		}

		err = producer.Push(notification)
		if err != nil {
			fmt.Fprintln(os.Stderr, "push:", err)
			os.Exit(1)
		}
	},
}

func SetupPush() {
	defaultFileFlag = pushCmd.Flags().BoolP("default-file", "d", false, "use a default file for notification (<config_dir>/notipie/producer/notification.yaml)")
	notificationFilePathArg = pushCmd.Flags().StringP("file", "f", "", "set a custom file path for notification")
	appNameArg = pushCmd.Flags().String("app-name", "", "set a custom app name")
	appIdDebugArg = pushCmd.Flags().String("app-id-debug-flag", "", "set a custom app id (for debugging)")
	appImgURIArg = pushCmd.Flags().String("app-img-uri", "", "set a custom icon")
	titleArg = pushCmd.Flags().StringP("title", "t", "", "set a custom title for the notification")
	subtitleArg = pushCmd.Flags().StringP("subtitle", "s", "", "set a custom subtitle for the notification")
	bodyArg = pushCmd.Flags().StringP("body", "b", "", "set a custom body for the notification")
	extURIArg = pushCmd.Flags().String("ext-uri", "", "set a custom external link uri")
	readURIArg = pushCmd.Flags().String("read-uri", "", "set a custom read link uri (to mark a notification as read in an external service)")
	archiveURIArg = pushCmd.Flags().String("archive-uri", "", "set a custom archive link uri (to mark a notification as archived in an external service)")
	timestampDebugArg = pushCmd.Flags().String("timestamp-debug-flag", "", "set a custom timestamp (for debugging)")
	apiKeyArg = pushCmd.Flags().String("api-key", "", "set a custom API key")
}
