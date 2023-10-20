package main

import (
	"os"
	"path"
)

func Env() (configPath string) {
	env := os.Getenv("FIREBOLT_MANAGER_V2_ENV")
	base, _ := os.Getwd()

	switch env {
	case "prod":
		configPath = path.Join(base, "/configs/prod.json")
	case "test":
		configPath = path.Join(base, "/configs/test.json")
	default:
		configPath = path.Join(base, "/configs/dev.json")
	}

	return
}
