// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CouponSetCreateRequest coupon set create request
// swagger:model CouponSetCreateRequest
type CouponSetCreateRequest struct {

	// coupon id
	CouponID string `json:"coupon_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// meta data
	MetaData interface{} `json:"meta_data,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this coupon set create request
func (m *CouponSetCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CouponSetCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CouponSetCreateRequest) UnmarshalBinary(b []byte) error {
	var res CouponSetCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
