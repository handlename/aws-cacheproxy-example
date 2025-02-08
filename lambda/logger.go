package main

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func initLogger(levelString string) {
	logger, _ = zap.NewProduction()

	level := zap.InfoLevel
	if levelString != "" {
		if parsed, err := zap.ParseAtomicLevel(levelString); err != nil {
			logger.Warn("failed to parse log level", zap.Error(err))
		} else {
			level = parsed.Level()
		}
	}

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	if l, err := config.Build(); err != nil {
		panic(err)
	} else {
		logger = l
	}
}
