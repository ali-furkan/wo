package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const (
	// Config Files Names
	ConfigFileName = "config.yml"
	StateFileName  = "state.yml"

	// RC File Names
	SpaceRCFileName   = ".wospacerc"
	WorkspaceFileName = ".worc"

	// UserConfigDir Env
	WO_CONFIG_DIR   = "WO_CONFIG_DIR"
	XDG_CONFIG_HOME = "XDG_CONFIG_HOME"
	APP_DATA        = "AppData"

	// WoConfigDir
	WoXDGConfigDir = "wo"
	WoAppDataDir   = "Wo CLI"
	WoHomeDir      = ".wo"
)

// ConfigDir is Wo configuration dir path
func ConfigDir() string {
	var path string

	if woConfigDir := os.Getenv(WO_CONFIG_DIR); woConfigDir != "" {
		path = woConfigDir
	} else if xdgConfigHome := os.Getenv(XDG_CONFIG_HOME); xdgConfigHome != "" {
		path = filepath.Join(xdgConfigHome, WoXDGConfigDir)
	} else if appData := os.Getenv(APP_DATA); appData != "" {
		path = filepath.Join(appData, WoAppDataDir)
	} else {
		userHomeDir, _ := os.UserHomeDir()
		path = filepath.Join(userHomeDir, WoHomeDir)
	}

	if !dirExists(path) {
		_ = createConfigDir(path)
	}

	return path
}

func dirExists(path string) bool {
	f, err := os.Stat(path)

	return err == nil && f.IsDir()
}

func createConfigDir(path string) error {
	if dirExists(path) {
		return os.ErrExist
	}

	return os.MkdirAll(path, 0755)
}

func createBlankConfigFile() error {
	configDir := ConfigDir()

	configPath := filepath.Join(configDir, ConfigFileName)

	data, err := yaml.Marshal(genSchemaMap(ConfigGlobalSchema))
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configPath, data, 0755)
}

func readConfigFile(filename string) ([]byte, error) {
	path := filepath.Clean(filename)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func writeConfigFile(filename string, data []byte) error {
	path := filepath.Clean(filename)

	return ioutil.WriteFile(path, data, 0600)
}

func parseConfigFile(filename string) (*map[string]interface{}, error) {
	data, err := readConfigFile(filename)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})

	err = yaml.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
