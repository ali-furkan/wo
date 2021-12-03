package editor

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Editor struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Exec string `json:"exec"`
}

var (
	EditorNameRule = []validation.Rule{validation.Required, validation.Length(3, 32), is.PrintableASCII}
	EditorIDRule   = []validation.Rule{validation.Required, validation.Length(3, 32), is.Alphanumeric}
	EditorExecRule = []validation.Rule{validation.Required}
)

func (e *Editor) Validate() error {
	return validation.ValidateStruct(e,
		validation.Field(&e.Name, EditorNameRule...),
		validation.Field(&e.ID, EditorIDRule...),
		validation.Field(&e.Exec, EditorExecRule...),
	)
}
