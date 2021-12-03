package space

type Options struct {
	Readme        string `json:"readme"`
	CodeOfConduct string `json:"code_of_conduct"`
	Contributing  string `json:"contributing"`
	License       string `json:"license"`
	Git           string `json:"git"`
	Gitignore     string `json:"gitignore"`
	Editor        string `json:"editor"`
}

func (o *Options) Map() map[string]interface{} {
	m := make(map[string]interface{})

	m["readme"] = o.Readme
	m["code_of_conduct"] = o.CodeOfConduct
	m["contributing"] = o.Contributing
	m["license"] = o.License
	m["git"] = o.Git
	m["gitignore"] = o.Gitignore
	m["editor"] = o.Editor

	return m
}

type Space struct {
	ID          string               `json:"id"`
	Description string               `json:"description"`
	TempDir     string               `json:"temp_dir"`
	RootDir     string               `json:"root_dir"`
	Defaults    Options              `json:"defaults"`
	CreatedAt   int64                `json:"created_at"`
	Workspaces  map[string]Workspace `json:"workspaces"`
}

func (s *Space) Map() map[string]interface{} {
	m := s.MapWithoutWorkspaces()

	m["defaults"] = s.Defaults.Map()

	if s.Workspaces != nil || len(s.Workspaces) > 0 {
		mws := make(map[string]interface{})
		for id, ws := range s.Workspaces {
			mws[id] = ws.Map()
		}

		m["workspaces"] = mws
	}

	return m
}

func (s *Space) MapForConfig() map[string]interface{} {
	m := s.MapWithoutWorkspaces()

	mws := make(map[string]interface{})
	for id, ws := range s.Workspaces {
		mws[id] = ws.MapForConfig()
	}

	m["workspaces"] = mws

	return m
}

func (s *Space) MapWithoutWorkspaces() map[string]interface{} {
	m := make(map[string]interface{})

	m["id"] = s.ID
	m["root_dir"] = s.RootDir
	m["temp_dir"] = s.TempDir
	m["description"] = s.Description
	m["created_at"] = s.CreatedAt

	return m
}

type Workspace struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Path        string   `json:"path"`
	Description string   `json:"description"`
	Run         string   `json:"run"`
	Actions     []Action `json:"actions"`
	Flows       []Flows  `json:"flows"`
	CreatedAt   int64    `json:"created_at"`
}

func (ws *Workspace) Map() map[string]interface{} {
	m := ws.MapForRC()

	m["id"] = ws.ID
	m["path"] = ws.Path
	m["created_at"] = ws.CreatedAt

	return m
}

func (ws *Workspace) MapForRC() map[string]interface{} {
	m := make(map[string]interface{})

	m["name"] = ws.Name
	m["description"] = ws.Description
	m["run"] = ws.Run

	if ws.Flows != nil {
		flows := []map[string]interface{}{}
		for i, f := range ws.Flows {
			flows[i] = f.Map()
		}
	}

	if ws.Actions != nil {
		actions := []map[string]interface{}{}
		for i, a := range ws.Actions {
			actions[i] = a.Map()
		}
	}

	return m
}

func (ws *Workspace) MapForConfig() map[string]interface{} {
	m := make(map[string]interface{})

	m["id"] = ws.ID
	m["name"] = ws.Name
	m["path"] = ws.Path
	m["created_at"] = ws.CreatedAt

	return m
}

type Flows struct {
	Name       string   `json:"name"`
	Env        string   `json:"env"`
	WorkingDir string   `json:"working_dir"`
	Steps      []Action `json:"steps"`
}

func (f *Flows) Map() map[string]interface{} {
	m := make(map[string]interface{})

	m["name"] = f.Name
	m["env"] = f.Env
	m["working_dir"] = f.WorkingDir

	steps := []map[string]interface{}{}

	for i, s := range f.Steps {
		steps[i] = s.Map()
	}

	m["steps"] = steps

	return m
}

type Action struct {
	Name       string   `json:"name"`
	Env        []string `json:"env"`
	Args       []string `json:"args"`
	Run        string   `json:"run"`
	Workingdir string   `json:"working_dir"`
}

func (a *Action) Map() map[string]interface{} {
	m := make(map[string]interface{})

	m["name"] = a.Name
	m["env"] = a.Env
	m["args"] = a.Args
	m["run"] = a.Run
	m["working_dir"] = a.Workingdir

	return m
}
