package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/cmd/gateway"
)

func main() {
	c := &cobra.Command{Use: "server-performance-tuning-2023 [command]"}
	gateway.RegisterCommand(c)
	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
