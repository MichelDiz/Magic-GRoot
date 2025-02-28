package cmd

import (
	"fmt"
	"mgr/internal/config"

	"github.com/spf13/cobra"
)

func SetLangCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set-lang [en|pt-br]",
		Short: "Define o idioma do CLI",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			lang := args[0]

			if lang != "en" && lang != "pt-br" {
				fmt.Println(config.Translate("suported_languages"))
				return
			}

			config.SetConfig("language", lang)

			fmt.Println(config.Translate("language_set"), lang)
		},
	}

}
