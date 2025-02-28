package cmd

import (
	"fmt"
	"mgr/internal/config"
	"mgr/internal/tui"
	"mgr/internal/utils"

	"github.com/spf13/cobra"
)

func GoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go [alias] [script]",
		Short: "Executa um script de um projeto via alias",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dryRun, _ := cmd.Flags().GetBool("dry-run")
			utils.DryRun = dryRun

			alias := args[0]

			projectPath, err := config.GetAlias(alias)
			if err != nil {
				fmt.Println(err)
				return
			}

			if len(args) == 1 {
				fmt.Println("\n Abrindo seleção de scripts para:", projectPath)
				scripts := config.GetScriptsFromDB(projectPath)
				if len(scripts) == 0 {
					fmt.Println("\n Nenhum script encontrado para esse projeto.")
					return
				}
				tui.RunTUI(projectPath, scripts)
				return
			}

			script := args[1]
			fmt.Printf("\n Executando '%s' em %s...\n", script, projectPath)
			utils.RunScript(projectPath, script)
		},
	}

	cmd.Flags().Bool("dry-run", false, "Simula a execução do script sem realmente rodá-lo")
	return cmd
}
