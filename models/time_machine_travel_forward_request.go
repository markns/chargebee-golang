// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TimeMachineTravelForwardRequest time machine travel forward request
// swagger:model TimeMachineTravelForwardRequest
type TimeMachineTravelForwardRequest struct {

	// destination time
	DestinationTime int64 `json:"destination_time,omitempty"`
}

// Validate validates this time machine travel forward request
func (m *TimeMachineTravelForwardRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TimeMachineTravelForwardRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeMachineTravelForwardRequest) UnmarshalBinary(b []byte) error {
	var res TimeMachineTravelForwardRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
