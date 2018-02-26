// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HostedPageResponse hosted page response
// swagger:model HostedPageResponse

type HostedPageResponse struct {

	// hostedpage
	Hostedpage *HostedPage `json:"hostedpage,omitempty"`
}

/* polymorph HostedPageResponse hostedpage false */

// Validate validates this hosted page response
func (m *HostedPageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostedpage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostedPageResponse) validateHostedpage(formats strfmt.Registry) error {

	if swag.IsZero(m.Hostedpage) { // not required
		return nil
	}

	if m.Hostedpage != nil {

		if err := m.Hostedpage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostedpage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostedPageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostedPageResponse) UnmarshalBinary(b []byte) error {
	var res HostedPageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}