package config

import "path/filepath"

type Config interface {
	Get(string) interface{}
	GetString(string) string
	Set(string, interface{}) error
	UnSet(string) error
	Reset()
	Write() error
	Map() map[string]interface{}
}

type fileConfig struct {
	m      ConfigMap
	source string
}

func NewConfig(source string, schema []FieldSchema) (Config, error) {
	m, err := parseConfigFile(source)
	if err != nil {
		return nil, err
	}

	configMap := NewConfigMap(*m, schema)

	return &fileConfig{
		source: source,
		m:      *configMap,
	}, nil
}

func NewGlobalConfig() (Config, error) {
	path := filepath.Join(ConfigDir(), ConfigFileName)

	return NewConfig(path, ConfigGlobalSchema)
}

func NewBlankConfig() (Config, error) {
	err := createBlankConfigFile()
	if err != nil {
		return nil, err
	}

	source := filepath.Join(ConfigDir(), ConfigFileName)

	return NewConfig(source, ConfigGlobalSchema)
}

func (fc *fileConfig) Get(field string) interface{} {
	return fc.m.Get(field)
}

func (fc *fileConfig) GetString(field string) string {
	return fc.m.GetString(field)
}

func (fc *fileConfig) Set(field string, val interface{}) error {
	return fc.m.Set(field, val)
}

func (fc *fileConfig) UnSet(field string) error {
	return fc.m.Reset(field)
}

func (fc *fileConfig) Reset() {
	fc.m.root = fc.m.defaults
}

func (fc fileConfig) Map() map[string]interface{} {
	return fc.m.root
}

func (fc fileConfig) Write() error {
	data, err := fc.m.Byte()
	if err != nil {
		return err
	}

	return writeConfigFile(fc.source, data)
}
