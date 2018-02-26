// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InvoiceRemovePaymentRequest invoice remove payment request
// swagger:model InvoiceRemovePaymentRequest

type InvoiceRemovePaymentRequest struct {

	// transaction id
	TransactionID string `json:"transaction[id],omitempty"`
}

/* polymorph InvoiceRemovePaymentRequest transaction[id] false */

// Validate validates this invoice remove payment request
func (m *InvoiceRemovePaymentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceRemovePaymentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceRemovePaymentRequest) UnmarshalBinary(b []byte) error {
	var res InvoiceRemovePaymentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
