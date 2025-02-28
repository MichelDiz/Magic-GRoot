package cmd

import (
	"fmt"
	"mgr/internal/config"
	"mgr/internal/utils"

	"github.com/spf13/cobra"
)

func SetRunnerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set-runner [npm|yarn|pnpm|bash]",
		Short: "Define o gerenciador de scripts preferido",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner := args[0]

			if !utils.IsCommandAvailable(runner) {
				fmt.Printf(" O gerenciador '%s' não está instalado no sistema.\n", runner)
				return
			}

			config.SetConfig("preferred_runner", runner)
			fmt.Printf(" Gerenciador de scripts definido como: %s\n", runner)
		},
	}
}
