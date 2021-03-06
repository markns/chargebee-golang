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

// PaymentSource payment source
// swagger:model PaymentSource
type PaymentSource struct {

	// amazon payment
	AmazonPayment *AmazonPayment `json:"amazon_payment,omitempty"`

	// bank account
	BankAccount *BankAccount `json:"bank_account,omitempty"`

	// card
	Card *Card `json:"card,omitempty"`

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// gateway
	Gateway string `json:"gateway,omitempty"`

	// gateway account id
	GatewayAccountID string `json:"gateway_account_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// ip address
	IPAddress string `json:"ip_address,omitempty"`

	// issuing country
	IssuingCountry string `json:"issuing_country,omitempty"`

	// paypal
	Paypal *Paypal `json:"paypal,omitempty"`

	// reference id
	ReferenceID string `json:"reference_id,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this payment source
func (m *PaymentSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmazonPayment(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBankAccount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCard(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaypal(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *PaymentSource) validateAmazonPayment(formats strfmt.Registry) error {

	if swag.IsZero(m.AmazonPayment) { // not required
		return nil
	}

	if m.AmazonPayment != nil {

		if err := m.AmazonPayment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amazon_payment")
			}
			return err
		}

	}

	return nil
}

func (m *PaymentSource) validateBankAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.BankAccount) { // not required
		return nil
	}

	if m.BankAccount != nil {

		if err := m.BankAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bank_account")
			}
			return err
		}

	}

	return nil
}

func (m *PaymentSource) validateCard(formats strfmt.Registry) error {

	if swag.IsZero(m.Card) { // not required
		return nil
	}

	if m.Card != nil {

		if err := m.Card.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card")
			}
			return err
		}

	}

	return nil
}

var paymentSourceTypeGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chargebee","stripe","wepay","braintree","authorize_net","paypal_pro","pin","eway","eway_rapid","worldpay","balanced_payments","beanstream","bluepay","elavon","first_data_global","hdfc","migs","nmi","ogone","paymill","paypal_payflow_pro","sage_pay","tco","wirecard","amazon_payments","paypal_express_checkout","gocardless","adyen","orbital","moneris_us","moneris","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentSourceTypeGatewayPropEnum = append(paymentSourceTypeGatewayPropEnum, v)
	}
}

const (

	// PaymentSourceGatewayChargebee captures enum value "chargebee"
	PaymentSourceGatewayChargebee string = "chargebee"

	// PaymentSourceGatewayStripe captures enum value "stripe"
	PaymentSourceGatewayStripe string = "stripe"

	// PaymentSourceGatewayWepay captures enum value "wepay"
	PaymentSourceGatewayWepay string = "wepay"

	// PaymentSourceGatewayBraintree captures enum value "braintree"
	PaymentSourceGatewayBraintree string = "braintree"

	// PaymentSourceGatewayAuthorizeNet captures enum value "authorize_net"
	PaymentSourceGatewayAuthorizeNet string = "authorize_net"

	// PaymentSourceGatewayPaypalPro captures enum value "paypal_pro"
	PaymentSourceGatewayPaypalPro string = "paypal_pro"

	// PaymentSourceGatewayPin captures enum value "pin"
	PaymentSourceGatewayPin string = "pin"

	// PaymentSourceGatewayEway captures enum value "eway"
	PaymentSourceGatewayEway string = "eway"

	// PaymentSourceGatewayEwayRapid captures enum value "eway_rapid"
	PaymentSourceGatewayEwayRapid string = "eway_rapid"

	// PaymentSourceGatewayWorldpay captures enum value "worldpay"
	PaymentSourceGatewayWorldpay string = "worldpay"

	// PaymentSourceGatewayBalancedPayments captures enum value "balanced_payments"
	PaymentSourceGatewayBalancedPayments string = "balanced_payments"

	// PaymentSourceGatewayBeanstream captures enum value "beanstream"
	PaymentSourceGatewayBeanstream string = "beanstream"

	// PaymentSourceGatewayBluepay captures enum value "bluepay"
	PaymentSourceGatewayBluepay string = "bluepay"

	// PaymentSourceGatewayElavon captures enum value "elavon"
	PaymentSourceGatewayElavon string = "elavon"

	// PaymentSourceGatewayFirstDataGlobal captures enum value "first_data_global"
	PaymentSourceGatewayFirstDataGlobal string = "first_data_global"

	// PaymentSourceGatewayHdfc captures enum value "hdfc"
	PaymentSourceGatewayHdfc string = "hdfc"

	// PaymentSourceGatewayMigs captures enum value "migs"
	PaymentSourceGatewayMigs string = "migs"

	// PaymentSourceGatewayNmi captures enum value "nmi"
	PaymentSourceGatewayNmi string = "nmi"

	// PaymentSourceGatewayOgone captures enum value "ogone"
	PaymentSourceGatewayOgone string = "ogone"

	// PaymentSourceGatewayPaymill captures enum value "paymill"
	PaymentSourceGatewayPaymill string = "paymill"

	// PaymentSourceGatewayPaypalPayflowPro captures enum value "paypal_payflow_pro"
	PaymentSourceGatewayPaypalPayflowPro string = "paypal_payflow_pro"

	// PaymentSourceGatewaySagePay captures enum value "sage_pay"
	PaymentSourceGatewaySagePay string = "sage_pay"

	// PaymentSourceGatewayTco captures enum value "tco"
	PaymentSourceGatewayTco string = "tco"

	// PaymentSourceGatewayWirecard captures enum value "wirecard"
	PaymentSourceGatewayWirecard string = "wirecard"

	// PaymentSourceGatewayAmazonPayments captures enum value "amazon_payments"
	PaymentSourceGatewayAmazonPayments string = "amazon_payments"

	// PaymentSourceGatewayPaypalExpressCheckout captures enum value "paypal_express_checkout"
	PaymentSourceGatewayPaypalExpressCheckout string = "paypal_express_checkout"

	// PaymentSourceGatewayGocardless captures enum value "gocardless"
	PaymentSourceGatewayGocardless string = "gocardless"

	// PaymentSourceGatewayAdyen captures enum value "adyen"
	PaymentSourceGatewayAdyen string = "adyen"

	// PaymentSourceGatewayOrbital captures enum value "orbital"
	PaymentSourceGatewayOrbital string = "orbital"

	// PaymentSourceGatewayMonerisUs captures enum value "moneris_us"
	PaymentSourceGatewayMonerisUs string = "moneris_us"

	// PaymentSourceGatewayMoneris captures enum value "moneris"
	PaymentSourceGatewayMoneris string = "moneris"

	// PaymentSourceGatewayNotApplicable captures enum value "not_applicable"
	PaymentSourceGatewayNotApplicable string = "not_applicable"
)

// prop value enum
func (m *PaymentSource) validateGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentSourceTypeGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentSource) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	// value enum
	if err := m.validateGatewayEnum("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSource) validatePaypal(formats strfmt.Registry) error {

	if swag.IsZero(m.Paypal) { // not required
		return nil
	}

	if m.Paypal != nil {

		if err := m.Paypal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paypal")
			}
			return err
		}

	}

	return nil
}

var paymentSourceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["future","in_trial","active","non_renewing","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentSourceTypeStatusPropEnum = append(paymentSourceTypeStatusPropEnum, v)
	}
}

const (

	// PaymentSourceStatusFuture captures enum value "future"
	PaymentSourceStatusFuture string = "future"

	// PaymentSourceStatusInTrial captures enum value "in_trial"
	PaymentSourceStatusInTrial string = "in_trial"

	// PaymentSourceStatusActive captures enum value "active"
	PaymentSourceStatusActive string = "active"

	// PaymentSourceStatusNonRenewing captures enum value "non_renewing"
	PaymentSourceStatusNonRenewing string = "non_renewing"

	// PaymentSourceStatusCancelled captures enum value "cancelled"
	PaymentSourceStatusCancelled string = "cancelled"
)

// prop value enum
func (m *PaymentSource) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentSourceTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentSource) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var paymentSourceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["card","paypal_express_checkout","amazon_payments","direct_debit","generic","alipay","unionpay","apple_pay"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentSourceTypeTypePropEnum = append(paymentSourceTypeTypePropEnum, v)
	}
}

const (

	// PaymentSourceTypeCard captures enum value "card"
	PaymentSourceTypeCard string = "card"

	// PaymentSourceTypePaypalExpressCheckout captures enum value "paypal_express_checkout"
	PaymentSourceTypePaypalExpressCheckout string = "paypal_express_checkout"

	// PaymentSourceTypeAmazonPayments captures enum value "amazon_payments"
	PaymentSourceTypeAmazonPayments string = "amazon_payments"

	// PaymentSourceTypeDirectDebit captures enum value "direct_debit"
	PaymentSourceTypeDirectDebit string = "direct_debit"

	// PaymentSourceTypeGeneric captures enum value "generic"
	PaymentSourceTypeGeneric string = "generic"

	// PaymentSourceTypeAlipay captures enum value "alipay"
	PaymentSourceTypeAlipay string = "alipay"

	// PaymentSourceTypeUnionpay captures enum value "unionpay"
	PaymentSourceTypeUnionpay string = "unionpay"

	// PaymentSourceTypeApplePay captures enum value "apple_pay"
	PaymentSourceTypeApplePay string = "apple_pay"
)

// prop value enum
func (m *PaymentSource) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentSourceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentSource) validateType(formats strfmt.Registry) error {

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
func (m *PaymentSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentSource) UnmarshalBinary(b []byte) error {
	var res PaymentSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
