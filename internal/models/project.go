package models

type Project struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Version     string     `yaml:"version"`
	Workflows   []Workflow `yaml:"workflows"`
}
