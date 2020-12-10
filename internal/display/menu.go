package display

import (
	"fmt"
	"mason/internal/models"
)

func Menu(project models.Project) {
	fmt.Printf("%s - %s\n", project.Name, project.Version)
	fmt.Printf("%s\n\n", project.Description)

	fmt.Printf("%s\n", "workflows")
	fmt.Printf("%s\n", "=========")
	for _, workflow := range project.Workflows {
		fmt.Printf("%s\t%s\n", workflow.Name, workflow.Description)
	}
}
