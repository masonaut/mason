package cmd

import (
	"fmt"
	"mason/internal/files"
	"mason/internal/parse"
	"os"

	"github.com/spf13/cobra"
)

var file string

var MainCommand = &cobra.Command{
	Use:   "mason",
	Short: "build system",
	Run: func(cmd *cobra.Command, args []string) {
		if file == "" {
			cmd.Help()
			os.Exit(0)
		}

		masonfile := parse.Yaml(files.ReadFileToByte(file))
		for _, v := range masonfile.Workflows {
			fmt.Println(v.Name)
		}
	},
}

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
