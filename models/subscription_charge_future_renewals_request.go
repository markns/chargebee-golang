// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SubscriptionChargeFutureRenewalsRequest subscription charge future renewals request
// swagger:model SubscriptionChargeFutureRenewalsRequest

type SubscriptionChargeFutureRenewalsRequest struct {

	// terms to charge
	TermsToCharge int32 `json:"terms_to_charge,omitempty"`
}

/* polymorph SubscriptionChargeFutureRenewalsRequest terms_to_charge false */

// Validate validates this subscription charge future renewals request
func (m *SubscriptionChargeFutureRenewalsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionChargeFutureRenewalsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionChargeFutureRenewalsRequest) UnmarshalBinary(b []byte) error {
	var res SubscriptionChargeFutureRenewalsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}