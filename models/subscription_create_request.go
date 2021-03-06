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

// SubscriptionCreateRequest subscription create request
// swagger:model SubscriptionCreateRequest
type SubscriptionCreateRequest struct {

	// affiliate token
	AffiliateToken string `json:"affiliate_token,omitempty"`

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

	// billing alignment mode
	BillingAlignmentMode string `json:"billing_alignment_mode,omitempty"`

	// billing cycles
	BillingCycles int32 `json:"billing_cycles,omitempty"`

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

	// coupon
	Coupon string `json:"coupon,omitempty"`

	// coupon ids
	CouponIds string `json:"coupon_ids,omitempty"`

	// created from ip
	CreatedFromIP string `json:"created_from_ip,omitempty"`

	// customer allow direct debit
	CustomerAllowDirectDebit bool `json:"customer[allow_direct_debit],omitempty"`

	// customer auto collection
	CustomerAutoCollection string `json:"customer[auto_collection],omitempty"`

	// customer company
	CustomerCompany string `json:"customer[company],omitempty"`

	// customer consolidated invoicing
	CustomerConsolidatedInvoicing bool `json:"customer[consolidated_invoicing],omitempty"`

	// customer email
	CustomerEmail string `json:"customer[email],omitempty"`

	// customer entity code
	CustomerEntityCode string `json:"customer[entity_code],omitempty"`

	// customer exempt number
	CustomerExemptNumber string `json:"customer[exempt_number],omitempty"`

	// customer first name
	CustomerFirstName string `json:"customer[first_name],omitempty"`

	// customer id
	CustomerID string `json:"customer[id],omitempty"`

	// customer last name
	CustomerLastName string `json:"customer[last_name],omitempty"`

	// customer locale
	CustomerLocale string `json:"customer[locale],omitempty"`

	// customer net term days
	CustomerNetTermDays int32 `json:"customer[net_term_days],omitempty"`

	// customer phone
	CustomerPhone string `json:"customer[phone],omitempty"`

	// customer registered for gst
	CustomerRegisteredForGst bool `json:"customer[registered_for_gst],omitempty"`

	// customer taxability
	CustomerTaxability string `json:"customer[taxability],omitempty"`

	// customer vat number
	CustomerVatNumber string `json:"customer[vat_number],omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// invoice immediately
	InvoiceImmediately bool `json:"invoice_immediately,omitempty"`

	// invoice notes
	InvoiceNotes string `json:"invoice_notes,omitempty"`

	// meta data
	MetaData interface{} `json:"meta_data,omitempty"`

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

	// plan id
	PlanID string `json:"plan_id,omitempty"`

	// plan quantity
	PlanQuantity int32 `json:"plan_quantity,omitempty"`

	// plan unit price
	PlanUnitPrice int32 `json:"plan_unit_price,omitempty"`

	// po number
	PoNumber string `json:"po_number,omitempty"`

	// setup fee
	SetupFee int32 `json:"setup_fee,omitempty"`

	// shipping address city
	ShippingAddressCity string `json:"shipping_address[city],omitempty"`

	// shipping address company
	ShippingAddressCompany string `json:"shipping_address[company],omitempty"`

	// shipping address country
	ShippingAddressCountry string `json:"shipping_address[country],omitempty"`

	// shipping address email
	ShippingAddressEmail string `json:"shipping_address[email],omitempty"`

	// shipping address first name
	ShippingAddressFirstName string `json:"shipping_address[first_name],omitempty"`

	// shipping address last name
	ShippingAddressLastName string `json:"shipping_address[last_name],omitempty"`

	// shipping address line1
	ShippingAddressLine1 string `json:"shipping_address[line1],omitempty"`

	// shipping address line2
	ShippingAddressLine2 string `json:"shipping_address[line2],omitempty"`

	// shipping address line3
	ShippingAddressLine3 string `json:"shipping_address[line3],omitempty"`

	// shipping address phone
	ShippingAddressPhone string `json:"shipping_address[phone],omitempty"`

	// shipping address state
	ShippingAddressState string `json:"shipping_address[state],omitempty"`

	// shipping address state code
	ShippingAddressStateCode string `json:"shipping_address[state_code],omitempty"`

	// shipping address validation status
	ShippingAddressValidationStatus string `json:"shipping_address[validation_status],omitempty"`

	// shipping address zip
	ShippingAddressZip string `json:"shipping_address[zip],omitempty"`

	// start date
	StartDate int64 `json:"start_date,omitempty"`

	// terms to charge
	TermsToCharge int32 `json:"terms_to_charge,omitempty"`

	// trial end
	TrialEnd int64 `json:"trial_end,omitempty"`
}

// Validate validates this subscription create request
func (m *SubscriptionCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoCollection(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBillingAddressValidationStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBillingAlignmentMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCardGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCustomerAutoCollection(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCustomerEntityCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCustomerTaxability(formats); err != nil {
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

	if err := m.validateShippingAddressValidationStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var subscriptionCreateRequestTypeAutoCollectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on","off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionCreateRequestTypeAutoCollectionPropEnum = append(subscriptionCreateRequestTypeAutoCollectionPropEnum, v)
	}
}

const (

	// SubscriptionCreateRequestAutoCollectionOn captures enum value "on"
	SubscriptionCreateRequestAutoCollectionOn string = "on"

	// SubscriptionCreateRequestAutoCollectionOff captures enum value "off"
	SubscriptionCreateRequestAutoCollectionOff string = "off"
)

// prop value enum
func (m *SubscriptionCreateRequest) validateAutoCollectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionCreateRequestTypeAutoCollectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionCreateRequest) validateAutoCollection(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoCollection) { // not required
		return nil
	}

	// value enum
	if err := m.validateAutoCollectionEnum("auto_collection", "body", m.AutoCollection); err != nil {
		return err
	}

	return nil
}

var subscriptionCreateRequestTypeBillingAddressValidationStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["not_validated","valid","partially_valid","invalid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionCreateRequestTypeBillingAddressValidationStatusPropEnum = append(subscriptionCreateRequestTypeBillingAddressValidationStatusPropEnum, v)
	}
}

const (

	// SubscriptionCreateRequestBillingAddressValidationStatusNotValidated captures enum value "not_validated"
	SubscriptionCreateRequestBillingAddressValidationStatusNotValidated string = "not_validated"

	// SubscriptionCreateRequestBillingAddressValidationStatusValid captures enum value "valid"
	SubscriptionCreateRequestBillingAddressValidationStatusValid string = "valid"

	// SubscriptionCreateRequestBillingAddressValidationStatusPartiallyValid captures enum value "partially_valid"
	SubscriptionCreateRequestBillingAddressValidationStatusPartiallyValid string = "partially_valid"

	// SubscriptionCreateRequestBillingAddressValidationStatusInvalid captures enum value "invalid"
	SubscriptionCreateRequestBillingAddressValidationStatusInvalid string = "invalid"
)

// prop value enum
func (m *SubscriptionCreateRequest) validateBillingAddressValidationStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionCreateRequestTypeBillingAddressValidationStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionCreateRequest) validateBillingAddressValidationStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingAddressValidationStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingAddressValidationStatusEnum("billing_address[validation_status]", "body", m.BillingAddressValidationStatus); err != nil {
		return err
	}

	return nil
}

var subscriptionCreateRequestTypeBillingAlignmentModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["immediate","delayed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionCreateRequestTypeBillingAlignmentModePropEnum = append(subscriptionCreateRequestTypeBillingAlignmentModePropEnum, v)
	}
}

const (

	// SubscriptionCreateRequestBillingAlignmentModeImmediate captures enum value "immediate"
	SubscriptionCreateRequestBillingAlignmentModeImmediate string = "immediate"

	// SubscriptionCreateRequestBillingAlignmentModeDelayed captures enum value "delayed"
	SubscriptionCreateRequestBillingAlignmentModeDelayed string = "delayed"
)

// prop value enum
func (m *SubscriptionCreateRequest) validateBillingAlignmentModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionCreateRequestTypeBillingAlignmentModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionCreateRequest) validateBillingAlignmentMode(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingAlignmentMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingAlignmentModeEnum("billing_alignment_mode", "body", m.BillingAlignmentMode); err != nil {
		return err
	}

	return nil
}

var subscriptionCreateRequestTypeCardGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chargebee","stripe","wepay","braintree","authorize_net","paypal_pro","pin","eway","eway_rapid","worldpay","balanced_payments","beanstream","bluepay","elavon","first_data_global","hdfc","migs","nmi","ogone","paymill","paypal_payflow_pro","sage_pay","tco","wirecard","amazon_payments","paypal_express_checkout","gocardless","adyen","orbital","moneris_us","moneris","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionCreateRequestTypeCardGatewayPropEnum = append(subscriptionCreateRequestTypeCardGatewayPropEnum, v)
	}
}

const (

	// SubscriptionCreateRequestCardGatewayChargebee captures enum value "chargebee"
	SubscriptionCreateRequestCardGatewayChargebee string = "chargebee"

	// SubscriptionCreateRequestCardGatewayStripe captures enum value "stripe"
	SubscriptionCreateRequestCardGatewayStripe string = "stripe"

	// SubscriptionCreateRequestCardGatewayWepay captures enum value "wepay"
	SubscriptionCreateRequestCardGatewayWepay string = "wepay"

	// SubscriptionCreateRequestCardGatewayBraintree captures enum value "braintree"
	SubscriptionCreateRequestCardGatewayBraintree string = "braintree"

	// SubscriptionCreateRequestCardGatewayAuthorizeNet captures enum value "authorize_net"
	SubscriptionCreateRequestCardGatewayAuthorizeNet string = "authorize_net"

	// SubscriptionCreateRequestCardGatewayPaypalPro captures enum value "paypal_pro"
	SubscriptionCreateRequestCardGatewayPaypalPro string = "paypal_pro"

	// SubscriptionCreateRequestCardGatewayPin captures enum value "pin"
	SubscriptionCreateRequestCardGatewayPin string = "pin"

	// SubscriptionCreateRequestCardGatewayEway captures enum value "eway"
	SubscriptionCreateRequestCardGatewayEway string = "eway"

	// SubscriptionCreateRequestCardGatewayEwayRapid captures enum value "eway_rapid"
	SubscriptionCreateRequestCardGatewayEwayRapid string = "eway_rapid"

	// SubscriptionCreateRequestCardGatewayWorldpay captures enum value "worldpay"
	SubscriptionCreateRequestCardGatewayWorldpay string = "worldpay"

	// SubscriptionCreateRequestCardGatewayBalancedPayments captures enum value "balanced_payments"
	SubscriptionCreateRequestCardGatewayBalancedPayments string = "balanced_payments"

	// SubscriptionCreateRequestCardGatewayBeanstream captures enum value "beanstream"
	SubscriptionCreateRequestCardGatewayBeanstream string = "beanstream"

	// SubscriptionCreateRequestCardGatewayBluepay captures enum value "bluepay"
	SubscriptionCreateRequestCardGatewayBluepay string = "bluepay"

	// SubscriptionCreateRequestCardGatewayElavon captures enum value "elavon"
	SubscriptionCreateRequestCardGatewayElavon string = "elavon"

	// SubscriptionCreateRequestCardGatewayFirstDataGlobal captures enum value "first_data_global"
	SubscriptionCreateRequestCardGatewayFirstDataGlobal string = "first_data_global"

	// SubscriptionCreateRequestCardGatewayHdfc captures enum value "hdfc"
	SubscriptionCreateRequestCardGatewayHdfc string = "hdfc"

	// SubscriptionCreateRequestCardGatewayMigs captures enum value "migs"
	SubscriptionCreateRequestCardGatewayMigs string = "migs"

	// SubscriptionCreateRequestCardGatewayNmi captures enum value "nmi"
	SubscriptionCreateRequestCardGatewayNmi string = "nmi"

	// SubscriptionCreateRequestCardGatewayOgone captures enum value "ogone"
	SubscriptionCreateRequestCardGatewayOgone string = "ogone"

	// SubscriptionCreateRequestCardGatewayPaymill captures enum value "paymill"
	SubscriptionCreateRequestCardGatewayPaymill string = "paymill"

	// SubscriptionCreateRequestCardGatewayPaypalPayflowPro captures enum value "paypal_payflow_pro"
	SubscriptionCreateRequestCardGatewayPaypalPayflowPro string = "paypal_payflow_pro"

	// SubscriptionCreateRequestCardGatewaySagePay captures enum value "sage_pay"
	SubscriptionCreateRequestCardGatewaySagePay string = "sage_pay"

	// SubscriptionCreateRequestCardGatewayTco captures enum value "tco"
	SubscriptionCreateRequestCardGatewayTco string = "tco"

	// SubscriptionCreateRequestCardGatewayWirecard captures enum value "wirecard"
	SubscriptionCreateRequestCardGatewayWirecard string = "wirecard"

	// SubscriptionCreateRequestCardGatewayAmazonPayments captures enum value "amazon_payments"
	SubscriptionCreateRequestCardGatewayAmazonPayments string = "amazon_payments"

	// SubscriptionCreateRequestCardGatewayPaypalExpressCheckout captures enum value "paypal_express_checkout"
	SubscriptionCreateRequestCardGatewayPaypalExpressCheckout string = "paypal_express_checkout"

	// SubscriptionCreateRequestCardGatewayGocardless captures enum value "gocardless"
	SubscriptionCreateRequestCardGatewayGocardless string = "gocardless"

	// SubscriptionCreateRequestCardGatewayAdyen captures enum value "adyen"
	SubscriptionCreateRequestCardGatewayAdyen string = "adyen"

	// SubscriptionCreateRequestCardGatewayOrbital captures enum value "orbital"
	SubscriptionCreateRequestCardGatewayOrbital string = "orbital"

	// SubscriptionCreateRequestCardGatewayMonerisUs captures enum value "moneris_us"
	SubscriptionCreateRequestCardGatewayMonerisUs string = "moneris_us"

	// SubscriptionCreateRequestCardGatewayMoneris captures enum value "moneris"
	SubscriptionCreateRequestCardGatewayMoneris string = "moneris"

	// SubscriptionCreateRequestCardGatewayNotApplicable captures enum value "not_applicable"
	SubscriptionCreateRequestCardGatewayNotApplicable string = "not_applicable"
)

// prop value enum
func (m *SubscriptionCreateRequest) validateCardGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionCreateRequestTypeCardGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionCreateRequest) validateCardGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.CardGateway) { // not required
		return nil
	}

	// value enum
	if err := m.validateCardGatewayEnum("card[gateway]", "body", m.CardGateway); err != nil {
		return err
	}

	return nil
}

var subscriptionCreateRequestTypeCustomerAutoCollectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on","off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionCreateRequestTypeCustomerAutoCollectionPropEnum = append(subscriptionCreateRequestTypeCustomerAutoCollectionPropEnum, v)
	}
}

const (

	// SubscriptionCreateRequestCustomerAutoCollectionOn captures enum value "on"
	SubscriptionCreateRequestCustomerAutoCollectionOn string = "on"

	// SubscriptionCreateRequestCustomerAutoCollectionOff captures enum value "off"
	SubscriptionCreateRequestCustomerAutoCollectionOff string = "off"
)

// prop value enum
func (m *SubscriptionCreateRequest) validateCustomerAutoCollectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionCreateRequestTypeCustomerAutoCollectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionCreateRequest) validateCustomerAutoCollection(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerAutoCollection) { // not required
		return nil
	}

	// value enum
	if err := m.validateCustomerAutoCollectionEnum("customer[auto_collection]", "body", m.CustomerAutoCollection); err != nil {
		return err
	}

	return nil
}

var subscriptionCreateRequestTypeCustomerEntityCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["a","b","c","d","e","f","g","h","i","j","k","l","n","p","q","r","med1","med2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionCreateRequestTypeCustomerEntityCodePropEnum = append(subscriptionCreateRequestTypeCustomerEntityCodePropEnum, v)
	}
}

const (

	// SubscriptionCreateRequestCustomerEntityCodeA captures enum value "a"
	SubscriptionCreateRequestCustomerEntityCodeA string = "a"

	// SubscriptionCreateRequestCustomerEntityCodeB captures enum value "b"
	SubscriptionCreateRequestCustomerEntityCodeB string = "b"

	// SubscriptionCreateRequestCustomerEntityCodeC captures enum value "c"
	SubscriptionCreateRequestCustomerEntityCodeC string = "c"

	// SubscriptionCreateRequestCustomerEntityCodeD captures enum value "d"
	SubscriptionCreateRequestCustomerEntityCodeD string = "d"

	// SubscriptionCreateRequestCustomerEntityCodeE captures enum value "e"
	SubscriptionCreateRequestCustomerEntityCodeE string = "e"

	// SubscriptionCreateRequestCustomerEntityCodeF captures enum value "f"
	SubscriptionCreateRequestCustomerEntityCodeF string = "f"

	// SubscriptionCreateRequestCustomerEntityCodeG captures enum value "g"
	SubscriptionCreateRequestCustomerEntityCodeG string = "g"

	// SubscriptionCreateRequestCustomerEntityCodeH captures enum value "h"
	SubscriptionCreateRequestCustomerEntityCodeH string = "h"

	// SubscriptionCreateRequestCustomerEntityCodeI captures enum value "i"
	SubscriptionCreateRequestCustomerEntityCodeI string = "i"

	// SubscriptionCreateRequestCustomerEntityCodeJ captures enum value "j"
	SubscriptionCreateRequestCustomerEntityCodeJ string = "j"

	// SubscriptionCreateRequestCustomerEntityCodeK captures enum value "k"
	SubscriptionCreateRequestCustomerEntityCodeK string = "k"

	// SubscriptionCreateRequestCustomerEntityCodeL captures enum value "l"
	SubscriptionCreateRequestCustomerEntityCodeL string = "l"

	// SubscriptionCreateRequestCustomerEntityCodeN captures enum value "n"
	SubscriptionCreateRequestCustomerEntityCodeN string = "n"

	// SubscriptionCreateRequestCustomerEntityCodeP captures enum value "p"
	SubscriptionCreateRequestCustomerEntityCodeP string = "p"

	// SubscriptionCreateRequestCustomerEntityCodeQ captures enum value "q"
	SubscriptionCreateRequestCustomerEntityCodeQ string = "q"

	// SubscriptionCreateRequestCustomerEntityCodeR captures enum value "r"
	SubscriptionCreateRequestCustomerEntityCodeR string = "r"

	// SubscriptionCreateRequestCustomerEntityCodeMed1 captures enum value "med1"
	SubscriptionCreateRequestCustomerEntityCodeMed1 string = "med1"

	// SubscriptionCreateRequestCustomerEntityCodeMed2 captures enum value "med2"
	SubscriptionCreateRequestCustomerEntityCodeMed2 string = "med2"
)

// prop value enum
func (m *SubscriptionCreateRequest) validateCustomerEntityCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionCreateRequestTypeCustomerEntityCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionCreateRequest) validateCustomerEntityCode(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerEntityCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateCustomerEntityCodeEnum("customer[entity_code]", "body", m.CustomerEntityCode); err != nil {
		return err
	}

	return nil
}

var subscriptionCreateRequestTypeCustomerTaxabilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["taxable","exempt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionCreateRequestTypeCustomerTaxabilityPropEnum = append(subscriptionCreateRequestTypeCustomerTaxabilityPropEnum, v)
	}
}

const (

	// SubscriptionCreateRequestCustomerTaxabilityTaxable captures enum value "taxable"
	SubscriptionCreateRequestCustomerTaxabilityTaxable string = "taxable"

	// SubscriptionCreateRequestCustomerTaxabilityExempt captures enum value "exempt"
	SubscriptionCreateRequestCustomerTaxabilityExempt string = "exempt"
)

// prop value enum
func (m *SubscriptionCreateRequest) validateCustomerTaxabilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionCreateRequestTypeCustomerTaxabilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionCreateRequest) validateCustomerTaxability(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerTaxability) { // not required
		return nil
	}

	// value enum
	if err := m.validateCustomerTaxabilityEnum("customer[taxability]", "body", m.CustomerTaxability); err != nil {
		return err
	}

	return nil
}

var subscriptionCreateRequestTypePaymentMethodGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chargebee","stripe","wepay","braintree","authorize_net","paypal_pro","pin","eway","eway_rapid","worldpay","balanced_payments","beanstream","bluepay","elavon","first_data_global","hdfc","migs","nmi","ogone","paymill","paypal_payflow_pro","sage_pay","tco","wirecard","amazon_payments","paypal_express_checkout","gocardless","adyen","orbital","moneris_us","moneris","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionCreateRequestTypePaymentMethodGatewayPropEnum = append(subscriptionCreateRequestTypePaymentMethodGatewayPropEnum, v)
	}
}

const (

	// SubscriptionCreateRequestPaymentMethodGatewayChargebee captures enum value "chargebee"
	SubscriptionCreateRequestPaymentMethodGatewayChargebee string = "chargebee"

	// SubscriptionCreateRequestPaymentMethodGatewayStripe captures enum value "stripe"
	SubscriptionCreateRequestPaymentMethodGatewayStripe string = "stripe"

	// SubscriptionCreateRequestPaymentMethodGatewayWepay captures enum value "wepay"
	SubscriptionCreateRequestPaymentMethodGatewayWepay string = "wepay"

	// SubscriptionCreateRequestPaymentMethodGatewayBraintree captures enum value "braintree"
	SubscriptionCreateRequestPaymentMethodGatewayBraintree string = "braintree"

	// SubscriptionCreateRequestPaymentMethodGatewayAuthorizeNet captures enum value "authorize_net"
	SubscriptionCreateRequestPaymentMethodGatewayAuthorizeNet string = "authorize_net"

	// SubscriptionCreateRequestPaymentMethodGatewayPaypalPro captures enum value "paypal_pro"
	SubscriptionCreateRequestPaymentMethodGatewayPaypalPro string = "paypal_pro"

	// SubscriptionCreateRequestPaymentMethodGatewayPin captures enum value "pin"
	SubscriptionCreateRequestPaymentMethodGatewayPin string = "pin"

	// SubscriptionCreateRequestPaymentMethodGatewayEway captures enum value "eway"
	SubscriptionCreateRequestPaymentMethodGatewayEway string = "eway"

	// SubscriptionCreateRequestPaymentMethodGatewayEwayRapid captures enum value "eway_rapid"
	SubscriptionCreateRequestPaymentMethodGatewayEwayRapid string = "eway_rapid"

	// SubscriptionCreateRequestPaymentMethodGatewayWorldpay captures enum value "worldpay"
	SubscriptionCreateRequestPaymentMethodGatewayWorldpay string = "worldpay"

	// SubscriptionCreateRequestPaymentMethodGatewayBalancedPayments captures enum value "balanced_payments"
	SubscriptionCreateRequestPaymentMethodGatewayBalancedPayments string = "balanced_payments"

	// SubscriptionCreateRequestPaymentMethodGatewayBeanstream captures enum value "beanstream"
	SubscriptionCreateRequestPaymentMethodGatewayBeanstream string = "beanstream"

	// SubscriptionCreateRequestPaymentMethodGatewayBluepay captures enum value "bluepay"
	SubscriptionCreateRequestPaymentMethodGatewayBluepay string = "bluepay"

	// SubscriptionCreateRequestPaymentMethodGatewayElavon captures enum value "elavon"
	SubscriptionCreateRequestPaymentMethodGatewayElavon string = "elavon"

	// SubscriptionCreateRequestPaymentMethodGatewayFirstDataGlobal captures enum value "first_data_global"
	SubscriptionCreateRequestPaymentMethodGatewayFirstDataGlobal string = "first_data_global"

	// SubscriptionCreateRequestPaymentMethodGatewayHdfc captures enum value "hdfc"
	SubscriptionCreateRequestPaymentMethodGatewayHdfc string = "hdfc"

	// SubscriptionCreateRequestPaymentMethodGatewayMigs captures enum value "migs"
	SubscriptionCreateRequestPaymentMethodGatewayMigs string = "migs"

	// SubscriptionCreateRequestPaymentMethodGatewayNmi captures enum value "nmi"
	SubscriptionCreateRequestPaymentMethodGatewayNmi string = "nmi"

	// SubscriptionCreateRequestPaymentMethodGatewayOgone captures enum value "ogone"
	SubscriptionCreateRequestPaymentMethodGatewayOgone string = "ogone"

	// SubscriptionCreateRequestPaymentMethodGatewayPaymill captures enum value "paymill"
	SubscriptionCreateRequestPaymentMethodGatewayPaymill string = "paymill"

	// SubscriptionCreateRequestPaymentMethodGatewayPaypalPayflowPro captures enum value "paypal_payflow_pro"
	SubscriptionCreateRequestPaymentMethodGatewayPaypalPayflowPro string = "paypal_payflow_pro"

	// SubscriptionCreateRequestPaymentMethodGatewaySagePay captures enum value "sage_pay"
	SubscriptionCreateRequestPaymentMethodGatewaySagePay string = "sage_pay"

	// SubscriptionCreateRequestPaymentMethodGatewayTco captures enum value "tco"
	SubscriptionCreateRequestPaymentMethodGatewayTco string = "tco"

	// SubscriptionCreateRequestPaymentMethodGatewayWirecard captures enum value "wirecard"
	SubscriptionCreateRequestPaymentMethodGatewayWirecard string = "wirecard"

	// SubscriptionCreateRequestPaymentMethodGatewayAmazonPayments captures enum value "amazon_payments"
	SubscriptionCreateRequestPaymentMethodGatewayAmazonPayments string = "amazon_payments"

	// SubscriptionCreateRequestPaymentMethodGatewayPaypalExpressCheckout captures enum value "paypal_express_checkout"
	SubscriptionCreateRequestPaymentMethodGatewayPaypalExpressCheckout string = "paypal_express_checkout"

	// SubscriptionCreateRequestPaymentMethodGatewayGocardless captures enum value "gocardless"
	SubscriptionCreateRequestPaymentMethodGatewayGocardless string = "gocardless"

	// SubscriptionCreateRequestPaymentMethodGatewayAdyen captures enum value "adyen"
	SubscriptionCreateRequestPaymentMethodGatewayAdyen string = "adyen"

	// SubscriptionCreateRequestPaymentMethodGatewayOrbital captures enum value "orbital"
	SubscriptionCreateRequestPaymentMethodGatewayOrbital string = "orbital"

	// SubscriptionCreateRequestPaymentMethodGatewayMonerisUs captures enum value "moneris_us"
	SubscriptionCreateRequestPaymentMethodGatewayMonerisUs string = "moneris_us"

	// SubscriptionCreateRequestPaymentMethodGatewayMoneris captures enum value "moneris"
	SubscriptionCreateRequestPaymentMethodGatewayMoneris string = "moneris"

	// SubscriptionCreateRequestPaymentMethodGatewayNotApplicable captures enum value "not_applicable"
	SubscriptionCreateRequestPaymentMethodGatewayNotApplicable string = "not_applicable"
)

// prop value enum
func (m *SubscriptionCreateRequest) validatePaymentMethodGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionCreateRequestTypePaymentMethodGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionCreateRequest) validatePaymentMethodGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethodGateway) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentMethodGatewayEnum("payment_method[gateway]", "body", m.PaymentMethodGateway); err != nil {
		return err
	}

	return nil
}

var subscriptionCreateRequestTypePaymentMethodTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["card","paypal_express_checkout","amazon_payments","direct_debit","generic","alipay","unionpay","apple_pay"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionCreateRequestTypePaymentMethodTypePropEnum = append(subscriptionCreateRequestTypePaymentMethodTypePropEnum, v)
	}
}

const (

	// SubscriptionCreateRequestPaymentMethodTypeCard captures enum value "card"
	SubscriptionCreateRequestPaymentMethodTypeCard string = "card"

	// SubscriptionCreateRequestPaymentMethodTypePaypalExpressCheckout captures enum value "paypal_express_checkout"
	SubscriptionCreateRequestPaymentMethodTypePaypalExpressCheckout string = "paypal_express_checkout"

	// SubscriptionCreateRequestPaymentMethodTypeAmazonPayments captures enum value "amazon_payments"
	SubscriptionCreateRequestPaymentMethodTypeAmazonPayments string = "amazon_payments"

	// SubscriptionCreateRequestPaymentMethodTypeDirectDebit captures enum value "direct_debit"
	SubscriptionCreateRequestPaymentMethodTypeDirectDebit string = "direct_debit"

	// SubscriptionCreateRequestPaymentMethodTypeGeneric captures enum value "generic"
	SubscriptionCreateRequestPaymentMethodTypeGeneric string = "generic"

	// SubscriptionCreateRequestPaymentMethodTypeAlipay captures enum value "alipay"
	SubscriptionCreateRequestPaymentMethodTypeAlipay string = "alipay"

	// SubscriptionCreateRequestPaymentMethodTypeUnionpay captures enum value "unionpay"
	SubscriptionCreateRequestPaymentMethodTypeUnionpay string = "unionpay"

	// SubscriptionCreateRequestPaymentMethodTypeApplePay captures enum value "apple_pay"
	SubscriptionCreateRequestPaymentMethodTypeApplePay string = "apple_pay"
)

// prop value enum
func (m *SubscriptionCreateRequest) validatePaymentMethodTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionCreateRequestTypePaymentMethodTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionCreateRequest) validatePaymentMethodType(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethodType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentMethodTypeEnum("payment_method[type]", "body", m.PaymentMethodType); err != nil {
		return err
	}

	return nil
}

var subscriptionCreateRequestTypeShippingAddressValidationStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["not_validated","valid","partially_valid","invalid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionCreateRequestTypeShippingAddressValidationStatusPropEnum = append(subscriptionCreateRequestTypeShippingAddressValidationStatusPropEnum, v)
	}
}

const (

	// SubscriptionCreateRequestShippingAddressValidationStatusNotValidated captures enum value "not_validated"
	SubscriptionCreateRequestShippingAddressValidationStatusNotValidated string = "not_validated"

	// SubscriptionCreateRequestShippingAddressValidationStatusValid captures enum value "valid"
	SubscriptionCreateRequestShippingAddressValidationStatusValid string = "valid"

	// SubscriptionCreateRequestShippingAddressValidationStatusPartiallyValid captures enum value "partially_valid"
	SubscriptionCreateRequestShippingAddressValidationStatusPartiallyValid string = "partially_valid"

	// SubscriptionCreateRequestShippingAddressValidationStatusInvalid captures enum value "invalid"
	SubscriptionCreateRequestShippingAddressValidationStatusInvalid string = "invalid"
)

// prop value enum
func (m *SubscriptionCreateRequest) validateShippingAddressValidationStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionCreateRequestTypeShippingAddressValidationStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionCreateRequest) validateShippingAddressValidationStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ShippingAddressValidationStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateShippingAddressValidationStatusEnum("shipping_address[validation_status]", "body", m.ShippingAddressValidationStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionCreateRequest) UnmarshalBinary(b []byte) error {
	var res SubscriptionCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
