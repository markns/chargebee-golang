// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Transaction transaction
// swagger:model Transaction
type Transaction struct {

	// amount
	Amount int32 `json:"amount,omitempty"`

	// amount unused
	AmountUnused int32 `json:"amount_unused,omitempty"`

	// currency code
	CurrencyCode string `json:"currency_code,omitempty"`

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// date
	Date int64 `json:"date,omitempty"`

	// deleted
	Deleted bool `json:"deleted,omitempty"`

	// error code
	ErrorCode string `json:"error_code,omitempty"`

	// error text
	ErrorText string `json:"error_text,omitempty"`

	// fraud flag
	FraudFlag string `json:"fraud_flag,omitempty"`

	// fraud reason
	FraudReason string `json:"fraud_reason,omitempty"`

	// gateway
	Gateway string `json:"gateway,omitempty"`

	// gateway account id
	GatewayAccountID string `json:"gateway_account_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// id at gateway
	IDAtGateway string `json:"id_at_gateway,omitempty"`

	// linked credit notes
	LinkedCreditNotes []*LinkedCreditNote `json:"linked_credit_notes"`

	// linked invoices
	LinkedInvoices []*LinkedInvoice `json:"linked_invoices"`

	// linked refunds
	LinkedRefunds []*LinkedRefund `json:"linked_refunds"`

	// masked card number
	MaskedCardNumber string `json:"masked_card_number,omitempty"`

	// payment method
	PaymentMethod *PaymentMethod `json:"payment_method,omitempty"`

	// payment source id
	PaymentSourceID string `json:"payment_source_id,omitempty"`

	// reference number
	ReferenceNumber string `json:"reference_number,omitempty"`

	// reference transaction id
	ReferenceTransactionID string `json:"reference_transaction_id,omitempty"`

	// refunded txn id
	RefundedTxnID string `json:"refunded_txn_id,omitempty"`

	// resource version
	ResourceVersion int64 `json:"resource_version,omitempty"`

	// reversal transaction id
	ReversalTransactionID string `json:"reversal_transaction_id,omitempty"`

	// settled at
	SettledAt int64 `json:"settled_at,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// subscription id
	SubscriptionID string `json:"subscription_id,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated at
	UpdatedAt int64 `json:"updated_at,omitempty"`

	// voided at
	VoidedAt int64 `json:"voided_at,omitempty"`
}

// Validate validates this transaction
func (m *Transaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFraudFlag(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLinkedCreditNotes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLinkedInvoices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLinkedRefunds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentMethod(formats); err != nil {
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

var transactionTypeFraudFlagPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["safe","suspicious","fraudulent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionTypeFraudFlagPropEnum = append(transactionTypeFraudFlagPropEnum, v)
	}
}

const (

	// TransactionFraudFlagSafe captures enum value "safe"
	TransactionFraudFlagSafe string = "safe"

	// TransactionFraudFlagSuspicious captures enum value "suspicious"
	TransactionFraudFlagSuspicious string = "suspicious"

	// TransactionFraudFlagFraudulent captures enum value "fraudulent"
	TransactionFraudFlagFraudulent string = "fraudulent"
)

// prop value enum
func (m *Transaction) validateFraudFlagEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, transactionTypeFraudFlagPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Transaction) validateFraudFlag(formats strfmt.Registry) error {

	if swag.IsZero(m.FraudFlag) { // not required
		return nil
	}

	// value enum
	if err := m.validateFraudFlagEnum("fraud_flag", "body", m.FraudFlag); err != nil {
		return err
	}

	return nil
}

var transactionTypeGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chargebee","stripe","wepay","braintree","authorize_net","paypal_pro","pin","eway","eway_rapid","worldpay","balanced_payments","beanstream","bluepay","elavon","first_data_global","hdfc","migs","nmi","ogone","paymill","paypal_payflow_pro","sage_pay","tco","wirecard","amazon_payments","paypal_express_checkout","gocardless","adyen","orbital","moneris_us","moneris","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionTypeGatewayPropEnum = append(transactionTypeGatewayPropEnum, v)
	}
}

const (

	// TransactionGatewayChargebee captures enum value "chargebee"
	TransactionGatewayChargebee string = "chargebee"

	// TransactionGatewayStripe captures enum value "stripe"
	TransactionGatewayStripe string = "stripe"

	// TransactionGatewayWepay captures enum value "wepay"
	TransactionGatewayWepay string = "wepay"

	// TransactionGatewayBraintree captures enum value "braintree"
	TransactionGatewayBraintree string = "braintree"

	// TransactionGatewayAuthorizeNet captures enum value "authorize_net"
	TransactionGatewayAuthorizeNet string = "authorize_net"

	// TransactionGatewayPaypalPro captures enum value "paypal_pro"
	TransactionGatewayPaypalPro string = "paypal_pro"

	// TransactionGatewayPin captures enum value "pin"
	TransactionGatewayPin string = "pin"

	// TransactionGatewayEway captures enum value "eway"
	TransactionGatewayEway string = "eway"

	// TransactionGatewayEwayRapid captures enum value "eway_rapid"
	TransactionGatewayEwayRapid string = "eway_rapid"

	// TransactionGatewayWorldpay captures enum value "worldpay"
	TransactionGatewayWorldpay string = "worldpay"

	// TransactionGatewayBalancedPayments captures enum value "balanced_payments"
	TransactionGatewayBalancedPayments string = "balanced_payments"

	// TransactionGatewayBeanstream captures enum value "beanstream"
	TransactionGatewayBeanstream string = "beanstream"

	// TransactionGatewayBluepay captures enum value "bluepay"
	TransactionGatewayBluepay string = "bluepay"

	// TransactionGatewayElavon captures enum value "elavon"
	TransactionGatewayElavon string = "elavon"

	// TransactionGatewayFirstDataGlobal captures enum value "first_data_global"
	TransactionGatewayFirstDataGlobal string = "first_data_global"

	// TransactionGatewayHdfc captures enum value "hdfc"
	TransactionGatewayHdfc string = "hdfc"

	// TransactionGatewayMigs captures enum value "migs"
	TransactionGatewayMigs string = "migs"

	// TransactionGatewayNmi captures enum value "nmi"
	TransactionGatewayNmi string = "nmi"

	// TransactionGatewayOgone captures enum value "ogone"
	TransactionGatewayOgone string = "ogone"

	// TransactionGatewayPaymill captures enum value "paymill"
	TransactionGatewayPaymill string = "paymill"

	// TransactionGatewayPaypalPayflowPro captures enum value "paypal_payflow_pro"
	TransactionGatewayPaypalPayflowPro string = "paypal_payflow_pro"

	// TransactionGatewaySagePay captures enum value "sage_pay"
	TransactionGatewaySagePay string = "sage_pay"

	// TransactionGatewayTco captures enum value "tco"
	TransactionGatewayTco string = "tco"

	// TransactionGatewayWirecard captures enum value "wirecard"
	TransactionGatewayWirecard string = "wirecard"

	// TransactionGatewayAmazonPayments captures enum value "amazon_payments"
	TransactionGatewayAmazonPayments string = "amazon_payments"

	// TransactionGatewayPaypalExpressCheckout captures enum value "paypal_express_checkout"
	TransactionGatewayPaypalExpressCheckout string = "paypal_express_checkout"

	// TransactionGatewayGocardless captures enum value "gocardless"
	TransactionGatewayGocardless string = "gocardless"

	// TransactionGatewayAdyen captures enum value "adyen"
	TransactionGatewayAdyen string = "adyen"

	// TransactionGatewayOrbital captures enum value "orbital"
	TransactionGatewayOrbital string = "orbital"

	// TransactionGatewayMonerisUs captures enum value "moneris_us"
	TransactionGatewayMonerisUs string = "moneris_us"

	// TransactionGatewayMoneris captures enum value "moneris"
	TransactionGatewayMoneris string = "moneris"

	// TransactionGatewayNotApplicable captures enum value "not_applicable"
	TransactionGatewayNotApplicable string = "not_applicable"
)

// prop value enum
func (m *Transaction) validateGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, transactionTypeGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Transaction) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	// value enum
	if err := m.validateGatewayEnum("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateLinkedCreditNotes(formats strfmt.Registry) error {

	if swag.IsZero(m.LinkedCreditNotes) { // not required
		return nil
	}

	for i := 0; i < len(m.LinkedCreditNotes); i++ {

		if swag.IsZero(m.LinkedCreditNotes[i]) { // not required
			continue
		}

		if m.LinkedCreditNotes[i] != nil {

			if err := m.LinkedCreditNotes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("linked_credit_notes" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *Transaction) validateLinkedInvoices(formats strfmt.Registry) error {

	if swag.IsZero(m.LinkedInvoices) { // not required
		return nil
	}

	for i := 0; i < len(m.LinkedInvoices); i++ {

		if swag.IsZero(m.LinkedInvoices[i]) { // not required
			continue
		}

		if m.LinkedInvoices[i] != nil {

			if err := m.LinkedInvoices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("linked_invoices" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *Transaction) validateLinkedRefunds(formats strfmt.Registry) error {

	if swag.IsZero(m.LinkedRefunds) { // not required
		return nil
	}

	for i := 0; i < len(m.LinkedRefunds); i++ {

		if swag.IsZero(m.LinkedRefunds[i]) { // not required
			continue
		}

		if m.LinkedRefunds[i] != nil {

			if err := m.LinkedRefunds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("linked_refunds" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *Transaction) validatePaymentMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethod) { // not required
		return nil
	}

	if m.PaymentMethod != nil {

		if err := m.PaymentMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_method")
			}
			return err
		}

	}

	return nil
}

var transactionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["future","in_trial","active","non_renewing","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionTypeStatusPropEnum = append(transactionTypeStatusPropEnum, v)
	}
}

const (

	// TransactionStatusFuture captures enum value "future"
	TransactionStatusFuture string = "future"

	// TransactionStatusInTrial captures enum value "in_trial"
	TransactionStatusInTrial string = "in_trial"

	// TransactionStatusActive captures enum value "active"
	TransactionStatusActive string = "active"

	// TransactionStatusNonRenewing captures enum value "non_renewing"
	TransactionStatusNonRenewing string = "non_renewing"

	// TransactionStatusCancelled captures enum value "cancelled"
	TransactionStatusCancelled string = "cancelled"
)

// prop value enum
func (m *Transaction) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, transactionTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Transaction) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var transactionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["card","paypal_express_checkout","amazon_payments","direct_debit","generic","alipay","unionpay","apple_pay"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionTypeTypePropEnum = append(transactionTypeTypePropEnum, v)
	}
}

const (

	// TransactionTypeCard captures enum value "card"
	TransactionTypeCard string = "card"

	// TransactionTypePaypalExpressCheckout captures enum value "paypal_express_checkout"
	TransactionTypePaypalExpressCheckout string = "paypal_express_checkout"

	// TransactionTypeAmazonPayments captures enum value "amazon_payments"
	TransactionTypeAmazonPayments string = "amazon_payments"

	// TransactionTypeDirectDebit captures enum value "direct_debit"
	TransactionTypeDirectDebit string = "direct_debit"

	// TransactionTypeGeneric captures enum value "generic"
	TransactionTypeGeneric string = "generic"

	// TransactionTypeAlipay captures enum value "alipay"
	TransactionTypeAlipay string = "alipay"

	// TransactionTypeUnionpay captures enum value "unionpay"
	TransactionTypeUnionpay string = "unionpay"

	// TransactionTypeApplePay captures enum value "apple_pay"
	TransactionTypeApplePay string = "apple_pay"
)

// prop value enum
func (m *Transaction) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, transactionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Transaction) validateType(formats strfmt.Registry) error {

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
func (m *Transaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transaction) UnmarshalBinary(b []byte) error {
	var res Transaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
