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

// CustomerCollectPaymentRequest customer collect payment request
// swagger:model CustomerCollectPaymentRequest

type CustomerCollectPaymentRequest struct {

	// amount
	Amount int32 `json:"amount,omitempty"`

	// card billing addr1
	CardBillingAddr1 string `json:"card[billing_addr1],omitempty"`

	// card billing addr2
	CardBillingAddr2 string `json:"card[billing_addr2],omitempty"`

	// card billing city
	CardBillingCity string `json:"card[billing_city],omitempty"`

	// card billing country
	CardBillingCountry string `json:"card[billing_country],omitempty"`

	// card billing state
	CardBillingState string `json:"card[billing_state],omitempty"`

	// card billing state code
	CardBillingStateCode string `json:"card[billing_state_code],omitempty"`

	// card billing zip
	CardBillingZip string `json:"card[billing_zip],omitempty"`

	// card cvv
	CardCvv string `json:"card[cvv],omitempty"`

	// card expiry month
	CardExpiryMonth int32 `json:"card[expiry_month],omitempty"`

	// card expiry year
	CardExpiryYear int32 `json:"card[expiry_year],omitempty"`

	// card first name
	CardFirstName string `json:"card[first_name],omitempty"`

	// card gateway account id
	CardGatewayAccountID string `json:"card[gateway_account_id],omitempty"`

	// card last name
	CardLastName string `json:"card[last_name],omitempty"`

	// card number
	CardNumber string `json:"card[number],omitempty"`

	// payment method gateway account id
	PaymentMethodGatewayAccountID string `json:"payment_method[gateway_account_id],omitempty"`

	// payment method reference id
	PaymentMethodReferenceID string `json:"payment_method[reference_id],omitempty"`

	// payment method tmp token
	PaymentMethodTmpToken string `json:"payment_method[tmp_token],omitempty"`

	// payment method type
	PaymentMethodType string `json:"payment_method[type],omitempty"`

	// payment source id
	PaymentSourceID string `json:"payment_source_id,omitempty"`

	// replace primary payment source
	ReplacePrimaryPaymentSource bool `json:"replace_primary_payment_source,omitempty"`

	// retain payment source
	RetainPaymentSource bool `json:"retain_payment_source,omitempty"`
}

/* polymorph CustomerCollectPaymentRequest amount false */

/* polymorph CustomerCollectPaymentRequest card[billing_addr1] false */

/* polymorph CustomerCollectPaymentRequest card[billing_addr2] false */

/* polymorph CustomerCollectPaymentRequest card[billing_city] false */

/* polymorph CustomerCollectPaymentRequest card[billing_country] false */

/* polymorph CustomerCollectPaymentRequest card[billing_state] false */

/* polymorph CustomerCollectPaymentRequest card[billing_state_code] false */

/* polymorph CustomerCollectPaymentRequest card[billing_zip] false */

/* polymorph CustomerCollectPaymentRequest card[cvv] false */

/* polymorph CustomerCollectPaymentRequest card[expiry_month] false */

/* polymorph CustomerCollectPaymentRequest card[expiry_year] false */

/* polymorph CustomerCollectPaymentRequest card[first_name] false */

/* polymorph CustomerCollectPaymentRequest card[gateway_account_id] false */

/* polymorph CustomerCollectPaymentRequest card[last_name] false */

/* polymorph CustomerCollectPaymentRequest card[number] false */

/* polymorph CustomerCollectPaymentRequest payment_method[gateway_account_id] false */

/* polymorph CustomerCollectPaymentRequest payment_method[reference_id] false */

/* polymorph CustomerCollectPaymentRequest payment_method[tmp_token] false */

/* polymorph CustomerCollectPaymentRequest payment_method[type] false */

/* polymorph CustomerCollectPaymentRequest payment_source_id false */

/* polymorph CustomerCollectPaymentRequest replace_primary_payment_source false */

/* polymorph CustomerCollectPaymentRequest retain_payment_source false */

// Validate validates this customer collect payment request
func (m *CustomerCollectPaymentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentMethodType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customerCollectPaymentRequestTypePaymentMethodTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["authorization","payment","refund","payment_reversal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerCollectPaymentRequestTypePaymentMethodTypePropEnum = append(customerCollectPaymentRequestTypePaymentMethodTypePropEnum, v)
	}
}

const (
	// CustomerCollectPaymentRequestPaymentMethodTypeAuthorization captures enum value "authorization"
	CustomerCollectPaymentRequestPaymentMethodTypeAuthorization string = "authorization"
	// CustomerCollectPaymentRequestPaymentMethodTypePayment captures enum value "payment"
	CustomerCollectPaymentRequestPaymentMethodTypePayment string = "payment"
	// CustomerCollectPaymentRequestPaymentMethodTypeRefund captures enum value "refund"
	CustomerCollectPaymentRequestPaymentMethodTypeRefund string = "refund"
	// CustomerCollectPaymentRequestPaymentMethodTypePaymentReversal captures enum value "payment_reversal"
	CustomerCollectPaymentRequestPaymentMethodTypePaymentReversal string = "payment_reversal"
)

// prop value enum
func (m *CustomerCollectPaymentRequest) validatePaymentMethodTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerCollectPaymentRequestTypePaymentMethodTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomerCollectPaymentRequest) validatePaymentMethodType(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethodType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentMethodTypeEnum("payment_method[type]", "body", m.PaymentMethodType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerCollectPaymentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerCollectPaymentRequest) UnmarshalBinary(b []byte) error {
	var res CustomerCollectPaymentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
