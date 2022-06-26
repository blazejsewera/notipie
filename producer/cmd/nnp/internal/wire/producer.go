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

	SetConfigPath(baseConfigPath)
	cfg, err := GetProducerConfig(patch)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read config:", err)
		os.Exit(2)
	}

	producerInstance = nnp.NewProducer(cfg, AppIDSaver)
	return producerInstance
}
