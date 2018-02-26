// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InvoiceChargeAddonRequest invoice charge addon request
// swagger:model InvoiceChargeAddonRequest

type InvoiceChargeAddonRequest struct {

	// addon id
	AddonID string `json:"addon_id,omitempty"`

	// addon quantity
	AddonQuantity int32 `json:"addon_quantity,omitempty"`

	// addon unit price
	AddonUnitPrice int32 `json:"addon_unit_price,omitempty"`

	// coupon
	Coupon string `json:"coupon,omitempty"`

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// payment source id
	PaymentSourceID string `json:"payment_source_id,omitempty"`

	// po number
	PoNumber string `json:"po_number,omitempty"`

	// subscription id
	SubscriptionID string `json:"subscription_id,omitempty"`
}

/* polymorph InvoiceChargeAddonRequest addon_id false */

/* polymorph InvoiceChargeAddonRequest addon_quantity false */

/* polymorph InvoiceChargeAddonRequest addon_unit_price false */

/* polymorph InvoiceChargeAddonRequest coupon false */

/* polymorph InvoiceChargeAddonRequest customer_id false */

/* polymorph InvoiceChargeAddonRequest payment_source_id false */

/* polymorph InvoiceChargeAddonRequest po_number false */

/* polymorph InvoiceChargeAddonRequest subscription_id false */

// Validate validates this invoice charge addon request
func (m *InvoiceChargeAddonRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceChargeAddonRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceChargeAddonRequest) UnmarshalBinary(b []byte) error {
	var res InvoiceChargeAddonRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
