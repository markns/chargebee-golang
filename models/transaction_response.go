// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TransactionResponse transaction response
// swagger:model TransactionResponse
type TransactionResponse struct {

	// transaction
	Transaction *Transaction `json:"transaction,omitempty"`
}

// Validate validates this transaction response
func (m *TransactionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionResponse) validateTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.Transaction) { // not required
		return nil
	}

	if m.Transaction != nil {

		if err := m.Transaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transaction")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionResponse) UnmarshalBinary(b []byte) error {
	var res TransactionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
