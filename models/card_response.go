// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CardResponse card response
// swagger:model CardResponse

type CardResponse struct {

	// card
	Card *Card `json:"card,omitempty"`
}

/* polymorph CardResponse card false */

// Validate validates this card response
func (m *CardResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCard(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CardResponse) validateCard(formats strfmt.Registry) error {

	if swag.IsZero(m.Card) { // not required
		return nil
	}

	if m.Card != nil {

		if err := m.Card.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CardResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CardResponse) UnmarshalBinary(b []byte) error {
	var res CardResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
