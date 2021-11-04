package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

type Config struct {
	rootPath     string
	configFile   *ConfigFile
	resourceFile *ResourceFile
}

// It create new config struct and allocates config file on memory
func NewConfig() (*Config, error) {
	woDir, err := homedir.Expand("~/.wo/")
	if err != nil {
		return nil, err
	}

	c := &Config{
		rootPath: woDir,
	}
	err = c.load()

	return c, err
}

// load provides to alloc the memory and create config file if it doesn't exist
func (c *Config) load() error {
	err := c.loadConfig()
	if err != nil {
		return err
	}

	return c.loadResource()
}

func (c *Config) loadResource() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	rcPath := filepath.Join(dir, ResourceFileName)

	data, err := ioutil.ReadFile(rcPath)
	if err != nil && os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &c.resourceFile)

	return err
}

func (c *Config) loadConfig() error {
	configPath := filepath.Join(c.rootPath, "config.yml")

	data, err := ioutil.ReadFile(configPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(c.rootPath, 0755)
		if err != nil {
			return err
		}

		file, err := os.Create(configPath)
		if err != nil {
			return err
		}

		data, err = yaml.Marshal(defaultConfigFile)
		if err != nil {
			return err
		}

		_, err = file.Write(data)
		if err != nil {
			return err
		}

		file.Close()
	}

	err = yaml.Unmarshal(data, &c.configFile)

	return err
}

// Write function config in memory to config file
// Also this func run on cycle.
func (c *Config) Write() error {
	configPath := filepath.Join(c.rootPath, "config.yml")

	data, err := yaml.Marshal(c.configFile)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configPath, data, 0755)
}

func (c *Config) WriteResourceFile() error {
	if c.resourceFile == nil {
		return nil
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	rcPath := filepath.Join(dir, ResourceFileName)

	_, err = os.ReadFile(rcPath)
	if err != nil && os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	data, err := yaml.Marshal(c.resourceFile)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(rcPath, data, 0755)
}

func CreateDefaultRCFile(path string) error {
	rcPath := filepath.Join(path, ResourceFileName)

	file, err := os.Create(rcPath)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(defaultResourceFile)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	file.Close()

	return nil
}

// Reset function sets value of key as default value
func (c *Config) Reset() error {
	c.configFile = defaultConfigFile

	return c.Write()
}

// Config function gets loaded configuration data
func (c *Config) Config() *ConfigFile {
	if c.configFile == nil {
		return nil
	}

	return c.configFile
}

func (c *Config) Resource() *ResourceFile {
	return c.resourceFile
}
