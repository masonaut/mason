package models

type Workflow struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Actions     []Action `yaml:"actions"`
}
