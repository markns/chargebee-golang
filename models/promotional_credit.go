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

// PromotionalCredit promotional credit
// swagger:model PromotionalCredit
type PromotionalCredit struct {

	// amount
	Amount int32 `json:"amount,omitempty"`

	// closing balance
	ClosingBalance int32 `json:"closing_balance,omitempty"`

	// created at
	CreatedAt int64 `json:"created_at,omitempty"`

	// credit type
	CreditType string `json:"credit_type,omitempty"`

	// currency code
	CurrencyCode string `json:"currency_code,omitempty"`

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// reference
	Reference string `json:"reference,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this promotional credit
func (m *PromotionalCredit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreditType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var promotionalCreditTypeCreditTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["loyalty_credits","referral_rewards","general"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		promotionalCreditTypeCreditTypePropEnum = append(promotionalCreditTypeCreditTypePropEnum, v)
	}
}

const (

	// PromotionalCreditCreditTypeLoyaltyCredits captures enum value "loyalty_credits"
	PromotionalCreditCreditTypeLoyaltyCredits string = "loyalty_credits"

	// PromotionalCreditCreditTypeReferralRewards captures enum value "referral_rewards"
	PromotionalCreditCreditTypeReferralRewards string = "referral_rewards"

	// PromotionalCreditCreditTypeGeneral captures enum value "general"
	PromotionalCreditCreditTypeGeneral string = "general"
)

// prop value enum
func (m *PromotionalCredit) validateCreditTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, promotionalCreditTypeCreditTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PromotionalCredit) validateCreditType(formats strfmt.Registry) error {

	if swag.IsZero(m.CreditType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCreditTypeEnum("credit_type", "body", m.CreditType); err != nil {
		return err
	}

	return nil
}

var promotionalCreditTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["card","paypal_express_checkout","amazon_payments","direct_debit","generic","alipay","unionpay","apple_pay"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		promotionalCreditTypeTypePropEnum = append(promotionalCreditTypeTypePropEnum, v)
	}
}

const (

	// PromotionalCreditTypeCard captures enum value "card"
	PromotionalCreditTypeCard string = "card"

	// PromotionalCreditTypePaypalExpressCheckout captures enum value "paypal_express_checkout"
	PromotionalCreditTypePaypalExpressCheckout string = "paypal_express_checkout"

	// PromotionalCreditTypeAmazonPayments captures enum value "amazon_payments"
	PromotionalCreditTypeAmazonPayments string = "amazon_payments"

	// PromotionalCreditTypeDirectDebit captures enum value "direct_debit"
	PromotionalCreditTypeDirectDebit string = "direct_debit"

	// PromotionalCreditTypeGeneric captures enum value "generic"
	PromotionalCreditTypeGeneric string = "generic"

	// PromotionalCreditTypeAlipay captures enum value "alipay"
	PromotionalCreditTypeAlipay string = "alipay"

	// PromotionalCreditTypeUnionpay captures enum value "unionpay"
	PromotionalCreditTypeUnionpay string = "unionpay"

	// PromotionalCreditTypeApplePay captures enum value "apple_pay"
	PromotionalCreditTypeApplePay string = "apple_pay"
)

// prop value enum
func (m *PromotionalCredit) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, promotionalCreditTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PromotionalCredit) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PromotionalCredit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PromotionalCredit) UnmarshalBinary(b []byte) error {
	var res PromotionalCredit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
