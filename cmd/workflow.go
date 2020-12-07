package cmd

import (
	"fmt"
	"mason/internal/files"
	"mason/internal/parse"
	"os"

	"github.com/spf13/cobra"
)

var WorkflowCommand = &cobra.Command{
	Use:   "workflow",
	Short: "Execute workflow",
	Long:  "Execute workflow",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		masonfile := parse.Yaml(files.ReadFileToByte(file))
		for _, val := range masonfile.Workflows {
			if val.Name == args[0] {
				fmt.Println(val.Actions)
				return
			}
		}

		fmt.Println("workflow not found")
	},
}

var WorkflowListCommand = &cobra.Command{
	Use:   "list",
	Short: "List workflows",
	Long:  "List workflows",
	Run: func(cmd *cobra.Command, args []string) {
		masonfile := parse.Yaml(files.ReadFileToByte(file))
		for _, v := range masonfile.Workflows {
			fmt.Println(v.Name)
		}
	},
}

func init() {
	WorkflowCommand.AddCommand(WorkflowListCommand)
	MainCommand.AddCommand(WorkflowCommand)
}
