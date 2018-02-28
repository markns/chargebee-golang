// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UnbilledChargeResponse unbilled charge response
// swagger:model UnbilledChargeResponse
type UnbilledChargeResponse struct {

	// unbilledcharge
	Unbilledcharge *UnbilledCharge `json:"unbilledcharge,omitempty"`
}

// Validate validates this unbilled charge response
func (m *UnbilledChargeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnbilledcharge(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnbilledChargeResponse) validateUnbilledcharge(formats strfmt.Registry) error {

	if swag.IsZero(m.Unbilledcharge) { // not required
		return nil
	}

	if m.Unbilledcharge != nil {

		if err := m.Unbilledcharge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unbilledcharge")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UnbilledChargeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnbilledChargeResponse) UnmarshalBinary(b []byte) error {
	var res UnbilledChargeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
