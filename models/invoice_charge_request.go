// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InvoiceChargeRequest invoice charge request
// swagger:model InvoiceChargeRequest

type InvoiceChargeRequest struct {

	// amount
	Amount int32 `json:"amount,omitempty"`

	// coupon
	Coupon string `json:"coupon,omitempty"`

	// currency code
	CurrencyCode string `json:"currency_code,omitempty"`

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// payment source id
	PaymentSourceID string `json:"payment_source_id,omitempty"`

	// po number
	PoNumber string `json:"po_number,omitempty"`

	// subscription id
	SubscriptionID string `json:"subscription_id,omitempty"`
}

/* polymorph InvoiceChargeRequest amount false */

/* polymorph InvoiceChargeRequest coupon false */

/* polymorph InvoiceChargeRequest currency_code false */

/* polymorph InvoiceChargeRequest customer_id false */

/* polymorph InvoiceChargeRequest description false */

/* polymorph InvoiceChargeRequest payment_source_id false */

/* polymorph InvoiceChargeRequest po_number false */

/* polymorph InvoiceChargeRequest subscription_id false */

// Validate validates this invoice charge request
func (m *InvoiceChargeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceChargeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceChargeRequest) UnmarshalBinary(b []byte) error {
	var res InvoiceChargeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
