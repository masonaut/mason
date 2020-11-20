package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var file string

var WorkflowCommand = &cobra.Command{
	Use:     "workflow",
	Short:   "Execute workflow",
	Long:    "Execute workflow",
	Aliases: []string{"w"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(file)
	},
}

func init() {
	MainCommand.AddCommand(WorkflowCommand)
	WorkflowCommand.Flags().StringVarP(&file, "file", "f", "", "Mason file to use")

	WorkflowCommand.MarkFlagRequired("file")
}
