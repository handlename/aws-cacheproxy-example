package main

import "github.com/kayac/go-config"

type Config struct {
	Secret       string               `yaml:"secret"`
	AllowedHosts []ConfigAllowedHost `yaml:"allowedHosts"`
}

type ConfigAllowedHost struct {
	// Name is a name of the allowed host.
	Name string `yaml:"name"`

	// Paths is a list of paths that are allowed to be accessed.
	// If a path contains `*`, it is treated as a wildcard.
	Paths []string `yaml:"paths"`
}

func loadConfig(path string) (*Config, error) {
	var conf Config
	err := config.Load(&conf, path)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
