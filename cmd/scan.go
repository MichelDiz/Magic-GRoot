package cmd

import (
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

			scripts := scanner.ScanForScripts(rootPath)
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

			for projectPath, scriptList := range scripts {
				tui.RunTUI(projectPath, scriptList)
				break
			}
		},
	}
}
