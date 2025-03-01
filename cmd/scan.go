package cmd

import (
	"context"
	"fmt"
	"mgr/internal/config"
	"mgr/internal/scanner"
	"mgr/internal/tui"

	"github.com/spf13/cobra"
)

func ScanCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "scan",
		Short: config.Translate("scan_short"),
		Run: func(cmd *cobra.Command, args []string) {
			rootPath := config.GetConfig("root_path")
			if rootPath == "" {
				fmt.Println(config.Translate("root_not_set"))
				return
			}

			scripts, err := scanner.ScanForScripts(context.Background(), rootPath)
			if err != nil {
				fmt.Println("Erro ao buscar scripts:", err)
				return
			}

			if len(scripts) == 0 {
				fmt.Println("Nenhum script encontrado.")
				return
			}

			if len(scripts) > 1 {
				fmt.Println("Múltiplos projetos encontrados. Escolha um para visualizar os scripts:")
				for projectPath := range scripts {
					fmt.Println("- ", projectPath)
				}
				fmt.Println("Use 'mgr scan [projeto]' para ver scripts de um projeto específico.")
				return
			}

			for projectPath, scriptMap := range scripts {
				scriptList := make([]string, 0, len(scriptMap))
				for script, command := range scriptMap {
					scriptList = append(scriptList, fmt.Sprintf("%s: %s", script, command))
				}
				tui.RunTUI(projectPath, scriptList)
				break
			}
		},
	}
}
