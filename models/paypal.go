// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Paypal paypal
// swagger:model Paypal

type Paypal struct {

	// agreement id
	AgreementID string `json:"agreement_id,omitempty"`

	// email
	Email string `json:"email,omitempty"`
}

/* polymorph Paypal agreement_id false */

/* polymorph Paypal email false */

// Validate validates this paypal
func (m *Paypal) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Paypal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Paypal) UnmarshalBinary(b []byte) error {
	var res Paypal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
