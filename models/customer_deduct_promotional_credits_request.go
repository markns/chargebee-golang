// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomerDeductPromotionalCreditsRequest customer deduct promotional credits request
// swagger:model CustomerDeductPromotionalCreditsRequest

type CustomerDeductPromotionalCreditsRequest struct {

	// amount
	Amount int32 `json:"amount,omitempty"`

	// credit type
	CreditType string `json:"credit_type,omitempty"`

	// currency code
	CurrencyCode string `json:"currency_code,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// reference
	Reference string `json:"reference,omitempty"`
}

/* polymorph CustomerDeductPromotionalCreditsRequest amount false */

/* polymorph CustomerDeductPromotionalCreditsRequest credit_type false */

/* polymorph CustomerDeductPromotionalCreditsRequest currency_code false */

/* polymorph CustomerDeductPromotionalCreditsRequest description false */

/* polymorph CustomerDeductPromotionalCreditsRequest reference false */

// Validate validates this customer deduct promotional credits request
func (m *CustomerDeductPromotionalCreditsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreditType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customerDeductPromotionalCreditsRequestTypeCreditTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["loyalty_credits","referral_rewards","general"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerDeductPromotionalCreditsRequestTypeCreditTypePropEnum = append(customerDeductPromotionalCreditsRequestTypeCreditTypePropEnum, v)
	}
}

const (
	// CustomerDeductPromotionalCreditsRequestCreditTypeLoyaltyCredits captures enum value "loyalty_credits"
	CustomerDeductPromotionalCreditsRequestCreditTypeLoyaltyCredits string = "loyalty_credits"
	// CustomerDeductPromotionalCreditsRequestCreditTypeReferralRewards captures enum value "referral_rewards"
	CustomerDeductPromotionalCreditsRequestCreditTypeReferralRewards string = "referral_rewards"
	// CustomerDeductPromotionalCreditsRequestCreditTypeGeneral captures enum value "general"
	CustomerDeductPromotionalCreditsRequestCreditTypeGeneral string = "general"
)

// prop value enum
func (m *CustomerDeductPromotionalCreditsRequest) validateCreditTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerDeductPromotionalCreditsRequestTypeCreditTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomerDeductPromotionalCreditsRequest) validateCreditType(formats strfmt.Registry) error {

	if swag.IsZero(m.CreditType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCreditTypeEnum("credit_type", "body", m.CreditType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerDeductPromotionalCreditsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerDeductPromotionalCreditsRequest) UnmarshalBinary(b []byte) error {
	var res CustomerDeductPromotionalCreditsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
