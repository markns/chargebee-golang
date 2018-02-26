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

// CustomerCreateRequest customer create request
// swagger:model CustomerCreateRequest

type CustomerCreateRequest struct {

	// allow direct debit
	AllowDirectDebit bool `json:"allow_direct_debit,omitempty"`

	// auto collection
	AutoCollection string `json:"auto_collection,omitempty"`

	// billing address city
	BillingAddressCity string `json:"billing_address[city],omitempty"`

	// billing address company
	BillingAddressCompany string `json:"billing_address[company],omitempty"`

	// billing address country
	BillingAddressCountry string `json:"billing_address[country],omitempty"`

	// billing address email
	BillingAddressEmail string `json:"billing_address[email],omitempty"`

	// billing address first name
	BillingAddressFirstName string `json:"billing_address[first_name],omitempty"`

	// billing address last name
	BillingAddressLastName string `json:"billing_address[last_name],omitempty"`

	// billing address line1
	BillingAddressLine1 string `json:"billing_address[line1],omitempty"`

	// billing address line2
	BillingAddressLine2 string `json:"billing_address[line2],omitempty"`

	// billing address line3
	BillingAddressLine3 string `json:"billing_address[line3],omitempty"`

	// billing address phone
	BillingAddressPhone string `json:"billing_address[phone],omitempty"`

	// billing address state
	BillingAddressState string `json:"billing_address[state],omitempty"`

	// billing address state code
	BillingAddressStateCode string `json:"billing_address[state_code],omitempty"`

	// billing address validation status
	BillingAddressValidationStatus string `json:"billing_address[validation_status],omitempty"`

	// billing address zip
	BillingAddressZip string `json:"billing_address[zip],omitempty"`

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

	// card gateway
	CardGateway string `json:"card[gateway],omitempty"`

	// card gateway account id
	CardGatewayAccountID string `json:"card[gateway_account_id],omitempty"`

	// card ip address
	CardIPAddress string `json:"card[ip_address],omitempty"`

	// card last name
	CardLastName string `json:"card[last_name],omitempty"`

	// card number
	CardNumber string `json:"card[number],omitempty"`

	// card tmp token
	CardTmpToken string `json:"card[tmp_token],omitempty"`

	// company
	Company string `json:"company,omitempty"`

	// consolidated invoicing
	ConsolidatedInvoicing bool `json:"consolidated_invoicing,omitempty"`

	// created from ip
	CreatedFromIP string `json:"created_from_ip,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// entity code
	EntityCode string `json:"entity_code,omitempty"`

	// exempt number
	ExemptNumber string `json:"exempt_number,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// invoice notes
	InvoiceNotes string `json:"invoice_notes,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// locale
	Locale string `json:"locale,omitempty"`

	// meta data
	MetaData string `json:"meta_data,omitempty"`

	// net term days
	NetTermDays int32 `json:"net_term_days,omitempty"`

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

	// phone
	Phone string `json:"phone,omitempty"`

	// preferred currency code
	PreferredCurrencyCode string `json:"preferred_currency_code,omitempty"`

	// registered for gst
	RegisteredForGst bool `json:"registered_for_gst,omitempty"`

	// taxability
	Taxability string `json:"taxability,omitempty"`

	// vat number
	VatNumber string `json:"vat_number,omitempty"`
}

/* polymorph CustomerCreateRequest allow_direct_debit false */

/* polymorph CustomerCreateRequest auto_collection false */

/* polymorph CustomerCreateRequest billing_address[city] false */

/* polymorph CustomerCreateRequest billing_address[company] false */

/* polymorph CustomerCreateRequest billing_address[country] false */

/* polymorph CustomerCreateRequest billing_address[email] false */

/* polymorph CustomerCreateRequest billing_address[first_name] false */

/* polymorph CustomerCreateRequest billing_address[last_name] false */

/* polymorph CustomerCreateRequest billing_address[line1] false */

/* polymorph CustomerCreateRequest billing_address[line2] false */

/* polymorph CustomerCreateRequest billing_address[line3] false */

/* polymorph CustomerCreateRequest billing_address[phone] false */

/* polymorph CustomerCreateRequest billing_address[state] false */

/* polymorph CustomerCreateRequest billing_address[state_code] false */

/* polymorph CustomerCreateRequest billing_address[validation_status] false */

/* polymorph CustomerCreateRequest billing_address[zip] false */

/* polymorph CustomerCreateRequest card[billing_addr1] false */

/* polymorph CustomerCreateRequest card[billing_addr2] false */

/* polymorph CustomerCreateRequest card[billing_city] false */

/* polymorph CustomerCreateRequest card[billing_country] false */

/* polymorph CustomerCreateRequest card[billing_state] false */

/* polymorph CustomerCreateRequest card[billing_state_code] false */

/* polymorph CustomerCreateRequest card[billing_zip] false */

/* polymorph CustomerCreateRequest card[cvv] false */

/* polymorph CustomerCreateRequest card[expiry_month] false */

/* polymorph CustomerCreateRequest card[expiry_year] false */

/* polymorph CustomerCreateRequest card[first_name] false */

/* polymorph CustomerCreateRequest card[gateway] false */

/* polymorph CustomerCreateRequest card[gateway_account_id] false */

/* polymorph CustomerCreateRequest card[ip_address] false */

/* polymorph CustomerCreateRequest card[last_name] false */

/* polymorph CustomerCreateRequest card[number] false */

/* polymorph CustomerCreateRequest card[tmp_token] false */

/* polymorph CustomerCreateRequest company false */

/* polymorph CustomerCreateRequest consolidated_invoicing false */

/* polymorph CustomerCreateRequest created_from_ip false */

/* polymorph CustomerCreateRequest email false */

/* polymorph CustomerCreateRequest entity_code false */

/* polymorph CustomerCreateRequest exempt_number false */

/* polymorph CustomerCreateRequest first_name false */

/* polymorph CustomerCreateRequest id false */

/* polymorph CustomerCreateRequest invoice_notes false */

/* polymorph CustomerCreateRequest last_name false */

/* polymorph CustomerCreateRequest locale false */

/* polymorph CustomerCreateRequest meta_data false */

/* polymorph CustomerCreateRequest net_term_days false */

/* polymorph CustomerCreateRequest payment_method[gateway] false */

/* polymorph CustomerCreateRequest payment_method[gateway_account_id] false */

/* polymorph CustomerCreateRequest payment_method[issuing_country] false */

/* polymorph CustomerCreateRequest payment_method[reference_id] false */

/* polymorph CustomerCreateRequest payment_method[tmp_token] false */

/* polymorph CustomerCreateRequest payment_method[type] false */

/* polymorph CustomerCreateRequest phone false */

/* polymorph CustomerCreateRequest preferred_currency_code false */

/* polymorph CustomerCreateRequest registered_for_gst false */

/* polymorph CustomerCreateRequest taxability false */

/* polymorph CustomerCreateRequest vat_number false */

// Validate validates this customer create request
func (m *CustomerCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoCollection(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBillingAddressValidationStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCardGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEntityCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentMethodGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentMethodType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTaxability(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customerCreateRequestTypeAutoCollectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on","off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerCreateRequestTypeAutoCollectionPropEnum = append(customerCreateRequestTypeAutoCollectionPropEnum, v)
	}
}

const (
	// CustomerCreateRequestAutoCollectionOn captures enum value "on"
	CustomerCreateRequestAutoCollectionOn string = "on"
	// CustomerCreateRequestAutoCollectionOff captures enum value "off"
	CustomerCreateRequestAutoCollectionOff string = "off"
)

// prop value enum
func (m *CustomerCreateRequest) validateAutoCollectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerCreateRequestTypeAutoCollectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomerCreateRequest) validateAutoCollection(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoCollection) { // not required
		return nil
	}

	// value enum
	if err := m.validateAutoCollectionEnum("auto_collection", "body", m.AutoCollection); err != nil {
		return err
	}

	return nil
}

var customerCreateRequestTypeBillingAddressValidationStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["not_validated","valid","partially_valid","invalid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerCreateRequestTypeBillingAddressValidationStatusPropEnum = append(customerCreateRequestTypeBillingAddressValidationStatusPropEnum, v)
	}
}

const (
	// CustomerCreateRequestBillingAddressValidationStatusNotValidated captures enum value "not_validated"
	CustomerCreateRequestBillingAddressValidationStatusNotValidated string = "not_validated"
	// CustomerCreateRequestBillingAddressValidationStatusValid captures enum value "valid"
	CustomerCreateRequestBillingAddressValidationStatusValid string = "valid"
	// CustomerCreateRequestBillingAddressValidationStatusPartiallyValid captures enum value "partially_valid"
	CustomerCreateRequestBillingAddressValidationStatusPartiallyValid string = "partially_valid"
	// CustomerCreateRequestBillingAddressValidationStatusInvalid captures enum value "invalid"
	CustomerCreateRequestBillingAddressValidationStatusInvalid string = "invalid"
)

// prop value enum
func (m *CustomerCreateRequest) validateBillingAddressValidationStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerCreateRequestTypeBillingAddressValidationStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomerCreateRequest) validateBillingAddressValidationStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingAddressValidationStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingAddressValidationStatusEnum("billing_address[validation_status]", "body", m.BillingAddressValidationStatus); err != nil {
		return err
	}

	return nil
}

var customerCreateRequestTypeCardGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chargebee","stripe","wepay","braintree","authorize_net","paypal_pro","pin","eway","eway_rapid","worldpay","balanced_payments","beanstream","bluepay","elavon","first_data_global","hdfc","migs","nmi","ogone","paymill","paypal_payflow_pro","sage_pay","tco","wirecard","amazon_payments","paypal_express_checkout","gocardless","adyen","orbital","moneris_us","moneris","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerCreateRequestTypeCardGatewayPropEnum = append(customerCreateRequestTypeCardGatewayPropEnum, v)
	}
}

const (
	// CustomerCreateRequestCardGatewayChargebee captures enum value "chargebee"
	CustomerCreateRequestCardGatewayChargebee string = "chargebee"
	// CustomerCreateRequestCardGatewayStripe captures enum value "stripe"
	CustomerCreateRequestCardGatewayStripe string = "stripe"
	// CustomerCreateRequestCardGatewayWepay captures enum value "wepay"
	CustomerCreateRequestCardGatewayWepay string = "wepay"
	// CustomerCreateRequestCardGatewayBraintree captures enum value "braintree"
	CustomerCreateRequestCardGatewayBraintree string = "braintree"
	// CustomerCreateRequestCardGatewayAuthorizeNet captures enum value "authorize_net"
	CustomerCreateRequestCardGatewayAuthorizeNet string = "authorize_net"
	// CustomerCreateRequestCardGatewayPaypalPro captures enum value "paypal_pro"
	CustomerCreateRequestCardGatewayPaypalPro string = "paypal_pro"
	// CustomerCreateRequestCardGatewayPin captures enum value "pin"
	CustomerCreateRequestCardGatewayPin string = "pin"
	// CustomerCreateRequestCardGatewayEway captures enum value "eway"
	CustomerCreateRequestCardGatewayEway string = "eway"
	// CustomerCreateRequestCardGatewayEwayRapid captures enum value "eway_rapid"
	CustomerCreateRequestCardGatewayEwayRapid string = "eway_rapid"
	// CustomerCreateRequestCardGatewayWorldpay captures enum value "worldpay"
	CustomerCreateRequestCardGatewayWorldpay string = "worldpay"
	// CustomerCreateRequestCardGatewayBalancedPayments captures enum value "balanced_payments"
	CustomerCreateRequestCardGatewayBalancedPayments string = "balanced_payments"
	// CustomerCreateRequestCardGatewayBeanstream captures enum value "beanstream"
	CustomerCreateRequestCardGatewayBeanstream string = "beanstream"
	// CustomerCreateRequestCardGatewayBluepay captures enum value "bluepay"
	CustomerCreateRequestCardGatewayBluepay string = "bluepay"
	// CustomerCreateRequestCardGatewayElavon captures enum value "elavon"
	CustomerCreateRequestCardGatewayElavon string = "elavon"
	// CustomerCreateRequestCardGatewayFirstDataGlobal captures enum value "first_data_global"
	CustomerCreateRequestCardGatewayFirstDataGlobal string = "first_data_global"
	// CustomerCreateRequestCardGatewayHdfc captures enum value "hdfc"
	CustomerCreateRequestCardGatewayHdfc string = "hdfc"
	// CustomerCreateRequestCardGatewayMigs captures enum value "migs"
	CustomerCreateRequestCardGatewayMigs string = "migs"
	// CustomerCreateRequestCardGatewayNmi captures enum value "nmi"
	CustomerCreateRequestCardGatewayNmi string = "nmi"
	// CustomerCreateRequestCardGatewayOgone captures enum value "ogone"
	CustomerCreateRequestCardGatewayOgone string = "ogone"
	// CustomerCreateRequestCardGatewayPaymill captures enum value "paymill"
	CustomerCreateRequestCardGatewayPaymill string = "paymill"
	// CustomerCreateRequestCardGatewayPaypalPayflowPro captures enum value "paypal_payflow_pro"
	CustomerCreateRequestCardGatewayPaypalPayflowPro string = "paypal_payflow_pro"
	// CustomerCreateRequestCardGatewaySagePay captures enum value "sage_pay"
	CustomerCreateRequestCardGatewaySagePay string = "sage_pay"
	// CustomerCreateRequestCardGatewayTco captures enum value "tco"
	CustomerCreateRequestCardGatewayTco string = "tco"
	// CustomerCreateRequestCardGatewayWirecard captures enum value "wirecard"
	CustomerCreateRequestCardGatewayWirecard string = "wirecard"
	// CustomerCreateRequestCardGatewayAmazonPayments captures enum value "amazon_payments"
	CustomerCreateRequestCardGatewayAmazonPayments string = "amazon_payments"
	// CustomerCreateRequestCardGatewayPaypalExpressCheckout captures enum value "paypal_express_checkout"
	CustomerCreateRequestCardGatewayPaypalExpressCheckout string = "paypal_express_checkout"
	// CustomerCreateRequestCardGatewayGocardless captures enum value "gocardless"
	CustomerCreateRequestCardGatewayGocardless string = "gocardless"
	// CustomerCreateRequestCardGatewayAdyen captures enum value "adyen"
	CustomerCreateRequestCardGatewayAdyen string = "adyen"
	// CustomerCreateRequestCardGatewayOrbital captures enum value "orbital"
	CustomerCreateRequestCardGatewayOrbital string = "orbital"
	// CustomerCreateRequestCardGatewayMonerisUs captures enum value "moneris_us"
	CustomerCreateRequestCardGatewayMonerisUs string = "moneris_us"
	// CustomerCreateRequestCardGatewayMoneris captures enum value "moneris"
	CustomerCreateRequestCardGatewayMoneris string = "moneris"
	// CustomerCreateRequestCardGatewayNotApplicable captures enum value "not_applicable"
	CustomerCreateRequestCardGatewayNotApplicable string = "not_applicable"
)

// prop value enum
func (m *CustomerCreateRequest) validateCardGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerCreateRequestTypeCardGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomerCreateRequest) validateCardGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.CardGateway) { // not required
		return nil
	}

	// value enum
	if err := m.validateCardGatewayEnum("card[gateway]", "body", m.CardGateway); err != nil {
		return err
	}

	return nil
}

var customerCreateRequestTypeEntityCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["a","b","c","d","e","f","g","h","i","j","k","l","n","p","q","r","med1","med2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerCreateRequestTypeEntityCodePropEnum = append(customerCreateRequestTypeEntityCodePropEnum, v)
	}
}

const (
	// CustomerCreateRequestEntityCodeA captures enum value "a"
	CustomerCreateRequestEntityCodeA string = "a"
	// CustomerCreateRequestEntityCodeB captures enum value "b"
	CustomerCreateRequestEntityCodeB string = "b"
	// CustomerCreateRequestEntityCodeC captures enum value "c"
	CustomerCreateRequestEntityCodeC string = "c"
	// CustomerCreateRequestEntityCodeD captures enum value "d"
	CustomerCreateRequestEntityCodeD string = "d"
	// CustomerCreateRequestEntityCodeE captures enum value "e"
	CustomerCreateRequestEntityCodeE string = "e"
	// CustomerCreateRequestEntityCodeF captures enum value "f"
	CustomerCreateRequestEntityCodeF string = "f"
	// CustomerCreateRequestEntityCodeG captures enum value "g"
	CustomerCreateRequestEntityCodeG string = "g"
	// CustomerCreateRequestEntityCodeH captures enum value "h"
	CustomerCreateRequestEntityCodeH string = "h"
	// CustomerCreateRequestEntityCodeI captures enum value "i"
	CustomerCreateRequestEntityCodeI string = "i"
	// CustomerCreateRequestEntityCodeJ captures enum value "j"
	CustomerCreateRequestEntityCodeJ string = "j"
	// CustomerCreateRequestEntityCodeK captures enum value "k"
	CustomerCreateRequestEntityCodeK string = "k"
	// CustomerCreateRequestEntityCodeL captures enum value "l"
	CustomerCreateRequestEntityCodeL string = "l"
	// CustomerCreateRequestEntityCodeN captures enum value "n"
	CustomerCreateRequestEntityCodeN string = "n"
	// CustomerCreateRequestEntityCodeP captures enum value "p"
	CustomerCreateRequestEntityCodeP string = "p"
	// CustomerCreateRequestEntityCodeQ captures enum value "q"
	CustomerCreateRequestEntityCodeQ string = "q"
	// CustomerCreateRequestEntityCodeR captures enum value "r"
	CustomerCreateRequestEntityCodeR string = "r"
	// CustomerCreateRequestEntityCodeMed1 captures enum value "med1"
	CustomerCreateRequestEntityCodeMed1 string = "med1"
	// CustomerCreateRequestEntityCodeMed2 captures enum value "med2"
	CustomerCreateRequestEntityCodeMed2 string = "med2"
)

// prop value enum
func (m *CustomerCreateRequest) validateEntityCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerCreateRequestTypeEntityCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomerCreateRequest) validateEntityCode(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityCodeEnum("entity_code", "body", m.EntityCode); err != nil {
		return err
	}

	return nil
}

var customerCreateRequestTypePaymentMethodGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chargebee","stripe","wepay","braintree","authorize_net","paypal_pro","pin","eway","eway_rapid","worldpay","balanced_payments","beanstream","bluepay","elavon","first_data_global","hdfc","migs","nmi","ogone","paymill","paypal_payflow_pro","sage_pay","tco","wirecard","amazon_payments","paypal_express_checkout","gocardless","adyen","orbital","moneris_us","moneris","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerCreateRequestTypePaymentMethodGatewayPropEnum = append(customerCreateRequestTypePaymentMethodGatewayPropEnum, v)
	}
}

const (
	// CustomerCreateRequestPaymentMethodGatewayChargebee captures enum value "chargebee"
	CustomerCreateRequestPaymentMethodGatewayChargebee string = "chargebee"
	// CustomerCreateRequestPaymentMethodGatewayStripe captures enum value "stripe"
	CustomerCreateRequestPaymentMethodGatewayStripe string = "stripe"
	// CustomerCreateRequestPaymentMethodGatewayWepay captures enum value "wepay"
	CustomerCreateRequestPaymentMethodGatewayWepay string = "wepay"
	// CustomerCreateRequestPaymentMethodGatewayBraintree captures enum value "braintree"
	CustomerCreateRequestPaymentMethodGatewayBraintree string = "braintree"
	// CustomerCreateRequestPaymentMethodGatewayAuthorizeNet captures enum value "authorize_net"
	CustomerCreateRequestPaymentMethodGatewayAuthorizeNet string = "authorize_net"
	// CustomerCreateRequestPaymentMethodGatewayPaypalPro captures enum value "paypal_pro"
	CustomerCreateRequestPaymentMethodGatewayPaypalPro string = "paypal_pro"
	// CustomerCreateRequestPaymentMethodGatewayPin captures enum value "pin"
	CustomerCreateRequestPaymentMethodGatewayPin string = "pin"
	// CustomerCreateRequestPaymentMethodGatewayEway captures enum value "eway"
	CustomerCreateRequestPaymentMethodGatewayEway string = "eway"
	// CustomerCreateRequestPaymentMethodGatewayEwayRapid captures enum value "eway_rapid"
	CustomerCreateRequestPaymentMethodGatewayEwayRapid string = "eway_rapid"
	// CustomerCreateRequestPaymentMethodGatewayWorldpay captures enum value "worldpay"
	CustomerCreateRequestPaymentMethodGatewayWorldpay string = "worldpay"
	// CustomerCreateRequestPaymentMethodGatewayBalancedPayments captures enum value "balanced_payments"
	CustomerCreateRequestPaymentMethodGatewayBalancedPayments string = "balanced_payments"
	// CustomerCreateRequestPaymentMethodGatewayBeanstream captures enum value "beanstream"
	CustomerCreateRequestPaymentMethodGatewayBeanstream string = "beanstream"
	// CustomerCreateRequestPaymentMethodGatewayBluepay captures enum value "bluepay"
	CustomerCreateRequestPaymentMethodGatewayBluepay string = "bluepay"
	// CustomerCreateRequestPaymentMethodGatewayElavon captures enum value "elavon"
	CustomerCreateRequestPaymentMethodGatewayElavon string = "elavon"
	// CustomerCreateRequestPaymentMethodGatewayFirstDataGlobal captures enum value "first_data_global"
	CustomerCreateRequestPaymentMethodGatewayFirstDataGlobal string = "first_data_global"
	// CustomerCreateRequestPaymentMethodGatewayHdfc captures enum value "hdfc"
	CustomerCreateRequestPaymentMethodGatewayHdfc string = "hdfc"
	// CustomerCreateRequestPaymentMethodGatewayMigs captures enum value "migs"
	CustomerCreateRequestPaymentMethodGatewayMigs string = "migs"
	// CustomerCreateRequestPaymentMethodGatewayNmi captures enum value "nmi"
	CustomerCreateRequestPaymentMethodGatewayNmi string = "nmi"
	// CustomerCreateRequestPaymentMethodGatewayOgone captures enum value "ogone"
	CustomerCreateRequestPaymentMethodGatewayOgone string = "ogone"
	// CustomerCreateRequestPaymentMethodGatewayPaymill captures enum value "paymill"
	CustomerCreateRequestPaymentMethodGatewayPaymill string = "paymill"
	// CustomerCreateRequestPaymentMethodGatewayPaypalPayflowPro captures enum value "paypal_payflow_pro"
	CustomerCreateRequestPaymentMethodGatewayPaypalPayflowPro string = "paypal_payflow_pro"
	// CustomerCreateRequestPaymentMethodGatewaySagePay captures enum value "sage_pay"
	CustomerCreateRequestPaymentMethodGatewaySagePay string = "sage_pay"
	// CustomerCreateRequestPaymentMethodGatewayTco captures enum value "tco"
	CustomerCreateRequestPaymentMethodGatewayTco string = "tco"
	// CustomerCreateRequestPaymentMethodGatewayWirecard captures enum value "wirecard"
	CustomerCreateRequestPaymentMethodGatewayWirecard string = "wirecard"
	// CustomerCreateRequestPaymentMethodGatewayAmazonPayments captures enum value "amazon_payments"
	CustomerCreateRequestPaymentMethodGatewayAmazonPayments string = "amazon_payments"
	// CustomerCreateRequestPaymentMethodGatewayPaypalExpressCheckout captures enum value "paypal_express_checkout"
	CustomerCreateRequestPaymentMethodGatewayPaypalExpressCheckout string = "paypal_express_checkout"
	// CustomerCreateRequestPaymentMethodGatewayGocardless captures enum value "gocardless"
	CustomerCreateRequestPaymentMethodGatewayGocardless string = "gocardless"
	// CustomerCreateRequestPaymentMethodGatewayAdyen captures enum value "adyen"
	CustomerCreateRequestPaymentMethodGatewayAdyen string = "adyen"
	// CustomerCreateRequestPaymentMethodGatewayOrbital captures enum value "orbital"
	CustomerCreateRequestPaymentMethodGatewayOrbital string = "orbital"
	// CustomerCreateRequestPaymentMethodGatewayMonerisUs captures enum value "moneris_us"
	CustomerCreateRequestPaymentMethodGatewayMonerisUs string = "moneris_us"
	// CustomerCreateRequestPaymentMethodGatewayMoneris captures enum value "moneris"
	CustomerCreateRequestPaymentMethodGatewayMoneris string = "moneris"
	// CustomerCreateRequestPaymentMethodGatewayNotApplicable captures enum value "not_applicable"
	CustomerCreateRequestPaymentMethodGatewayNotApplicable string = "not_applicable"
)

// prop value enum
func (m *CustomerCreateRequest) validatePaymentMethodGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerCreateRequestTypePaymentMethodGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomerCreateRequest) validatePaymentMethodGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethodGateway) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentMethodGatewayEnum("payment_method[gateway]", "body", m.PaymentMethodGateway); err != nil {
		return err
	}

	return nil
}

var customerCreateRequestTypePaymentMethodTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["authorization","payment","refund","payment_reversal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerCreateRequestTypePaymentMethodTypePropEnum = append(customerCreateRequestTypePaymentMethodTypePropEnum, v)
	}
}

const (
	// CustomerCreateRequestPaymentMethodTypeAuthorization captures enum value "authorization"
	CustomerCreateRequestPaymentMethodTypeAuthorization string = "authorization"
	// CustomerCreateRequestPaymentMethodTypePayment captures enum value "payment"
	CustomerCreateRequestPaymentMethodTypePayment string = "payment"
	// CustomerCreateRequestPaymentMethodTypeRefund captures enum value "refund"
	CustomerCreateRequestPaymentMethodTypeRefund string = "refund"
	// CustomerCreateRequestPaymentMethodTypePaymentReversal captures enum value "payment_reversal"
	CustomerCreateRequestPaymentMethodTypePaymentReversal string = "payment_reversal"
)

// prop value enum
func (m *CustomerCreateRequest) validatePaymentMethodTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerCreateRequestTypePaymentMethodTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomerCreateRequest) validatePaymentMethodType(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethodType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentMethodTypeEnum("payment_method[type]", "body", m.PaymentMethodType); err != nil {
		return err
	}

	return nil
}

var customerCreateRequestTypeTaxabilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["taxable","exempt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerCreateRequestTypeTaxabilityPropEnum = append(customerCreateRequestTypeTaxabilityPropEnum, v)
	}
}

const (
	// CustomerCreateRequestTaxabilityTaxable captures enum value "taxable"
	CustomerCreateRequestTaxabilityTaxable string = "taxable"
	// CustomerCreateRequestTaxabilityExempt captures enum value "exempt"
	CustomerCreateRequestTaxabilityExempt string = "exempt"
)

// prop value enum
func (m *CustomerCreateRequest) validateTaxabilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerCreateRequestTypeTaxabilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomerCreateRequest) validateTaxability(formats strfmt.Registry) error {

	if swag.IsZero(m.Taxability) { // not required
		return nil
	}

	// value enum
	if err := m.validateTaxabilityEnum("taxability", "body", m.Taxability); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerCreateRequest) UnmarshalBinary(b []byte) error {
	var res CustomerCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}