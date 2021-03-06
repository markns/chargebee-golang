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

// CustomerUpdatePaymentMethodRequest customer update payment method request
// swagger:model CustomerUpdatePaymentMethodRequest
type CustomerUpdatePaymentMethodRequest struct {

	// payment method gateway
	PaymentMethodGateway string `json:"payment_method[gateway],omitempty"`

	// payment method gateway account id
	PaymentMethodGatewayAccountID string `json:"payment_method[gateway_account_id],omitempty"`

	// payment method issuing country
	PaymentMethodIssuingCountry string `json:"payment_method[issuing_country],omitempty"`

	// payment method reference id
	PaymentMethodReferenceID string `json:"payment_method[reference_id],omitempty"`

	// payment method tmp token
	PaymentMethodTmpToken string `json:"payment_method[tmp_token],omitempty"`

	// payment method type
	PaymentMethodType string `json:"payment_method[type],omitempty"`
}

// Validate validates this customer update payment method request
func (m *CustomerUpdatePaymentMethodRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentMethodGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentMethodType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customerUpdatePaymentMethodRequestTypePaymentMethodGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chargebee","stripe","wepay","braintree","authorize_net","paypal_pro","pin","eway","eway_rapid","worldpay","balanced_payments","beanstream","bluepay","elavon","first_data_global","hdfc","migs","nmi","ogone","paymill","paypal_payflow_pro","sage_pay","tco","wirecard","amazon_payments","paypal_express_checkout","gocardless","adyen","orbital","moneris_us","moneris","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerUpdatePaymentMethodRequestTypePaymentMethodGatewayPropEnum = append(customerUpdatePaymentMethodRequestTypePaymentMethodGatewayPropEnum, v)
	}
}

const (

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayChargebee captures enum value "chargebee"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayChargebee string = "chargebee"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayStripe captures enum value "stripe"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayStripe string = "stripe"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayWepay captures enum value "wepay"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayWepay string = "wepay"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayBraintree captures enum value "braintree"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayBraintree string = "braintree"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayAuthorizeNet captures enum value "authorize_net"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayAuthorizeNet string = "authorize_net"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayPaypalPro captures enum value "paypal_pro"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayPaypalPro string = "paypal_pro"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayPin captures enum value "pin"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayPin string = "pin"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayEway captures enum value "eway"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayEway string = "eway"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayEwayRapid captures enum value "eway_rapid"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayEwayRapid string = "eway_rapid"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayWorldpay captures enum value "worldpay"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayWorldpay string = "worldpay"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayBalancedPayments captures enum value "balanced_payments"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayBalancedPayments string = "balanced_payments"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayBeanstream captures enum value "beanstream"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayBeanstream string = "beanstream"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayBluepay captures enum value "bluepay"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayBluepay string = "bluepay"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayElavon captures enum value "elavon"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayElavon string = "elavon"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayFirstDataGlobal captures enum value "first_data_global"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayFirstDataGlobal string = "first_data_global"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayHdfc captures enum value "hdfc"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayHdfc string = "hdfc"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayMigs captures enum value "migs"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayMigs string = "migs"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayNmi captures enum value "nmi"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayNmi string = "nmi"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayOgone captures enum value "ogone"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayOgone string = "ogone"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayPaymill captures enum value "paymill"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayPaymill string = "paymill"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayPaypalPayflowPro captures enum value "paypal_payflow_pro"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayPaypalPayflowPro string = "paypal_payflow_pro"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewaySagePay captures enum value "sage_pay"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewaySagePay string = "sage_pay"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayTco captures enum value "tco"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayTco string = "tco"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayWirecard captures enum value "wirecard"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayWirecard string = "wirecard"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayAmazonPayments captures enum value "amazon_payments"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayAmazonPayments string = "amazon_payments"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayPaypalExpressCheckout captures enum value "paypal_express_checkout"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayPaypalExpressCheckout string = "paypal_express_checkout"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayGocardless captures enum value "gocardless"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayGocardless string = "gocardless"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayAdyen captures enum value "adyen"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayAdyen string = "adyen"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayOrbital captures enum value "orbital"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayOrbital string = "orbital"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayMonerisUs captures enum value "moneris_us"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayMonerisUs string = "moneris_us"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayMoneris captures enum value "moneris"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayMoneris string = "moneris"

	// CustomerUpdatePaymentMethodRequestPaymentMethodGatewayNotApplicable captures enum value "not_applicable"
	CustomerUpdatePaymentMethodRequestPaymentMethodGatewayNotApplicable string = "not_applicable"
)

// prop value enum
func (m *CustomerUpdatePaymentMethodRequest) validatePaymentMethodGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerUpdatePaymentMethodRequestTypePaymentMethodGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomerUpdatePaymentMethodRequest) validatePaymentMethodGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethodGateway) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentMethodGatewayEnum("payment_method[gateway]", "body", m.PaymentMethodGateway); err != nil {
		return err
	}

	return nil
}

var customerUpdatePaymentMethodRequestTypePaymentMethodTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["card","paypal_express_checkout","amazon_payments","direct_debit","generic","alipay","unionpay","apple_pay"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerUpdatePaymentMethodRequestTypePaymentMethodTypePropEnum = append(customerUpdatePaymentMethodRequestTypePaymentMethodTypePropEnum, v)
	}
}

const (

	// CustomerUpdatePaymentMethodRequestPaymentMethodTypeCard captures enum value "card"
	CustomerUpdatePaymentMethodRequestPaymentMethodTypeCard string = "card"

	// CustomerUpdatePaymentMethodRequestPaymentMethodTypePaypalExpressCheckout captures enum value "paypal_express_checkout"
	CustomerUpdatePaymentMethodRequestPaymentMethodTypePaypalExpressCheckout string = "paypal_express_checkout"

	// CustomerUpdatePaymentMethodRequestPaymentMethodTypeAmazonPayments captures enum value "amazon_payments"
	CustomerUpdatePaymentMethodRequestPaymentMethodTypeAmazonPayments string = "amazon_payments"

	// CustomerUpdatePaymentMethodRequestPaymentMethodTypeDirectDebit captures enum value "direct_debit"
	CustomerUpdatePaymentMethodRequestPaymentMethodTypeDirectDebit string = "direct_debit"

	// CustomerUpdatePaymentMethodRequestPaymentMethodTypeGeneric captures enum value "generic"
	CustomerUpdatePaymentMethodRequestPaymentMethodTypeGeneric string = "generic"

	// CustomerUpdatePaymentMethodRequestPaymentMethodTypeAlipay captures enum value "alipay"
	CustomerUpdatePaymentMethodRequestPaymentMethodTypeAlipay string = "alipay"

	// CustomerUpdatePaymentMethodRequestPaymentMethodTypeUnionpay captures enum value "unionpay"
	CustomerUpdatePaymentMethodRequestPaymentMethodTypeUnionpay string = "unionpay"

	// CustomerUpdatePaymentMethodRequestPaymentMethodTypeApplePay captures enum value "apple_pay"
	CustomerUpdatePaymentMethodRequestPaymentMethodTypeApplePay string = "apple_pay"
)

// prop value enum
func (m *CustomerUpdatePaymentMethodRequest) validatePaymentMethodTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerUpdatePaymentMethodRequestTypePaymentMethodTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomerUpdatePaymentMethodRequest) validatePaymentMethodType(formats strfmt.Registry) error {

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
func (m *CustomerUpdatePaymentMethodRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerUpdatePaymentMethodRequest) UnmarshalBinary(b []byte) error {
	var res CustomerUpdatePaymentMethodRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
