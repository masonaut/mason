package cmd

import (
	"fmt"
	"mason/internal/files"
	"mason/internal/models"
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

		var actions []models.Action

		masonfile := parse.Yaml(files.ReadFileToByte(file))
		for _, val := range masonfile.Workflows {
			if val.Name == args[0] {
				actions = val.Actions
				fmt.Println(actions)
				return
			}
		}

		fmt.Println("workflow not found")
	},
}

func init() {
	MainCommand.AddCommand(WorkflowCommand)
}
