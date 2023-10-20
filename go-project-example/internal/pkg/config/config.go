package config

import (
	"encoding/json"
	"go-project-example/internal/pkg/log"
	"os"
	"path"
)

var c *Config

func Initialize(configPath string) (err error) {
	if c, err = static(configPath); err != nil {
		return
	}
	env(c)

	return
}

func static(configPath string) (conf *Config, err error) {
	var f []byte
	f, err = os.ReadFile(configPath)
	if err != nil {
		return
	}

	conf = new(Config)
	if err = json.Unmarshal(f, conf); err != nil {
		return
	}

	return
}

func env(config *Config) {
	var appLogDir = os.Getenv("MATRIX_APPLOGS_DIR")
	var accessLogDir = os.Getenv("MATRIX_ACCESSLOGS_DIR")

	if appLogDir != "" {
		(*config).Server.Log[log.AppLog] = path.Join(appLogDir, (*config).Server.Log[log.AppLog])
	}
	if accessLogDir != "" {
		(*config).Server.Log[log.AccessLog] = path.Join(accessLogDir, (*config).Server.Log[log.AccessLog])
	}
}

func Get() *Config {
	return c
}
