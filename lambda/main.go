package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fujiwara/ridge"
	"go.uber.org/zap"
)

func main() {
	initLogger(os.Getenv("LOG_LEVEL"))
	defer logger.Sync()

	// TODO: use flag
	configPath := "config.yaml"
	port := 8080

	conf, err := loadConfig(configPath)
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}
	logger.Debug("loaded config", zap.Any("config", conf))

	server := NewServer(conf)

	var mux = http.NewServeMux()
	mux.HandleFunc("/", server.Handler)

	logger.Info("start server", zap.String("listen", fmt.Sprintf("http://localhost:%d", port)))
	ridge.Run(fmt.Sprintf(":%d", port), "/", mux)
}
