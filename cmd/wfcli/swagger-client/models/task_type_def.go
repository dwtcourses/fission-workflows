package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// TaskTypeDef task type def
// swagger:model TaskTypeDef
type TaskTypeDef struct {

	// resolved
	Resolved string `json:"resolved,omitempty"`

	// runtime
	Runtime string `json:"runtime,omitempty"`

	// src
	Src string `json:"src,omitempty"`
}

// Validate validates this task type def
func (m *TaskTypeDef) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
