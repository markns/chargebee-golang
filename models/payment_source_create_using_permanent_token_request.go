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

// PaymentSourceCreateUsingPermanentTokenRequest payment source create using permanent token request
// swagger:model PaymentSourceCreateUsingPermanentTokenRequest
type PaymentSourceCreateUsingPermanentTokenRequest struct {

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// gateway account id
	GatewayAccountID string `json:"gateway_account_id,omitempty"`

	// issuing country
	IssuingCountry string `json:"issuing_country,omitempty"`

	// reference id
	ReferenceID string `json:"reference_id,omitempty"`

	// replace primary payment source
	ReplacePrimaryPaymentSource bool `json:"replace_primary_payment_source,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this payment source create using permanent token request
func (m *PaymentSourceCreateUsingPermanentTokenRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var paymentSourceCreateUsingPermanentTokenRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["card","paypal_express_checkout","amazon_payments","direct_debit","generic","alipay","unionpay","apple_pay"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentSourceCreateUsingPermanentTokenRequestTypeTypePropEnum = append(paymentSourceCreateUsingPermanentTokenRequestTypeTypePropEnum, v)
	}
}

const (

	// PaymentSourceCreateUsingPermanentTokenRequestTypeCard captures enum value "card"
	PaymentSourceCreateUsingPermanentTokenRequestTypeCard string = "card"

	// PaymentSourceCreateUsingPermanentTokenRequestTypePaypalExpressCheckout captures enum value "paypal_express_checkout"
	PaymentSourceCreateUsingPermanentTokenRequestTypePaypalExpressCheckout string = "paypal_express_checkout"

	// PaymentSourceCreateUsingPermanentTokenRequestTypeAmazonPayments captures enum value "amazon_payments"
	PaymentSourceCreateUsingPermanentTokenRequestTypeAmazonPayments string = "amazon_payments"

	// PaymentSourceCreateUsingPermanentTokenRequestTypeDirectDebit captures enum value "direct_debit"
	PaymentSourceCreateUsingPermanentTokenRequestTypeDirectDebit string = "direct_debit"

	// PaymentSourceCreateUsingPermanentTokenRequestTypeGeneric captures enum value "generic"
	PaymentSourceCreateUsingPermanentTokenRequestTypeGeneric string = "generic"

	// PaymentSourceCreateUsingPermanentTokenRequestTypeAlipay captures enum value "alipay"
	PaymentSourceCreateUsingPermanentTokenRequestTypeAlipay string = "alipay"

	// PaymentSourceCreateUsingPermanentTokenRequestTypeUnionpay captures enum value "unionpay"
	PaymentSourceCreateUsingPermanentTokenRequestTypeUnionpay string = "unionpay"

	// PaymentSourceCreateUsingPermanentTokenRequestTypeApplePay captures enum value "apple_pay"
	PaymentSourceCreateUsingPermanentTokenRequestTypeApplePay string = "apple_pay"
)

// prop value enum
func (m *PaymentSourceCreateUsingPermanentTokenRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentSourceCreateUsingPermanentTokenRequestTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentSourceCreateUsingPermanentTokenRequest) validateType(formats strfmt.Registry) error {

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
func (m *PaymentSourceCreateUsingPermanentTokenRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentSourceCreateUsingPermanentTokenRequest) UnmarshalBinary(b []byte) error {
	var res PaymentSourceCreateUsingPermanentTokenRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
