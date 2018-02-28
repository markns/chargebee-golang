// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CouponSetAddCouponCodesRequest coupon set add coupon codes request
// swagger:model CouponSetAddCouponCodesRequest
type CouponSetAddCouponCodesRequest struct {

	// code
	Code string `json:"code,omitempty"`
}

// Validate validates this coupon set add coupon codes request
func (m *CouponSetAddCouponCodesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CouponSetAddCouponCodesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CouponSetAddCouponCodesRequest) UnmarshalBinary(b []byte) error {
	var res CouponSetAddCouponCodesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
