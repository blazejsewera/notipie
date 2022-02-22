package log

import (
	"fmt"
	"go.uber.org/zap"
)

func Init(prod bool) {
	var logger *zap.Logger
	var err error

	if prod {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(fmt.Sprintf("could not initialize logger, err: %s", err))
	}

	zap.ReplaceGlobals(logger)
}

func For(name string) *zap.Logger {
	return zap.L().Named(name)
}
