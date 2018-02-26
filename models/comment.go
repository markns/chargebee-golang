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

// Comment comment
// swagger:model Comment

type Comment struct {

	// added by
	AddedBy string `json:"added_by,omitempty"`

	// created at
	CreatedAt int64 `json:"created_at,omitempty"`

	// entity id
	EntityID string `json:"entity_id,omitempty"`

	// entity type
	EntityType string `json:"entity_type,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// notes
	Notes string `json:"notes,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

/* polymorph Comment added_by false */

/* polymorph Comment created_at false */

/* polymorph Comment entity_id false */

/* polymorph Comment entity_type false */

/* polymorph Comment id false */

/* polymorph Comment notes false */

/* polymorph Comment type false */

// Validate validates this comment
func (m *Comment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var commentTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["plan_setup","plan","addon","adhoc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commentTypeEntityTypePropEnum = append(commentTypeEntityTypePropEnum, v)
	}
}

const (
	// CommentEntityTypePlanSetup captures enum value "plan_setup"
	CommentEntityTypePlanSetup string = "plan_setup"
	// CommentEntityTypePlan captures enum value "plan"
	CommentEntityTypePlan string = "plan"
	// CommentEntityTypeAddon captures enum value "addon"
	CommentEntityTypeAddon string = "addon"
	// CommentEntityTypeAdhoc captures enum value "adhoc"
	CommentEntityTypeAdhoc string = "adhoc"
)

// prop value enum
func (m *Comment) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, commentTypeEntityTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Comment) validateEntityType(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityTypeEnum("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

var commentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["authorization","payment","refund","payment_reversal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commentTypeTypePropEnum = append(commentTypeTypePropEnum, v)
	}
}

const (
	// CommentTypeAuthorization captures enum value "authorization"
	CommentTypeAuthorization string = "authorization"
	// CommentTypePayment captures enum value "payment"
	CommentTypePayment string = "payment"
	// CommentTypeRefund captures enum value "refund"
	CommentTypeRefund string = "refund"
	// CommentTypePaymentReversal captures enum value "payment_reversal"
	CommentTypePaymentReversal string = "payment_reversal"
)

// prop value enum
func (m *Comment) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, commentTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Comment) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Comment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Comment) UnmarshalBinary(b []byte) error {
	var res Comment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}