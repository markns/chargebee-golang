// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreditNoteDeleteRequest credit note delete request
// swagger:model CreditNoteDeleteRequest

type CreditNoteDeleteRequest struct {

	// comment
	Comment string `json:"comment,omitempty"`
}

/* polymorph CreditNoteDeleteRequest comment false */

// Validate validates this credit note delete request
func (m *CreditNoteDeleteRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CreditNoteDeleteRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreditNoteDeleteRequest) UnmarshalBinary(b []byte) error {
	var res CreditNoteDeleteRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}