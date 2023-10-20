package main

import (
	"go-project-example/internal/app/router"
)

func main() {
	var err error
	var configPath string

	configPath = Env()
	if err = load(configPath); err != nil {
		panic(err)
	}

	var errCh = make(chan error)
	go func() {
		errCh <- router.Router()
	}()

	for err = range errCh {
		if err != nil {
			panic(err)
		}
	}
}
