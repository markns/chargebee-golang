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

// Customer customer
// swagger:model Customer
type Customer struct {

	// allow direct debit
	AllowDirectDebit bool `json:"allow_direct_debit,omitempty"`

	// auto collection
	AutoCollection string `json:"auto_collection,omitempty"`

	// backup payment source id
	BackupPaymentSourceID string `json:"backup_payment_source_id,omitempty"`

	// balances
	Balances []*Balance `json:"balances"`

	// billing address
	BillingAddress *BillingAddress `json:"billing_address,omitempty"`

	// billing date
	BillingDate int32 `json:"billing_date,omitempty"`

	// billing date mode
	BillingDateMode string `json:"billing_date_mode,omitempty"`

	// billing day of week
	BillingDayOfWeek string `json:"billing_day_of_week,omitempty"`

	// billing day of week mode
	BillingDayOfWeekMode string `json:"billing_day_of_week_mode,omitempty"`

	// company
	Company string `json:"company,omitempty"`

	// consolidated invoicing
	ConsolidatedInvoicing bool `json:"consolidated_invoicing,omitempty"`

	// contacts
	Contacts []*Contact `json:"contacts"`

	// created at
	CreatedAt int64 `json:"created_at,omitempty"`

	// created from ip
	CreatedFromIP string `json:"created_from_ip,omitempty"`

	// deleted
	Deleted bool `json:"deleted,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// entity code
	EntityCode string `json:"entity_code,omitempty"`

	// excess payments
	ExcessPayments int32 `json:"excess_payments,omitempty"`

	// exempt number
	ExemptNumber string `json:"exempt_number,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// fraud flag
	FraudFlag string `json:"fraud_flag,omitempty"`

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

	// payment method
	PaymentMethod *PaymentMethod `json:"payment_method,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// preferred currency code
	PreferredCurrencyCode string `json:"preferred_currency_code,omitempty"`

	// primary payment source id
	PrimaryPaymentSourceID string `json:"primary_payment_source_id,omitempty"`

	// promotional credits
	PromotionalCredits int32 `json:"promotional_credits,omitempty"`

	// referral urls
	ReferralUrls []*ReferralURL `json:"referral_urls"`

	// refundable credits
	RefundableCredits int32 `json:"refundable_credits,omitempty"`

	// registered for gst
	RegisteredForGst bool `json:"registered_for_gst,omitempty"`

	// resource version
	ResourceVersion int64 `json:"resource_version,omitempty"`

	// taxability
	Taxability string `json:"taxability,omitempty"`

	// unbilled charges
	UnbilledCharges int32 `json:"unbilled_charges,omitempty"`

	// updated at
	UpdatedAt int64 `json:"updated_at,omitempty"`

	// vat number
	VatNumber string `json:"vat_number,omitempty"`
}

// Validate validates this customer
func (m *Customer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoCollection(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBalances(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBillingAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBillingDateMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBillingDayOfWeek(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBillingDayOfWeekMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContacts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEntityCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFraudFlag(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentMethod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReferralUrls(formats); err != nil {
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

var customerTypeAutoCollectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on","off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypeAutoCollectionPropEnum = append(customerTypeAutoCollectionPropEnum, v)
	}
}

const (

	// CustomerAutoCollectionOn captures enum value "on"
	CustomerAutoCollectionOn string = "on"

	// CustomerAutoCollectionOff captures enum value "off"
	CustomerAutoCollectionOff string = "off"
)

// prop value enum
func (m *Customer) validateAutoCollectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerTypeAutoCollectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validateAutoCollection(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoCollection) { // not required
		return nil
	}

	// value enum
	if err := m.validateAutoCollectionEnum("auto_collection", "body", m.AutoCollection); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateBalances(formats strfmt.Registry) error {

	if swag.IsZero(m.Balances) { // not required
		return nil
	}

	for i := 0; i < len(m.Balances); i++ {

		if swag.IsZero(m.Balances[i]) { // not required
			continue
		}

		if m.Balances[i] != nil {

			if err := m.Balances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("balances" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *Customer) validateBillingAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingAddress) { // not required
		return nil
	}

	if m.BillingAddress != nil {

		if err := m.BillingAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_address")
			}
			return err
		}

	}

	return nil
}

var customerTypeBillingDateModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["using_defaults","manually_set"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypeBillingDateModePropEnum = append(customerTypeBillingDateModePropEnum, v)
	}
}

const (

	// CustomerBillingDateModeUsingDefaults captures enum value "using_defaults"
	CustomerBillingDateModeUsingDefaults string = "using_defaults"

	// CustomerBillingDateModeManuallySet captures enum value "manually_set"
	CustomerBillingDateModeManuallySet string = "manually_set"
)

// prop value enum
func (m *Customer) validateBillingDateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerTypeBillingDateModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validateBillingDateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingDateMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingDateModeEnum("billing_date_mode", "body", m.BillingDateMode); err != nil {
		return err
	}

	return nil
}

var customerTypeBillingDayOfWeekPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sunday","monday","tuesday","wednesday","thursday","friday","saturday"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypeBillingDayOfWeekPropEnum = append(customerTypeBillingDayOfWeekPropEnum, v)
	}
}

const (

	// CustomerBillingDayOfWeekSunday captures enum value "sunday"
	CustomerBillingDayOfWeekSunday string = "sunday"

	// CustomerBillingDayOfWeekMonday captures enum value "monday"
	CustomerBillingDayOfWeekMonday string = "monday"

	// CustomerBillingDayOfWeekTuesday captures enum value "tuesday"
	CustomerBillingDayOfWeekTuesday string = "tuesday"

	// CustomerBillingDayOfWeekWednesday captures enum value "wednesday"
	CustomerBillingDayOfWeekWednesday string = "wednesday"

	// CustomerBillingDayOfWeekThursday captures enum value "thursday"
	CustomerBillingDayOfWeekThursday string = "thursday"

	// CustomerBillingDayOfWeekFriday captures enum value "friday"
	CustomerBillingDayOfWeekFriday string = "friday"

	// CustomerBillingDayOfWeekSaturday captures enum value "saturday"
	CustomerBillingDayOfWeekSaturday string = "saturday"
)

// prop value enum
func (m *Customer) validateBillingDayOfWeekEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerTypeBillingDayOfWeekPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validateBillingDayOfWeek(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingDayOfWeek) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingDayOfWeekEnum("billing_day_of_week", "body", m.BillingDayOfWeek); err != nil {
		return err
	}

	return nil
}

var customerTypeBillingDayOfWeekModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["using_defaults","manually_set"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypeBillingDayOfWeekModePropEnum = append(customerTypeBillingDayOfWeekModePropEnum, v)
	}
}

const (

	// CustomerBillingDayOfWeekModeUsingDefaults captures enum value "using_defaults"
	CustomerBillingDayOfWeekModeUsingDefaults string = "using_defaults"

	// CustomerBillingDayOfWeekModeManuallySet captures enum value "manually_set"
	CustomerBillingDayOfWeekModeManuallySet string = "manually_set"
)

// prop value enum
func (m *Customer) validateBillingDayOfWeekModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerTypeBillingDayOfWeekModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validateBillingDayOfWeekMode(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingDayOfWeekMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingDayOfWeekModeEnum("billing_day_of_week_mode", "body", m.BillingDayOfWeekMode); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateContacts(formats strfmt.Registry) error {

	if swag.IsZero(m.Contacts) { // not required
		return nil
	}

	for i := 0; i < len(m.Contacts); i++ {

		if swag.IsZero(m.Contacts[i]) { // not required
			continue
		}

		if m.Contacts[i] != nil {

			if err := m.Contacts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contacts" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

var customerTypeEntityCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["a","b","c","d","e","f","g","h","i","j","k","l","n","p","q","r","med1","med2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypeEntityCodePropEnum = append(customerTypeEntityCodePropEnum, v)
	}
}

const (

	// CustomerEntityCodeA captures enum value "a"
	CustomerEntityCodeA string = "a"

	// CustomerEntityCodeB captures enum value "b"
	CustomerEntityCodeB string = "b"

	// CustomerEntityCodeC captures enum value "c"
	CustomerEntityCodeC string = "c"

	// CustomerEntityCodeD captures enum value "d"
	CustomerEntityCodeD string = "d"

	// CustomerEntityCodeE captures enum value "e"
	CustomerEntityCodeE string = "e"

	// CustomerEntityCodeF captures enum value "f"
	CustomerEntityCodeF string = "f"

	// CustomerEntityCodeG captures enum value "g"
	CustomerEntityCodeG string = "g"

	// CustomerEntityCodeH captures enum value "h"
	CustomerEntityCodeH string = "h"

	// CustomerEntityCodeI captures enum value "i"
	CustomerEntityCodeI string = "i"

	// CustomerEntityCodeJ captures enum value "j"
	CustomerEntityCodeJ string = "j"

	// CustomerEntityCodeK captures enum value "k"
	CustomerEntityCodeK string = "k"

	// CustomerEntityCodeL captures enum value "l"
	CustomerEntityCodeL string = "l"

	// CustomerEntityCodeN captures enum value "n"
	CustomerEntityCodeN string = "n"

	// CustomerEntityCodeP captures enum value "p"
	CustomerEntityCodeP string = "p"

	// CustomerEntityCodeQ captures enum value "q"
	CustomerEntityCodeQ string = "q"

	// CustomerEntityCodeR captures enum value "r"
	CustomerEntityCodeR string = "r"

	// CustomerEntityCodeMed1 captures enum value "med1"
	CustomerEntityCodeMed1 string = "med1"

	// CustomerEntityCodeMed2 captures enum value "med2"
	CustomerEntityCodeMed2 string = "med2"
)

// prop value enum
func (m *Customer) validateEntityCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerTypeEntityCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validateEntityCode(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityCodeEnum("entity_code", "body", m.EntityCode); err != nil {
		return err
	}

	return nil
}

var customerTypeFraudFlagPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["safe","suspicious","fraudulent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypeFraudFlagPropEnum = append(customerTypeFraudFlagPropEnum, v)
	}
}

const (

	// CustomerFraudFlagSafe captures enum value "safe"
	CustomerFraudFlagSafe string = "safe"

	// CustomerFraudFlagSuspicious captures enum value "suspicious"
	CustomerFraudFlagSuspicious string = "suspicious"

	// CustomerFraudFlagFraudulent captures enum value "fraudulent"
	CustomerFraudFlagFraudulent string = "fraudulent"
)

// prop value enum
func (m *Customer) validateFraudFlagEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerTypeFraudFlagPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validateFraudFlag(formats strfmt.Registry) error {

	if swag.IsZero(m.FraudFlag) { // not required
		return nil
	}

	// value enum
	if err := m.validateFraudFlagEnum("fraud_flag", "body", m.FraudFlag); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validatePaymentMethod(formats strfmt.Registry) error {

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

func (m *Customer) validateReferralUrls(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferralUrls) { // not required
		return nil
	}

	for i := 0; i < len(m.ReferralUrls); i++ {

		if swag.IsZero(m.ReferralUrls[i]) { // not required
			continue
		}

		if m.ReferralUrls[i] != nil {

			if err := m.ReferralUrls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("referral_urls" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

var customerTypeTaxabilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["taxable","exempt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypeTaxabilityPropEnum = append(customerTypeTaxabilityPropEnum, v)
	}
}

const (

	// CustomerTaxabilityTaxable captures enum value "taxable"
	CustomerTaxabilityTaxable string = "taxable"

	// CustomerTaxabilityExempt captures enum value "exempt"
	CustomerTaxabilityExempt string = "exempt"
)

// prop value enum
func (m *Customer) validateTaxabilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerTypeTaxabilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validateTaxability(formats strfmt.Registry) error {

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
func (m *Customer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Customer) UnmarshalBinary(b []byte) error {
	var res Customer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
