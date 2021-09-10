package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

type Config struct {
	rootPath   string
	configFile *ConfigFile
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
	configPath := filepath.Join(c.rootPath, "config.yml")

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
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
	if err != nil {
		return err
	}

	return nil
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
