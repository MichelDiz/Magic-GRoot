package cmd

import (
	"fmt"
	"mgr/internal/config"
	"mgr/internal/tui"

	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "Lista projetos disponíveis",
		Run: func(cmd *cobra.Command, args []string) {
			projects := config.GetProjectsFromDB()
			if len(projects) == 0 {
				fmt.Println("\n Nenhum projeto encontrado no banco de dados.")
				fmt.Println("Execute 'mgr scan' para atualizar os projetos.")
				return
			}

			var projectPaths []string
			for projectPath := range projects {
				projectPaths = append(projectPaths, projectPath)
			}

			tui.RunListTUI(projectPaths)
		},
	}
}
