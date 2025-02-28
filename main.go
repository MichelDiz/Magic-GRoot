package main

import (
	"mgr/cmd"
	"mgr/internal/config"

	"github.com/spf13/cobra"
)

func execute() {

	var RootCmd = &cobra.Command{
		Use:   "mgr",
		Short: "Magic GRoot CLI",
		Long:  config.Translate("root"),
	}

	RootCmd.AddCommand(cmd.ScanCmd())
	RootCmd.AddCommand(cmd.SetRootCmd())
	RootCmd.AddCommand(cmd.SetLangCmd())
	cobra.CheckErr(RootCmd.Execute())
}

func main() {
	config.InitDB()
	config.InitI18n()
	lang := config.GetLanguage()
	config.SetConfig("language", lang)
	execute()
}
