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

// SubscriptionImportForCustomerRequest subscription import for customer request
// swagger:model SubscriptionImportForCustomerRequest
type SubscriptionImportForCustomerRequest struct {

	// auto collection
	AutoCollection string `json:"auto_collection,omitempty"`

	// billing cycles
	BillingCycles int32 `json:"billing_cycles,omitempty"`

	// cancelled at
	CancelledAt int64 `json:"cancelled_at,omitempty"`

	// coupon ids
	CouponIds string `json:"coupon_ids,omitempty"`

	// current term end
	CurrentTermEnd int64 `json:"current_term_end,omitempty"`

	// current term start
	CurrentTermStart int64 `json:"current_term_start,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// invoice notes
	InvoiceNotes string `json:"invoice_notes,omitempty"`

	// meta data
	MetaData string `json:"meta_data,omitempty"`

	// payment source id
	PaymentSourceID string `json:"payment_source_id,omitempty"`

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

	// started at
	StartedAt int64 `json:"started_at,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// trial end
	TrialEnd int64 `json:"trial_end,omitempty"`

	// trial start
	TrialStart int64 `json:"trial_start,omitempty"`
}

// Validate validates this subscription import for customer request
func (m *SubscriptionImportForCustomerRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoCollection(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShippingAddressValidationStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var subscriptionImportForCustomerRequestTypeAutoCollectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on","off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionImportForCustomerRequestTypeAutoCollectionPropEnum = append(subscriptionImportForCustomerRequestTypeAutoCollectionPropEnum, v)
	}
}

const (

	// SubscriptionImportForCustomerRequestAutoCollectionOn captures enum value "on"
	SubscriptionImportForCustomerRequestAutoCollectionOn string = "on"

	// SubscriptionImportForCustomerRequestAutoCollectionOff captures enum value "off"
	SubscriptionImportForCustomerRequestAutoCollectionOff string = "off"
)

// prop value enum
func (m *SubscriptionImportForCustomerRequest) validateAutoCollectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionImportForCustomerRequestTypeAutoCollectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionImportForCustomerRequest) validateAutoCollection(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoCollection) { // not required
		return nil
	}

	// value enum
	if err := m.validateAutoCollectionEnum("auto_collection", "body", m.AutoCollection); err != nil {
		return err
	}

	return nil
}

var subscriptionImportForCustomerRequestTypeShippingAddressValidationStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["not_validated","valid","partially_valid","invalid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionImportForCustomerRequestTypeShippingAddressValidationStatusPropEnum = append(subscriptionImportForCustomerRequestTypeShippingAddressValidationStatusPropEnum, v)
	}
}

const (

	// SubscriptionImportForCustomerRequestShippingAddressValidationStatusNotValidated captures enum value "not_validated"
	SubscriptionImportForCustomerRequestShippingAddressValidationStatusNotValidated string = "not_validated"

	// SubscriptionImportForCustomerRequestShippingAddressValidationStatusValid captures enum value "valid"
	SubscriptionImportForCustomerRequestShippingAddressValidationStatusValid string = "valid"

	// SubscriptionImportForCustomerRequestShippingAddressValidationStatusPartiallyValid captures enum value "partially_valid"
	SubscriptionImportForCustomerRequestShippingAddressValidationStatusPartiallyValid string = "partially_valid"

	// SubscriptionImportForCustomerRequestShippingAddressValidationStatusInvalid captures enum value "invalid"
	SubscriptionImportForCustomerRequestShippingAddressValidationStatusInvalid string = "invalid"
)

// prop value enum
func (m *SubscriptionImportForCustomerRequest) validateShippingAddressValidationStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionImportForCustomerRequestTypeShippingAddressValidationStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionImportForCustomerRequest) validateShippingAddressValidationStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ShippingAddressValidationStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateShippingAddressValidationStatusEnum("shipping_address[validation_status]", "body", m.ShippingAddressValidationStatus); err != nil {
		return err
	}

	return nil
}

var subscriptionImportForCustomerRequestTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["future","in_trial","active","non_renewing","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionImportForCustomerRequestTypeStatusPropEnum = append(subscriptionImportForCustomerRequestTypeStatusPropEnum, v)
	}
}

const (

	// SubscriptionImportForCustomerRequestStatusFuture captures enum value "future"
	SubscriptionImportForCustomerRequestStatusFuture string = "future"

	// SubscriptionImportForCustomerRequestStatusInTrial captures enum value "in_trial"
	SubscriptionImportForCustomerRequestStatusInTrial string = "in_trial"

	// SubscriptionImportForCustomerRequestStatusActive captures enum value "active"
	SubscriptionImportForCustomerRequestStatusActive string = "active"

	// SubscriptionImportForCustomerRequestStatusNonRenewing captures enum value "non_renewing"
	SubscriptionImportForCustomerRequestStatusNonRenewing string = "non_renewing"

	// SubscriptionImportForCustomerRequestStatusCancelled captures enum value "cancelled"
	SubscriptionImportForCustomerRequestStatusCancelled string = "cancelled"
)

// prop value enum
func (m *SubscriptionImportForCustomerRequest) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionImportForCustomerRequestTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionImportForCustomerRequest) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionImportForCustomerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionImportForCustomerRequest) UnmarshalBinary(b []byte) error {
	var res SubscriptionImportForCustomerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
