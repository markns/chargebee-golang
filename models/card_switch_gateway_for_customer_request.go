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

// CardSwitchGatewayForCustomerRequest card switch gateway for customer request
// swagger:model CardSwitchGatewayForCustomerRequest

type CardSwitchGatewayForCustomerRequest struct {

	// gateway
	Gateway string `json:"gateway,omitempty"`

	// gateway account id
	GatewayAccountID string `json:"gateway_account_id,omitempty"`
}

/* polymorph CardSwitchGatewayForCustomerRequest gateway false */

/* polymorph CardSwitchGatewayForCustomerRequest gateway_account_id false */

// Validate validates this card switch gateway for customer request
func (m *CardSwitchGatewayForCustomerRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cardSwitchGatewayForCustomerRequestTypeGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chargebee","stripe","wepay","braintree","authorize_net","paypal_pro","pin","eway","eway_rapid","worldpay","balanced_payments","beanstream","bluepay","elavon","first_data_global","hdfc","migs","nmi","ogone","paymill","paypal_payflow_pro","sage_pay","tco","wirecard","amazon_payments","paypal_express_checkout","gocardless","adyen","orbital","moneris_us","moneris","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cardSwitchGatewayForCustomerRequestTypeGatewayPropEnum = append(cardSwitchGatewayForCustomerRequestTypeGatewayPropEnum, v)
	}
}

const (
	// CardSwitchGatewayForCustomerRequestGatewayChargebee captures enum value "chargebee"
	CardSwitchGatewayForCustomerRequestGatewayChargebee string = "chargebee"
	// CardSwitchGatewayForCustomerRequestGatewayStripe captures enum value "stripe"
	CardSwitchGatewayForCustomerRequestGatewayStripe string = "stripe"
	// CardSwitchGatewayForCustomerRequestGatewayWepay captures enum value "wepay"
	CardSwitchGatewayForCustomerRequestGatewayWepay string = "wepay"
	// CardSwitchGatewayForCustomerRequestGatewayBraintree captures enum value "braintree"
	CardSwitchGatewayForCustomerRequestGatewayBraintree string = "braintree"
	// CardSwitchGatewayForCustomerRequestGatewayAuthorizeNet captures enum value "authorize_net"
	CardSwitchGatewayForCustomerRequestGatewayAuthorizeNet string = "authorize_net"
	// CardSwitchGatewayForCustomerRequestGatewayPaypalPro captures enum value "paypal_pro"
	CardSwitchGatewayForCustomerRequestGatewayPaypalPro string = "paypal_pro"
	// CardSwitchGatewayForCustomerRequestGatewayPin captures enum value "pin"
	CardSwitchGatewayForCustomerRequestGatewayPin string = "pin"
	// CardSwitchGatewayForCustomerRequestGatewayEway captures enum value "eway"
	CardSwitchGatewayForCustomerRequestGatewayEway string = "eway"
	// CardSwitchGatewayForCustomerRequestGatewayEwayRapid captures enum value "eway_rapid"
	CardSwitchGatewayForCustomerRequestGatewayEwayRapid string = "eway_rapid"
	// CardSwitchGatewayForCustomerRequestGatewayWorldpay captures enum value "worldpay"
	CardSwitchGatewayForCustomerRequestGatewayWorldpay string = "worldpay"
	// CardSwitchGatewayForCustomerRequestGatewayBalancedPayments captures enum value "balanced_payments"
	CardSwitchGatewayForCustomerRequestGatewayBalancedPayments string = "balanced_payments"
	// CardSwitchGatewayForCustomerRequestGatewayBeanstream captures enum value "beanstream"
	CardSwitchGatewayForCustomerRequestGatewayBeanstream string = "beanstream"
	// CardSwitchGatewayForCustomerRequestGatewayBluepay captures enum value "bluepay"
	CardSwitchGatewayForCustomerRequestGatewayBluepay string = "bluepay"
	// CardSwitchGatewayForCustomerRequestGatewayElavon captures enum value "elavon"
	CardSwitchGatewayForCustomerRequestGatewayElavon string = "elavon"
	// CardSwitchGatewayForCustomerRequestGatewayFirstDataGlobal captures enum value "first_data_global"
	CardSwitchGatewayForCustomerRequestGatewayFirstDataGlobal string = "first_data_global"
	// CardSwitchGatewayForCustomerRequestGatewayHdfc captures enum value "hdfc"
	CardSwitchGatewayForCustomerRequestGatewayHdfc string = "hdfc"
	// CardSwitchGatewayForCustomerRequestGatewayMigs captures enum value "migs"
	CardSwitchGatewayForCustomerRequestGatewayMigs string = "migs"
	// CardSwitchGatewayForCustomerRequestGatewayNmi captures enum value "nmi"
	CardSwitchGatewayForCustomerRequestGatewayNmi string = "nmi"
	// CardSwitchGatewayForCustomerRequestGatewayOgone captures enum value "ogone"
	CardSwitchGatewayForCustomerRequestGatewayOgone string = "ogone"
	// CardSwitchGatewayForCustomerRequestGatewayPaymill captures enum value "paymill"
	CardSwitchGatewayForCustomerRequestGatewayPaymill string = "paymill"
	// CardSwitchGatewayForCustomerRequestGatewayPaypalPayflowPro captures enum value "paypal_payflow_pro"
	CardSwitchGatewayForCustomerRequestGatewayPaypalPayflowPro string = "paypal_payflow_pro"
	// CardSwitchGatewayForCustomerRequestGatewaySagePay captures enum value "sage_pay"
	CardSwitchGatewayForCustomerRequestGatewaySagePay string = "sage_pay"
	// CardSwitchGatewayForCustomerRequestGatewayTco captures enum value "tco"
	CardSwitchGatewayForCustomerRequestGatewayTco string = "tco"
	// CardSwitchGatewayForCustomerRequestGatewayWirecard captures enum value "wirecard"
	CardSwitchGatewayForCustomerRequestGatewayWirecard string = "wirecard"
	// CardSwitchGatewayForCustomerRequestGatewayAmazonPayments captures enum value "amazon_payments"
	CardSwitchGatewayForCustomerRequestGatewayAmazonPayments string = "amazon_payments"
	// CardSwitchGatewayForCustomerRequestGatewayPaypalExpressCheckout captures enum value "paypal_express_checkout"
	CardSwitchGatewayForCustomerRequestGatewayPaypalExpressCheckout string = "paypal_express_checkout"
	// CardSwitchGatewayForCustomerRequestGatewayGocardless captures enum value "gocardless"
	CardSwitchGatewayForCustomerRequestGatewayGocardless string = "gocardless"
	// CardSwitchGatewayForCustomerRequestGatewayAdyen captures enum value "adyen"
	CardSwitchGatewayForCustomerRequestGatewayAdyen string = "adyen"
	// CardSwitchGatewayForCustomerRequestGatewayOrbital captures enum value "orbital"
	CardSwitchGatewayForCustomerRequestGatewayOrbital string = "orbital"
	// CardSwitchGatewayForCustomerRequestGatewayMonerisUs captures enum value "moneris_us"
	CardSwitchGatewayForCustomerRequestGatewayMonerisUs string = "moneris_us"
	// CardSwitchGatewayForCustomerRequestGatewayMoneris captures enum value "moneris"
	CardSwitchGatewayForCustomerRequestGatewayMoneris string = "moneris"
	// CardSwitchGatewayForCustomerRequestGatewayNotApplicable captures enum value "not_applicable"
	CardSwitchGatewayForCustomerRequestGatewayNotApplicable string = "not_applicable"
)

// prop value enum
func (m *CardSwitchGatewayForCustomerRequest) validateGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cardSwitchGatewayForCustomerRequestTypeGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CardSwitchGatewayForCustomerRequest) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	// value enum
	if err := m.validateGatewayEnum("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CardSwitchGatewayForCustomerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CardSwitchGatewayForCustomerRequest) UnmarshalBinary(b []byte) error {
	var res CardSwitchGatewayForCustomerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}