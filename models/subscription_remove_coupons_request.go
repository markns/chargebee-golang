// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SubscriptionRemoveCouponsRequest subscription remove coupons request
// swagger:model SubscriptionRemoveCouponsRequest
type SubscriptionRemoveCouponsRequest struct {

	// coupon ids
	CouponIds string `json:"coupon_ids,omitempty"`
}

// Validate validates this subscription remove coupons request
func (m *SubscriptionRemoveCouponsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionRemoveCouponsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionRemoveCouponsRequest) UnmarshalBinary(b []byte) error {
	var res SubscriptionRemoveCouponsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
