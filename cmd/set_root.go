package cmd

import (
	"fmt"
	"mgr/internal/config"

	"github.com/spf13/cobra"
)

func SetRootCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "set-root [path]",
		Short: config.Translate("set_root_short"),
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) > 0 {
				config.SetConfig("root_path", args[0])
				fmt.Println(config.Translate("root_path_set"), args[0])
			} else {
				fmt.Println(config.Translate("set_root_usage"))
			}
			if len(args) < 1 {
				fmt.Println("Uso: mgr set-root [path]")
				return
			}

		},
	}

}
