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

// CommentCreateRequest comment create request
// swagger:model CommentCreateRequest
type CommentCreateRequest struct {

	// added by
	AddedBy string `json:"added_by,omitempty"`

	// entity id
	EntityID string `json:"entity_id,omitempty"`

	// entity type
	EntityType string `json:"entity_type,omitempty"`

	// notes
	Notes string `json:"notes,omitempty"`
}

// Validate validates this comment create request
func (m *CommentCreateRequest) Validate(formats strfmt.Registry) error {
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

var commentCreateRequestTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["customer","subscription","invoice","credit_note","transaction","plan","addon","coupon"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commentCreateRequestTypeEntityTypePropEnum = append(commentCreateRequestTypeEntityTypePropEnum, v)
	}
}

const (

	// CommentCreateRequestEntityTypeCustomer captures enum value "customer"
	CommentCreateRequestEntityTypeCustomer string = "customer"

	// CommentCreateRequestEntityTypeSubscription captures enum value "subscription"
	CommentCreateRequestEntityTypeSubscription string = "subscription"

	// CommentCreateRequestEntityTypeInvoice captures enum value "invoice"
	CommentCreateRequestEntityTypeInvoice string = "invoice"

	// CommentCreateRequestEntityTypeCreditNote captures enum value "credit_note"
	CommentCreateRequestEntityTypeCreditNote string = "credit_note"

	// CommentCreateRequestEntityTypeTransaction captures enum value "transaction"
	CommentCreateRequestEntityTypeTransaction string = "transaction"

	// CommentCreateRequestEntityTypePlan captures enum value "plan"
	CommentCreateRequestEntityTypePlan string = "plan"

	// CommentCreateRequestEntityTypeAddon captures enum value "addon"
	CommentCreateRequestEntityTypeAddon string = "addon"

	// CommentCreateRequestEntityTypeCoupon captures enum value "coupon"
	CommentCreateRequestEntityTypeCoupon string = "coupon"
)

// prop value enum
func (m *CommentCreateRequest) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, commentCreateRequestTypeEntityTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CommentCreateRequest) validateEntityType(formats strfmt.Registry) error {

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
func (m *CommentCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommentCreateRequest) UnmarshalBinary(b []byte) error {
	var res CommentCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
