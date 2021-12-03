package config

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gopkg.in/yaml.v2"
)

var (
	ErrConfigMapFieldNotFound    = "field value not found"
	ErrConfigMapInvalidFieldAddr = "invalid field addr"
	ErrConfigValidation          = "%s validation err: %s"
)

type ConfigMap struct {
	root     map[string]interface{}
	defaults map[string]interface{}
	schema   []FieldSchema
}

func NewConfigMap(root map[string]interface{}, schema []FieldSchema) *ConfigMap {
	defaults := genSchemaMap(schema)

	cm := &ConfigMap{
		root:     root,
		defaults: defaults,
		schema:   schema,
	}
	return cm
}

func (cm *ConfigMap) Get(field string) interface{} {
	ks := strings.Split(field, ".")
	lastKey := ks[len(ks)-1]

	nestedMap, err := cm.searchMap(cm.root, ks[:len(ks)-1]...)
	if err != nil {
		m, err := cm.searchMap(cm.defaults, ks[:len(ks)-1]...)
		if err != nil {
			return nil
		}

		return m[lastKey]
	}

	return nestedMap[lastKey]
}

func (cm *ConfigMap) GetString(field string) string {
	val := cm.Get(field)
	if val == nil || reflect.TypeOf(val).Kind() != reflect.String {
		return ""
	}

	return fmt.Sprintf("%v", val)
}

func (cm *ConfigMap) GetSchema(key string) (*FieldSchema, error) {
	for _, s := range cm.schema {
		if s.Key == key {
			return &s, nil
		}
	}

	return nil, errors.New(ErrConfigMapFieldNotFound)
}

func (cm *ConfigMap) Set(field string, val interface{}) error {
	ks := strings.Split(field, ".")
	lastKey := ks[len(ks)-1]

	err := cm.validate(val, field)
	if err != nil {
		return err
	}

	nestedMap := cm.deepSearchMap(cm.root, ks[:len(ks)-1]...)

	nestedMap[lastKey] = val
	return nil
}

func (cm *ConfigMap) Reset(field string) error {
	ks := strings.Split(field, ".")
	lastKey := ks[len(ks)-1]

	nestedMap, err := cm.searchMap(cm.root, ks[:len(ks)-1]...)
	if err != nil {
		return err
	}

	return cm.Set(field, nestedMap[lastKey])
}

func (cm *ConfigMap) Byte() ([]byte, error) {
	return yaml.Marshal(cm.root)
}

func (cm ConfigMap) Map() map[string]interface{} {
	return cm.root
}

func (cm *ConfigMap) validate(val interface{}, field string) error {
	ks := strings.Split(field, ".")

	if len(ks) > 2 {
		return nil
	}

	schema, err := cm.GetSchema(ks[0])
	if err != nil {
		return fmt.Errorf(ErrConfigValidation, field, err)
	}

	err = validation.Validate(val, schema.Rules...)
	if err != nil {
		return fmt.Errorf(ErrConfigValidation, schema.Key, err)
	}

	return nil
}

func (cm *ConfigMap) searchMap(m map[string]interface{}, keys ...string) (map[string]interface{}, error) {
	if len(keys) == 0 {
		return m, nil
	}

	for _, k := range keys {
		m2, ok := m[k]
		if !ok {
			return nil, errors.New(ErrConfigMapInvalidFieldAddr)
		}

		m3, ok := m2.(map[string]interface{})
		if !ok {
			switch m2.(type) {
			case map[string]interface{}, map[interface{}]interface{}:
				m3 = make(map[string]interface{})
				m[k] = m3
			default:
				return nil, errors.New(ErrConfigMapInvalidFieldAddr)
			}
		}

		m = m3
	}

	return m, nil
}

func (cm *ConfigMap) deepSearchMap(m map[string]interface{}, keys ...string) map[string]interface{} {
	if len(keys) == 0 {
		return m
	}

	for _, k := range keys {
		m2, ok := m[k]
		if !ok {
			m3 := make(map[string]interface{})
			m[k] = m3
			m = m3
			continue
		}

		m3, ok := m2.(map[string]interface{})
		if !ok {
			m3 = make(map[string]interface{})
			m[k] = m3
		}

		m = m3
	}

	return m
}
