package space

import "time"

type Options struct {
	Readme        string `json:"readme"`
	CodeOfConduct string `json:"code_of_conduct"`
	Contributing  string `json:"contributing"`
	License       string `json:"license"`
	Git           string `json:"git"`
	Gitignore     string `json:"gitignore"`
	Editor        string `json:"editor"`
}

type Space struct {
	TempDir  string               `json:"temp_dir"`
	RootDir  string               `json:"root_dir"`
	Defaults Options              `json:"defaults"`
	Works    map[string]Workspace `json:"works"`
}

type Workspace struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Description string    `json:"description"`
	Run         string    `json:"run"`
	Actions     []Action  `json:"actions"`
	Flows       []Flows   `json:"flows"`
	CreatedAt   time.Time `json:"created_at"`
}

type Action struct {
	Name       string   `json:"name"`
	Env        []string `json:"env"`
	Args       []string `json:"args"`
	Run        string   `json:"run"`
	Workingdir string   `json:"working_dir"`
}

type Flows struct {
	Name       string   `json:"name"`
	Env        string   `json:"env"`
	WorkingDir string   `json:"working_dir"`
	Steps      []Action `json:"steps"`
}
