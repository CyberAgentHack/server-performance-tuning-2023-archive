package gateway

import (
	"github.com/spf13/cobra"

	v1 "github.com/CyberAgentHack/server-performance-tuning-2023/pkg/cmd/gateway/v1"
)

func RegisterCommand(registry *cobra.Command) {
	registry.AddCommand(
		v1.New(),
	)
}
