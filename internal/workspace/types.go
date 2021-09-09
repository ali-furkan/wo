package workspace

import "time"

type WorkTypes uint

const (
	Template WorkTypes = iota
	Created
	Init
	Temporary
)

type Workspace struct {
	TempWorkDir    string    `yaml:"temp_ws_dir"`
	WorkDir        string    `yaml:"ws_dir"`
	DefaultEditor  string    `yaml:"default_editor"`
	LastScanEditor time.Time `yaml:"last_scan_editor"`
	DefaultGit     bool      `yaml:"default_git"`
	DefaultReadme  bool      `yaml:"default_readme"`
	DefaultLicense string    `yaml:"default_license"`
	Works          []Work    `yaml:"works"`
}

type Work struct {
	ID          string    `yaml:"id"`
	Type        WorkTypes `yaml:"type"`
	Name        string    `yaml:"name"`
	Path        string    `yaml:"path"`
	Description string    `yaml:"description"`
	Template    string    `yaml:"template"`
	InitReadme  bool      `yaml:"init_readme"`
	InitGit     bool      `yaml:"init_git"`
	License     string    `yaml:"license"`
	RunScript   string    `yaml:"run_script"`
	CreatedAt   time.Time `yaml:"created_at"`
	UpdatedAt   time.Time `yaml:"updated_at"`
}
