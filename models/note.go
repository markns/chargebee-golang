// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Note note
// swagger:model Note

type Note struct {

	// entity id
	EntityID string `json:"entity_id,omitempty"`

	// entity type
	EntityType string `json:"entity_type,omitempty"`

	// note
	Note string `json:"note,omitempty"`
}

/* polymorph Note entity_id false */

/* polymorph Note entity_type false */

/* polymorph Note note false */

// Validate validates this note
func (m *Note) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var noteTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["plan_setup","plan","addon","adhoc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		noteTypeEntityTypePropEnum = append(noteTypeEntityTypePropEnum, v)
	}
}

const (
	// NoteEntityTypePlanSetup captures enum value "plan_setup"
	NoteEntityTypePlanSetup string = "plan_setup"
	// NoteEntityTypePlan captures enum value "plan"
	NoteEntityTypePlan string = "plan"
	// NoteEntityTypeAddon captures enum value "addon"
	NoteEntityTypeAddon string = "addon"
	// NoteEntityTypeAdhoc captures enum value "adhoc"
	NoteEntityTypeAdhoc string = "adhoc"
)

// prop value enum
func (m *Note) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, noteTypeEntityTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Note) validateEntityType(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityTypeEnum("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Note) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Note) UnmarshalBinary(b []byte) error {
	var res Note
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}