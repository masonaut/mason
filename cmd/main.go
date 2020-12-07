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

var file string

func Execute() {
	MainCommand.Version = "0.1.0"
	MainCommand.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	MainCommand.PersistentFlags().StringVarP(&file, "file", "f", "", "Build config file to use")

	if err := MainCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
