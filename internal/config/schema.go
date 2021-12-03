package config

import (
	"path/filepath"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type FieldSchema struct {
	Key          string
	Description  string
	Rules        []validation.Rule
	DefaultValue interface{}
}

var (
	// Rules
	toggleRules = []validation.Rule{validation.Required, validation.In("enabled", "disabled")}

	// Utils Schemas
	editorsSchema = FieldSchema{
		Key:          "editors",
		Description:  "",
		DefaultValue: map[string]interface{}{},
	}
	spacesSchema = FieldSchema{
		Key:         "spaces",
		Description: "",
		DefaultValue: map[string]interface{}{
			"global": map[string]interface{}{
				"type":       "global",
				"id":         "global",
				"temp_dir":   filepath.Join(ConfigDir(), "./spaces/global-temp"),
				"root_dir":   filepath.Join(ConfigDir(), "./spaces/global"),
				"workspaces": map[string]interface{}{},
			},
		},
	}
	actionsSchema = FieldSchema{
		Key:          "actions",
		Description:  "",
		DefaultValue: map[string]interface{}{},
	}
	flowsSchema = FieldSchema{
		Key:          "flows",
		Description:  "",
		DefaultValue: map[string]interface{}{},
	}
	cyclesSchema = FieldSchema{
		Key:          "cycles",
		Description:  "",
		DefaultValue: map[string]interface{}{},
	}
	appsSchema = FieldSchema{
		Key:          "apps",
		Description:  "",
		DefaultValue: map[string]interface{}{},
	}
	defaultsSchema = FieldSchema{
		Key:         "defaults",
		Description: "defaults",
		Rules: []validation.Rule{validation.Map(
			validation.Key("readme", toggleRules...),
			validation.Key("code_of_conduct", toggleRules...),
			validation.Key("contributing", toggleRules...),
			validation.Key("license", toggleRules...),
			validation.Key("init_git", toggleRules...),
			validation.Key("gitignore", toggleRules...),
			validation.Key("author", is.PrintableASCII),
			validation.Key("editor", is.PrintableASCII),
		)},
		DefaultValue: map[string]interface{}{
			"readme":          "enabled",
			"code_of_conduct": "disabled",
			"contributing":    "disabled",
			"license":         "enabled",
			"init_git":        "enabled",
			"gitignore":       "enabled",
			"author":          "",
			"editor":          "",
		},
	}
	authSchema = FieldSchema{
		Key:         "auth",
		Description: "auth",
		Rules: []validation.Rule{
			validation.Map(
				validation.Key("github", validation.Map(
					validation.Key("token", validation.Required),
				)),
				validation.Key("gitlab", validation.Map(
					validation.Key("username", validation.Required),
					validation.Key("password", validation.Required),
				)),
			),
		},
		DefaultValue: map[string]map[string]string{
			"github": {
				"token": "",
			},
			"gitlab": {
				"username": "",
				"password": "",
			},
		},
	}

	// Global Config for schemas
	ConfigGlobalSchema = []FieldSchema{
		{
			Key:          "last_scan_editor",
			Description:  "the last editor scan time",
			Rules:        []validation.Rule{},
			DefaultValue: 0,
		},
		authSchema,
		defaultsSchema,
		editorsSchema,
		spacesSchema,
	}
	// Schema of spaceRC (.wospace)
	SpaceRCSchema = []FieldSchema{
		defaultsSchema,
		actionsSchema,
		flowsSchema,
		cyclesSchema,
	}
	// Schema of workspaceRC (.worc)
	WorkspaceRCSchema = []FieldSchema{
		{
			Key:          "name",
			Description:  "",
			DefaultValue: "",
		},
		{
			Key:          "description",
			Description:  "",
			DefaultValue: "",
		},
		{
			Key:          "version",
			Description:  "",
			DefaultValue: "v0.1",
		},
		actionsSchema,
		flowsSchema,
		cyclesSchema,
		appsSchema,
	}
)

func genSchemaMap(schema []FieldSchema) map[string]interface{} {
	m := make(map[string]interface{})

	for _, f := range schema {
		m[f.Key] = f.DefaultValue
	}

	return m
}
