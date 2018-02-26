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

// PromotionalCreditDeductRequest promotional credit deduct request
// swagger:model PromotionalCreditDeductRequest

type PromotionalCreditDeductRequest struct {

	// amount
	Amount int32 `json:"amount,omitempty"`

	// credit type
	CreditType string `json:"credit_type,omitempty"`

	// currency code
	CurrencyCode string `json:"currency_code,omitempty"`

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// reference
	Reference string `json:"reference,omitempty"`
}

/* polymorph PromotionalCreditDeductRequest amount false */

/* polymorph PromotionalCreditDeductRequest credit_type false */

/* polymorph PromotionalCreditDeductRequest currency_code false */

/* polymorph PromotionalCreditDeductRequest customer_id false */

/* polymorph PromotionalCreditDeductRequest description false */

/* polymorph PromotionalCreditDeductRequest reference false */

// Validate validates this promotional credit deduct request
func (m *PromotionalCreditDeductRequest) Validate(formats strfmt.Registry) error {
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

var promotionalCreditDeductRequestTypeCreditTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["loyalty_credits","referral_rewards","general"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		promotionalCreditDeductRequestTypeCreditTypePropEnum = append(promotionalCreditDeductRequestTypeCreditTypePropEnum, v)
	}
}

const (
	// PromotionalCreditDeductRequestCreditTypeLoyaltyCredits captures enum value "loyalty_credits"
	PromotionalCreditDeductRequestCreditTypeLoyaltyCredits string = "loyalty_credits"
	// PromotionalCreditDeductRequestCreditTypeReferralRewards captures enum value "referral_rewards"
	PromotionalCreditDeductRequestCreditTypeReferralRewards string = "referral_rewards"
	// PromotionalCreditDeductRequestCreditTypeGeneral captures enum value "general"
	PromotionalCreditDeductRequestCreditTypeGeneral string = "general"
)

// prop value enum
func (m *PromotionalCreditDeductRequest) validateCreditTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, promotionalCreditDeductRequestTypeCreditTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PromotionalCreditDeductRequest) validateCreditType(formats strfmt.Registry) error {

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
func (m *PromotionalCreditDeductRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PromotionalCreditDeductRequest) UnmarshalBinary(b []byte) error {
	var res PromotionalCreditDeductRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
