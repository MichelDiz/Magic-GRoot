package cmd

import (
	"encoding/json"
	"fmt"
	"mgr/internal/config"
	"mgr/internal/tui"

	"github.com/spf13/cobra"
)

type ProjectData struct {
	Name    string
	Scripts map[string]string
}

func ListCmd() *cobra.Command {
	var listByPath bool
	var listAll bool

	cmd := &cobra.Command{
		Use:   "ls",
		Short: "Lista projetos disponíveis",
		Run: func(cmd *cobra.Command, args []string) {
			projectsRaw := config.GetProjectsFromDB()
			if len(projectsRaw) == 0 {
				fmt.Println("\n Nenhum projeto encontrado no banco de dados.")
				fmt.Println("Execute 'mgr scan' para atualizar os projetos.")
				return
			}

			projects := make(map[string]ProjectData)
			for path, rawData := range projectsRaw {
				if len(rawData) < 2 {
					fmt.Printf("\nErro ao processar projeto em %s: dados incompletos.\n", path)
					continue
				}

				var project ProjectData
				_ = json.Unmarshal([]byte(rawData[0]), &project.Name)
				project.Scripts = make(map[string]string)
				_ = json.Unmarshal([]byte(rawData[1]), &project.Scripts)
				projects[path] = project
			}

			if listAll {
				for path, project := range projects {
					fmt.Printf("\nNome: %s\nPath: %s\nScripts:\n", project.Name, path)
					for script, command := range project.Scripts {
						fmt.Printf("  - %s: %s\n", script, command)
					}
				}
				return
			}

			var items []string
			for path, project := range projects {
				if listByPath {
					items = append(items, path)
				} else {
					if project.Name != "" {
						items = append(items, project.Name)
					} else {
						items = append(items, path) // Usa o path como fallback
					}
				}
			}

			tui.RunListTUI(items)
		},
	}

	cmd.Flags().BoolVarP(&listByPath, "path", "p", false, "Lista os projetos pelo caminho em vez do nome")
	cmd.Flags().BoolVarP(&listAll, "all", "a", false, "Lista todos os detalhes dos projetos")

	return cmd
}
