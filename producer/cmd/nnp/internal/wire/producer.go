package wire

import (
	"fmt"
	"github.com/blazejsewera/notipie/producer/pkg/lib/config"
	"github.com/blazejsewera/notipie/producer/pkg/lib/nnp"
	"os"
)

var producerInstance nnp.Producer

func GetProducer(baseConfigPath string, patch config.Config) nnp.Producer {
	if producerInstance != nil {
		return producerInstance
	}

	cfg, err := GetProducerConfig(baseConfigPath, patch)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "read config:", err)
		os.Exit(2)
	}

	producerInstance = nnp.NewProducer(cfg, nnp.AppIDSaverFunc(func(appID string) error {
		fmt.Println("appID:", appID)
		return nil
	}))
	return producerInstance
}
