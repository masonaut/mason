package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var MainCommand = &cobra.Command{
	Use:   "mason",
	Short: "build system",
}

func Execute() {
	MainCommand.Version = "0.1.0"
	MainCommand.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	if err := MainCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
