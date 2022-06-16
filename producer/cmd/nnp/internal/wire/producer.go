package wire

import (
	"fmt"
	"github.com/blazejsewera/notipie/producer/pkg/lib/nnp"
	"os"
)

var producerInstance nnp.Producer

func init() {
	if producerInstance != nil {
		return
	}

	cfg, err := getProducerConfig()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "read config:", err)
		os.Exit(2)
	}

	producerInstance = nnp.NewProducer(cfg, nnp.AppIDSaverFunc(func(appID string) error {
		fmt.Println("appID:", appID)
		return nil
	}))
}

func GetProducer() nnp.Producer {
	return producerInstance
}
