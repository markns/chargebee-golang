// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CouponCopyRequest coupon copy request
// swagger:model CouponCopyRequest
type CouponCopyRequest struct {

	// for site merging
	ForSiteMerging bool `json:"for_site_merging,omitempty"`

	// from site
	FromSite string `json:"from_site,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// id at from site
	IDAtFromSite string `json:"id_at_from_site,omitempty"`
}

// Validate validates this coupon copy request
func (m *CouponCopyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CouponCopyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CouponCopyRequest) UnmarshalBinary(b []byte) error {
	var res CouponCopyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
