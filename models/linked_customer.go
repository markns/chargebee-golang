// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LinkedCustomer linked customer
// swagger:model LinkedCustomer

type LinkedCustomer struct {

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// has active subscription
	HasActiveSubscription bool `json:"has_active_subscription,omitempty"`

	// has billing address
	HasBillingAddress bool `json:"has_billing_address,omitempty"`

	// has payment method
	HasPaymentMethod bool `json:"has_payment_method,omitempty"`
}

/* polymorph LinkedCustomer customer_id false */

/* polymorph LinkedCustomer email false */

/* polymorph LinkedCustomer has_active_subscription false */

/* polymorph LinkedCustomer has_billing_address false */

/* polymorph LinkedCustomer has_payment_method false */

// Validate validates this linked customer
func (m *LinkedCustomer) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *LinkedCustomer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkedCustomer) UnmarshalBinary(b []byte) error {
	var res LinkedCustomer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}