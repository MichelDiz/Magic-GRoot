package main

import (
	"context"
	"mgr/cmd"
	"mgr/internal/config"
	"os"
	"os/signal"
	"syscall"

	"mgr/internal/logger"

	"github.com/spf13/cobra"
)

func execute(ctx context.Context) error {

	var RootCmd = &cobra.Command{
		Use:   "mgr",
		Short: "Magic GRoot CLI",
		Long:  config.Translate("root"),
	}

	RootCmd.AddCommand(cmd.ScanCmd())
	RootCmd.AddCommand(cmd.SetRootCmd())
	RootCmd.AddCommand(cmd.SetLangCmd())
	RootCmd.AddCommand(cmd.ListCmd())
	RootCmd.AddCommand(cmd.AliasCmd())
	RootCmd.AddCommand(cmd.GoCmd())
	RootCmd.AddCommand(cmd.SetRunnerCmd())

	return RootCmd.ExecuteContext(ctx)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		logger.Info.Println("Shutting down gracefully...")
		cancel()
	}()

	config.InitDB()
	config.InitI18n()
	lang := config.GetLanguage()
	config.SetConfig("language", lang)

	if err := execute(ctx); err != nil {
		logger.Error.Printf("Application error: %v\n", err)
		os.Exit(1)
	}
}
