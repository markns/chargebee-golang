// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UnbilledChargeInvoiceNowEstimateRequest unbilled charge invoice now estimate request
// swagger:model UnbilledChargeInvoiceNowEstimateRequest
type UnbilledChargeInvoiceNowEstimateRequest struct {

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// subscription id
	SubscriptionID string `json:"subscription_id,omitempty"`
}

// Validate validates this unbilled charge invoice now estimate request
func (m *UnbilledChargeInvoiceNowEstimateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UnbilledChargeInvoiceNowEstimateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnbilledChargeInvoiceNowEstimateRequest) UnmarshalBinary(b []byte) error {
	var res UnbilledChargeInvoiceNowEstimateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
