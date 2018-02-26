// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InvoiceResponse invoice response
// swagger:model InvoiceResponse

type InvoiceResponse struct {

	// invoice
	Invoice *Invoice `json:"invoice,omitempty"`
}

/* polymorph InvoiceResponse invoice false */

// Validate validates this invoice response
func (m *InvoiceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvoice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceResponse) validateInvoice(formats strfmt.Registry) error {

	if swag.IsZero(m.Invoice) { // not required
		return nil
	}

	if m.Invoice != nil {

		if err := m.Invoice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("invoice")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceResponse) UnmarshalBinary(b []byte) error {
	var res InvoiceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
