package cmd

import (
	"fmt"
	"mgr/internal/config"
	"mgr/internal/tui"

	"github.com/spf13/cobra"
)

func AliasCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alias",
		Short: "Gerencia aliases para projetos",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "add",
		Short: "Adiciona um novo alias para um projeto interativamente",
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

			tui.RunAliasTUI(projectPaths)
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "ls",
		Short: "Lista todos os aliases salvos",
		Run: func(cmd *cobra.Command, args []string) {
			aliases := config.GetAllAliases()
			if len(aliases) == 0 {
				fmt.Println(" Nenhum alias encontrado. Use 'mgr alias add' para adicionar um.")
				return
			}
			fmt.Println("\n Aliases registrados:")
			for alias, path := range aliases {
				fmt.Printf("  %s → %s\n", alias, path)
			}
		},
	})

	return cmd
}
