// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EstimateResponse estimate response
// swagger:model EstimateResponse
type EstimateResponse struct {

	// estimate
	Estimate *Estimate `json:"estimate,omitempty"`
}

// Validate validates this estimate response
func (m *EstimateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEstimate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EstimateResponse) validateEstimate(formats strfmt.Registry) error {

	if swag.IsZero(m.Estimate) { // not required
		return nil
	}

	if m.Estimate != nil {

		if err := m.Estimate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("estimate")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EstimateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EstimateResponse) UnmarshalBinary(b []byte) error {
	var res EstimateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
