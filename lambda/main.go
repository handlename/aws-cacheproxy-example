package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/fujiwara/ridge"
	"go.uber.org/zap"
)

type Flags struct {
	ConfigPath string
	Port       int
}

func main() {
	initLogger(os.Getenv("LOG_LEVEL"))
	defer logger.Sync()

	f := Flags{}
	flag.StringVar(&f.ConfigPath, "config", "config.yaml", "config file path")
	flag.IntVar(&f.Port, "port", 8080, "port number")
	flag.Parse()

	conf, err := loadConfig(f.ConfigPath)
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}
	logger.Debug("loaded config", zap.Any("config", conf))

	server := NewServer(conf)

	var mux = http.NewServeMux()
	mux.HandleFunc("/", server.Handler)

	logger.Info("start server", zap.String("listen", fmt.Sprintf("http://localhost:%d", f.Port)))
	ridge.Run(fmt.Sprintf(":%d", f.Port), "/", mux)
}
