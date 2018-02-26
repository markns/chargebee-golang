// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PortalSessionActivateRequest portal session activate request
// swagger:model PortalSessionActivateRequest

type PortalSessionActivateRequest struct {

	// token
	Token string `json:"token,omitempty"`
}

/* polymorph PortalSessionActivateRequest token false */

// Validate validates this portal session activate request
func (m *PortalSessionActivateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PortalSessionActivateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortalSessionActivateRequest) UnmarshalBinary(b []byte) error {
	var res PortalSessionActivateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
