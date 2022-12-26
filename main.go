package main

import (
	"log"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Failed to run app %v", err)
		return
	}
}

func run() error {
	a, err := app.New()
	if err != nil {
		return err
	}
	return a.Run()
}
