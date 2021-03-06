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

// CustomerSetPromotionalCreditsRequest customer set promotional credits request
// swagger:model CustomerSetPromotionalCreditsRequest
type CustomerSetPromotionalCreditsRequest struct {

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

// Validate validates this customer set promotional credits request
func (m *CustomerSetPromotionalCreditsRequest) Validate(formats strfmt.Registry) error {
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

var customerSetPromotionalCreditsRequestTypeCreditTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["loyalty_credits","referral_rewards","general"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerSetPromotionalCreditsRequestTypeCreditTypePropEnum = append(customerSetPromotionalCreditsRequestTypeCreditTypePropEnum, v)
	}
}

const (

	// CustomerSetPromotionalCreditsRequestCreditTypeLoyaltyCredits captures enum value "loyalty_credits"
	CustomerSetPromotionalCreditsRequestCreditTypeLoyaltyCredits string = "loyalty_credits"

	// CustomerSetPromotionalCreditsRequestCreditTypeReferralRewards captures enum value "referral_rewards"
	CustomerSetPromotionalCreditsRequestCreditTypeReferralRewards string = "referral_rewards"

	// CustomerSetPromotionalCreditsRequestCreditTypeGeneral captures enum value "general"
	CustomerSetPromotionalCreditsRequestCreditTypeGeneral string = "general"
)

// prop value enum
func (m *CustomerSetPromotionalCreditsRequest) validateCreditTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerSetPromotionalCreditsRequestTypeCreditTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomerSetPromotionalCreditsRequest) validateCreditType(formats strfmt.Registry) error {

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
func (m *CustomerSetPromotionalCreditsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerSetPromotionalCreditsRequest) UnmarshalBinary(b []byte) error {
	var res CustomerSetPromotionalCreditsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
