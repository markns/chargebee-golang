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

// PaymentMethod payment method
// swagger:model PaymentMethod
type PaymentMethod struct {

	// gateway
	Gateway string `json:"gateway,omitempty"`

	// gateway account id
	GatewayAccountID string `json:"gateway_account_id,omitempty"`

	// reference id
	ReferenceID string `json:"reference_id,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this payment method
func (m *PaymentMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
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

var paymentMethodTypeGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chargebee","stripe","wepay","braintree","authorize_net","paypal_pro","pin","eway","eway_rapid","worldpay","balanced_payments","beanstream","bluepay","elavon","first_data_global","hdfc","migs","nmi","ogone","paymill","paypal_payflow_pro","sage_pay","tco","wirecard","amazon_payments","paypal_express_checkout","gocardless","adyen","orbital","moneris_us","moneris","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentMethodTypeGatewayPropEnum = append(paymentMethodTypeGatewayPropEnum, v)
	}
}

const (

	// PaymentMethodGatewayChargebee captures enum value "chargebee"
	PaymentMethodGatewayChargebee string = "chargebee"

	// PaymentMethodGatewayStripe captures enum value "stripe"
	PaymentMethodGatewayStripe string = "stripe"

	// PaymentMethodGatewayWepay captures enum value "wepay"
	PaymentMethodGatewayWepay string = "wepay"

	// PaymentMethodGatewayBraintree captures enum value "braintree"
	PaymentMethodGatewayBraintree string = "braintree"

	// PaymentMethodGatewayAuthorizeNet captures enum value "authorize_net"
	PaymentMethodGatewayAuthorizeNet string = "authorize_net"

	// PaymentMethodGatewayPaypalPro captures enum value "paypal_pro"
	PaymentMethodGatewayPaypalPro string = "paypal_pro"

	// PaymentMethodGatewayPin captures enum value "pin"
	PaymentMethodGatewayPin string = "pin"

	// PaymentMethodGatewayEway captures enum value "eway"
	PaymentMethodGatewayEway string = "eway"

	// PaymentMethodGatewayEwayRapid captures enum value "eway_rapid"
	PaymentMethodGatewayEwayRapid string = "eway_rapid"

	// PaymentMethodGatewayWorldpay captures enum value "worldpay"
	PaymentMethodGatewayWorldpay string = "worldpay"

	// PaymentMethodGatewayBalancedPayments captures enum value "balanced_payments"
	PaymentMethodGatewayBalancedPayments string = "balanced_payments"

	// PaymentMethodGatewayBeanstream captures enum value "beanstream"
	PaymentMethodGatewayBeanstream string = "beanstream"

	// PaymentMethodGatewayBluepay captures enum value "bluepay"
	PaymentMethodGatewayBluepay string = "bluepay"

	// PaymentMethodGatewayElavon captures enum value "elavon"
	PaymentMethodGatewayElavon string = "elavon"

	// PaymentMethodGatewayFirstDataGlobal captures enum value "first_data_global"
	PaymentMethodGatewayFirstDataGlobal string = "first_data_global"

	// PaymentMethodGatewayHdfc captures enum value "hdfc"
	PaymentMethodGatewayHdfc string = "hdfc"

	// PaymentMethodGatewayMigs captures enum value "migs"
	PaymentMethodGatewayMigs string = "migs"

	// PaymentMethodGatewayNmi captures enum value "nmi"
	PaymentMethodGatewayNmi string = "nmi"

	// PaymentMethodGatewayOgone captures enum value "ogone"
	PaymentMethodGatewayOgone string = "ogone"

	// PaymentMethodGatewayPaymill captures enum value "paymill"
	PaymentMethodGatewayPaymill string = "paymill"

	// PaymentMethodGatewayPaypalPayflowPro captures enum value "paypal_payflow_pro"
	PaymentMethodGatewayPaypalPayflowPro string = "paypal_payflow_pro"

	// PaymentMethodGatewaySagePay captures enum value "sage_pay"
	PaymentMethodGatewaySagePay string = "sage_pay"

	// PaymentMethodGatewayTco captures enum value "tco"
	PaymentMethodGatewayTco string = "tco"

	// PaymentMethodGatewayWirecard captures enum value "wirecard"
	PaymentMethodGatewayWirecard string = "wirecard"

	// PaymentMethodGatewayAmazonPayments captures enum value "amazon_payments"
	PaymentMethodGatewayAmazonPayments string = "amazon_payments"

	// PaymentMethodGatewayPaypalExpressCheckout captures enum value "paypal_express_checkout"
	PaymentMethodGatewayPaypalExpressCheckout string = "paypal_express_checkout"

	// PaymentMethodGatewayGocardless captures enum value "gocardless"
	PaymentMethodGatewayGocardless string = "gocardless"

	// PaymentMethodGatewayAdyen captures enum value "adyen"
	PaymentMethodGatewayAdyen string = "adyen"

	// PaymentMethodGatewayOrbital captures enum value "orbital"
	PaymentMethodGatewayOrbital string = "orbital"

	// PaymentMethodGatewayMonerisUs captures enum value "moneris_us"
	PaymentMethodGatewayMonerisUs string = "moneris_us"

	// PaymentMethodGatewayMoneris captures enum value "moneris"
	PaymentMethodGatewayMoneris string = "moneris"

	// PaymentMethodGatewayNotApplicable captures enum value "not_applicable"
	PaymentMethodGatewayNotApplicable string = "not_applicable"
)

// prop value enum
func (m *PaymentMethod) validateGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentMethodTypeGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentMethod) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	// value enum
	if err := m.validateGatewayEnum("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

var paymentMethodTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["future","in_trial","active","non_renewing","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentMethodTypeStatusPropEnum = append(paymentMethodTypeStatusPropEnum, v)
	}
}

const (

	// PaymentMethodStatusFuture captures enum value "future"
	PaymentMethodStatusFuture string = "future"

	// PaymentMethodStatusInTrial captures enum value "in_trial"
	PaymentMethodStatusInTrial string = "in_trial"

	// PaymentMethodStatusActive captures enum value "active"
	PaymentMethodStatusActive string = "active"

	// PaymentMethodStatusNonRenewing captures enum value "non_renewing"
	PaymentMethodStatusNonRenewing string = "non_renewing"

	// PaymentMethodStatusCancelled captures enum value "cancelled"
	PaymentMethodStatusCancelled string = "cancelled"
)

// prop value enum
func (m *PaymentMethod) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentMethodTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentMethod) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var paymentMethodTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["card","paypal_express_checkout","amazon_payments","direct_debit","generic","alipay","unionpay","apple_pay"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentMethodTypeTypePropEnum = append(paymentMethodTypeTypePropEnum, v)
	}
}

const (

	// PaymentMethodTypeCard captures enum value "card"
	PaymentMethodTypeCard string = "card"

	// PaymentMethodTypePaypalExpressCheckout captures enum value "paypal_express_checkout"
	PaymentMethodTypePaypalExpressCheckout string = "paypal_express_checkout"

	// PaymentMethodTypeAmazonPayments captures enum value "amazon_payments"
	PaymentMethodTypeAmazonPayments string = "amazon_payments"

	// PaymentMethodTypeDirectDebit captures enum value "direct_debit"
	PaymentMethodTypeDirectDebit string = "direct_debit"

	// PaymentMethodTypeGeneric captures enum value "generic"
	PaymentMethodTypeGeneric string = "generic"

	// PaymentMethodTypeAlipay captures enum value "alipay"
	PaymentMethodTypeAlipay string = "alipay"

	// PaymentMethodTypeUnionpay captures enum value "unionpay"
	PaymentMethodTypeUnionpay string = "unionpay"

	// PaymentMethodTypeApplePay captures enum value "apple_pay"
	PaymentMethodTypeApplePay string = "apple_pay"
)

// prop value enum
func (m *PaymentMethod) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentMethodTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentMethod) validateType(formats strfmt.Registry) error {

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
func (m *PaymentMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentMethod) UnmarshalBinary(b []byte) error {
	var res PaymentMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
