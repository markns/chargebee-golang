// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ListPlanOKBody list plan o k body
// swagger:model listPlanOKBody
type ListPlanOKBody struct {

	// list
	List []*PlanResponse `json:"list"`

	// next offset
	NextOffset string `json:"next_offset,omitempty"`
}

// Validate validates this list plan o k body
func (m *ListPlanOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListPlanOKBody) validateList(formats strfmt.Registry) error {

	if swag.IsZero(m.List) { // not required
		return nil
	}

	for i := 0; i < len(m.List); i++ {

		if swag.IsZero(m.List[i]) { // not required
			continue
		}

		if m.List[i] != nil {

			if err := m.List[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListPlanOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListPlanOKBody) UnmarshalBinary(b []byte) error {
	var res ListPlanOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
