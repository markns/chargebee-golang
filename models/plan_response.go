// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PlanResponse plan response
// swagger:model PlanResponse

type PlanResponse struct {

	// plan
	Plan *Plan `json:"plan,omitempty"`
}

/* polymorph PlanResponse plan false */

// Validate validates this plan response
func (m *PlanResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlan(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanResponse) validatePlan(formats strfmt.Registry) error {

	if swag.IsZero(m.Plan) { // not required
		return nil
	}

	if m.Plan != nil {

		if err := m.Plan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanResponse) UnmarshalBinary(b []byte) error {
	var res PlanResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}