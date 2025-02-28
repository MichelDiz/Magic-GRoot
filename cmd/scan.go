package cmd

import (
	"fmt"
	"mgr/internal/config"
	"mgr/internal/scanner"
	"strings"

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
			for project, scriptList := range scripts {
				fmt.Println(config.Translate("project"), project)
				fmt.Println(config.Translate("available_scripts"), strings.Join(scriptList, ", "))
			}
		},
	}
}
