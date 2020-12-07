package models

type Mason struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Version     string     `yaml:"version"`
	Workflows   []Workflow `yaml:"workflows"`
}

type Workflow struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Actions     []Action `yaml:"actions"`
}

type Action struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Commands    []string `yaml:"commands"`
}
