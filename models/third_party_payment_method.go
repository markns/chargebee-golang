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

// ThirdPartyPaymentMethod third party payment method
// swagger:model ThirdPartyPaymentMethod

type ThirdPartyPaymentMethod struct {

	// gateway
	Gateway string `json:"gateway,omitempty"`

	// gateway account id
	GatewayAccountID string `json:"gateway_account_id,omitempty"`

	// reference id
	ReferenceID string `json:"reference_id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

/* polymorph ThirdPartyPaymentMethod gateway false */

/* polymorph ThirdPartyPaymentMethod gateway_account_id false */

/* polymorph ThirdPartyPaymentMethod reference_id false */

/* polymorph ThirdPartyPaymentMethod type false */

// Validate validates this third party payment method
func (m *ThirdPartyPaymentMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
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

var thirdPartyPaymentMethodTypeGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chargebee","stripe","wepay","braintree","authorize_net","paypal_pro","pin","eway","eway_rapid","worldpay","balanced_payments","beanstream","bluepay","elavon","first_data_global","hdfc","migs","nmi","ogone","paymill","paypal_payflow_pro","sage_pay","tco","wirecard","amazon_payments","paypal_express_checkout","gocardless","adyen","orbital","moneris_us","moneris","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		thirdPartyPaymentMethodTypeGatewayPropEnum = append(thirdPartyPaymentMethodTypeGatewayPropEnum, v)
	}
}

const (
	// ThirdPartyPaymentMethodGatewayChargebee captures enum value "chargebee"
	ThirdPartyPaymentMethodGatewayChargebee string = "chargebee"
	// ThirdPartyPaymentMethodGatewayStripe captures enum value "stripe"
	ThirdPartyPaymentMethodGatewayStripe string = "stripe"
	// ThirdPartyPaymentMethodGatewayWepay captures enum value "wepay"
	ThirdPartyPaymentMethodGatewayWepay string = "wepay"
	// ThirdPartyPaymentMethodGatewayBraintree captures enum value "braintree"
	ThirdPartyPaymentMethodGatewayBraintree string = "braintree"
	// ThirdPartyPaymentMethodGatewayAuthorizeNet captures enum value "authorize_net"
	ThirdPartyPaymentMethodGatewayAuthorizeNet string = "authorize_net"
	// ThirdPartyPaymentMethodGatewayPaypalPro captures enum value "paypal_pro"
	ThirdPartyPaymentMethodGatewayPaypalPro string = "paypal_pro"
	// ThirdPartyPaymentMethodGatewayPin captures enum value "pin"
	ThirdPartyPaymentMethodGatewayPin string = "pin"
	// ThirdPartyPaymentMethodGatewayEway captures enum value "eway"
	ThirdPartyPaymentMethodGatewayEway string = "eway"
	// ThirdPartyPaymentMethodGatewayEwayRapid captures enum value "eway_rapid"
	ThirdPartyPaymentMethodGatewayEwayRapid string = "eway_rapid"
	// ThirdPartyPaymentMethodGatewayWorldpay captures enum value "worldpay"
	ThirdPartyPaymentMethodGatewayWorldpay string = "worldpay"
	// ThirdPartyPaymentMethodGatewayBalancedPayments captures enum value "balanced_payments"
	ThirdPartyPaymentMethodGatewayBalancedPayments string = "balanced_payments"
	// ThirdPartyPaymentMethodGatewayBeanstream captures enum value "beanstream"
	ThirdPartyPaymentMethodGatewayBeanstream string = "beanstream"
	// ThirdPartyPaymentMethodGatewayBluepay captures enum value "bluepay"
	ThirdPartyPaymentMethodGatewayBluepay string = "bluepay"
	// ThirdPartyPaymentMethodGatewayElavon captures enum value "elavon"
	ThirdPartyPaymentMethodGatewayElavon string = "elavon"
	// ThirdPartyPaymentMethodGatewayFirstDataGlobal captures enum value "first_data_global"
	ThirdPartyPaymentMethodGatewayFirstDataGlobal string = "first_data_global"
	// ThirdPartyPaymentMethodGatewayHdfc captures enum value "hdfc"
	ThirdPartyPaymentMethodGatewayHdfc string = "hdfc"
	// ThirdPartyPaymentMethodGatewayMigs captures enum value "migs"
	ThirdPartyPaymentMethodGatewayMigs string = "migs"
	// ThirdPartyPaymentMethodGatewayNmi captures enum value "nmi"
	ThirdPartyPaymentMethodGatewayNmi string = "nmi"
	// ThirdPartyPaymentMethodGatewayOgone captures enum value "ogone"
	ThirdPartyPaymentMethodGatewayOgone string = "ogone"
	// ThirdPartyPaymentMethodGatewayPaymill captures enum value "paymill"
	ThirdPartyPaymentMethodGatewayPaymill string = "paymill"
	// ThirdPartyPaymentMethodGatewayPaypalPayflowPro captures enum value "paypal_payflow_pro"
	ThirdPartyPaymentMethodGatewayPaypalPayflowPro string = "paypal_payflow_pro"
	// ThirdPartyPaymentMethodGatewaySagePay captures enum value "sage_pay"
	ThirdPartyPaymentMethodGatewaySagePay string = "sage_pay"
	// ThirdPartyPaymentMethodGatewayTco captures enum value "tco"
	ThirdPartyPaymentMethodGatewayTco string = "tco"
	// ThirdPartyPaymentMethodGatewayWirecard captures enum value "wirecard"
	ThirdPartyPaymentMethodGatewayWirecard string = "wirecard"
	// ThirdPartyPaymentMethodGatewayAmazonPayments captures enum value "amazon_payments"
	ThirdPartyPaymentMethodGatewayAmazonPayments string = "amazon_payments"
	// ThirdPartyPaymentMethodGatewayPaypalExpressCheckout captures enum value "paypal_express_checkout"
	ThirdPartyPaymentMethodGatewayPaypalExpressCheckout string = "paypal_express_checkout"
	// ThirdPartyPaymentMethodGatewayGocardless captures enum value "gocardless"
	ThirdPartyPaymentMethodGatewayGocardless string = "gocardless"
	// ThirdPartyPaymentMethodGatewayAdyen captures enum value "adyen"
	ThirdPartyPaymentMethodGatewayAdyen string = "adyen"
	// ThirdPartyPaymentMethodGatewayOrbital captures enum value "orbital"
	ThirdPartyPaymentMethodGatewayOrbital string = "orbital"
	// ThirdPartyPaymentMethodGatewayMonerisUs captures enum value "moneris_us"
	ThirdPartyPaymentMethodGatewayMonerisUs string = "moneris_us"
	// ThirdPartyPaymentMethodGatewayMoneris captures enum value "moneris"
	ThirdPartyPaymentMethodGatewayMoneris string = "moneris"
	// ThirdPartyPaymentMethodGatewayNotApplicable captures enum value "not_applicable"
	ThirdPartyPaymentMethodGatewayNotApplicable string = "not_applicable"
)

// prop value enum
func (m *ThirdPartyPaymentMethod) validateGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, thirdPartyPaymentMethodTypeGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ThirdPartyPaymentMethod) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	// value enum
	if err := m.validateGatewayEnum("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

var thirdPartyPaymentMethodTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["card","paypal_express_checkout","amazon_payments","direct_debit","generic","alipay","unionpay","apple_pay"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		thirdPartyPaymentMethodTypeTypePropEnum = append(thirdPartyPaymentMethodTypeTypePropEnum, v)
	}
}

const (
	// ThirdPartyPaymentMethodTypeCard captures enum value "card"
	ThirdPartyPaymentMethodTypeCard string = "card"
	// ThirdPartyPaymentMethodTypePaypalExpressCheckout captures enum value "paypal_express_checkout"
	ThirdPartyPaymentMethodTypePaypalExpressCheckout string = "paypal_express_checkout"
	// ThirdPartyPaymentMethodTypeAmazonPayments captures enum value "amazon_payments"
	ThirdPartyPaymentMethodTypeAmazonPayments string = "amazon_payments"
	// ThirdPartyPaymentMethodTypeDirectDebit captures enum value "direct_debit"
	ThirdPartyPaymentMethodTypeDirectDebit string = "direct_debit"
	// ThirdPartyPaymentMethodTypeGeneric captures enum value "generic"
	ThirdPartyPaymentMethodTypeGeneric string = "generic"
	// ThirdPartyPaymentMethodTypeAlipay captures enum value "alipay"
	ThirdPartyPaymentMethodTypeAlipay string = "alipay"
	// ThirdPartyPaymentMethodTypeUnionpay captures enum value "unionpay"
	ThirdPartyPaymentMethodTypeUnionpay string = "unionpay"
	// ThirdPartyPaymentMethodTypeApplePay captures enum value "apple_pay"
	ThirdPartyPaymentMethodTypeApplePay string = "apple_pay"
)

// prop value enum
func (m *ThirdPartyPaymentMethod) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, thirdPartyPaymentMethodTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ThirdPartyPaymentMethod) validateType(formats strfmt.Registry) error {

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
func (m *ThirdPartyPaymentMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThirdPartyPaymentMethod) UnmarshalBinary(b []byte) error {
	var res ThirdPartyPaymentMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
