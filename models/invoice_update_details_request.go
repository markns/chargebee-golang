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

// InvoiceUpdateDetailsRequest invoice update details request
// swagger:model InvoiceUpdateDetailsRequest

type InvoiceUpdateDetailsRequest struct {

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

	// po number
	PoNumber string `json:"po_number,omitempty"`

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

	// vat number
	VatNumber string `json:"vat_number,omitempty"`
}

/* polymorph InvoiceUpdateDetailsRequest billing_address[city] false */

/* polymorph InvoiceUpdateDetailsRequest billing_address[company] false */

/* polymorph InvoiceUpdateDetailsRequest billing_address[country] false */

/* polymorph InvoiceUpdateDetailsRequest billing_address[email] false */

/* polymorph InvoiceUpdateDetailsRequest billing_address[first_name] false */

/* polymorph InvoiceUpdateDetailsRequest billing_address[last_name] false */

/* polymorph InvoiceUpdateDetailsRequest billing_address[line1] false */

/* polymorph InvoiceUpdateDetailsRequest billing_address[line2] false */

/* polymorph InvoiceUpdateDetailsRequest billing_address[line3] false */

/* polymorph InvoiceUpdateDetailsRequest billing_address[phone] false */

/* polymorph InvoiceUpdateDetailsRequest billing_address[state] false */

/* polymorph InvoiceUpdateDetailsRequest billing_address[state_code] false */

/* polymorph InvoiceUpdateDetailsRequest billing_address[validation_status] false */

/* polymorph InvoiceUpdateDetailsRequest billing_address[zip] false */

/* polymorph InvoiceUpdateDetailsRequest po_number false */

/* polymorph InvoiceUpdateDetailsRequest shipping_address[city] false */

/* polymorph InvoiceUpdateDetailsRequest shipping_address[company] false */

/* polymorph InvoiceUpdateDetailsRequest shipping_address[country] false */

/* polymorph InvoiceUpdateDetailsRequest shipping_address[email] false */

/* polymorph InvoiceUpdateDetailsRequest shipping_address[first_name] false */

/* polymorph InvoiceUpdateDetailsRequest shipping_address[last_name] false */

/* polymorph InvoiceUpdateDetailsRequest shipping_address[line1] false */

/* polymorph InvoiceUpdateDetailsRequest shipping_address[line2] false */

/* polymorph InvoiceUpdateDetailsRequest shipping_address[line3] false */

/* polymorph InvoiceUpdateDetailsRequest shipping_address[phone] false */

/* polymorph InvoiceUpdateDetailsRequest shipping_address[state] false */

/* polymorph InvoiceUpdateDetailsRequest shipping_address[state_code] false */

/* polymorph InvoiceUpdateDetailsRequest shipping_address[validation_status] false */

/* polymorph InvoiceUpdateDetailsRequest shipping_address[zip] false */

/* polymorph InvoiceUpdateDetailsRequest vat_number false */

// Validate validates this invoice update details request
func (m *InvoiceUpdateDetailsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingAddressValidationStatus(formats); err != nil {
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

var invoiceUpdateDetailsRequestTypeBillingAddressValidationStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["not_validated","valid","partially_valid","invalid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceUpdateDetailsRequestTypeBillingAddressValidationStatusPropEnum = append(invoiceUpdateDetailsRequestTypeBillingAddressValidationStatusPropEnum, v)
	}
}

const (
	// InvoiceUpdateDetailsRequestBillingAddressValidationStatusNotValidated captures enum value "not_validated"
	InvoiceUpdateDetailsRequestBillingAddressValidationStatusNotValidated string = "not_validated"
	// InvoiceUpdateDetailsRequestBillingAddressValidationStatusValid captures enum value "valid"
	InvoiceUpdateDetailsRequestBillingAddressValidationStatusValid string = "valid"
	// InvoiceUpdateDetailsRequestBillingAddressValidationStatusPartiallyValid captures enum value "partially_valid"
	InvoiceUpdateDetailsRequestBillingAddressValidationStatusPartiallyValid string = "partially_valid"
	// InvoiceUpdateDetailsRequestBillingAddressValidationStatusInvalid captures enum value "invalid"
	InvoiceUpdateDetailsRequestBillingAddressValidationStatusInvalid string = "invalid"
)

// prop value enum
func (m *InvoiceUpdateDetailsRequest) validateBillingAddressValidationStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, invoiceUpdateDetailsRequestTypeBillingAddressValidationStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceUpdateDetailsRequest) validateBillingAddressValidationStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingAddressValidationStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingAddressValidationStatusEnum("billing_address[validation_status]", "body", m.BillingAddressValidationStatus); err != nil {
		return err
	}

	return nil
}

var invoiceUpdateDetailsRequestTypeShippingAddressValidationStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["not_validated","valid","partially_valid","invalid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceUpdateDetailsRequestTypeShippingAddressValidationStatusPropEnum = append(invoiceUpdateDetailsRequestTypeShippingAddressValidationStatusPropEnum, v)
	}
}

const (
	// InvoiceUpdateDetailsRequestShippingAddressValidationStatusNotValidated captures enum value "not_validated"
	InvoiceUpdateDetailsRequestShippingAddressValidationStatusNotValidated string = "not_validated"
	// InvoiceUpdateDetailsRequestShippingAddressValidationStatusValid captures enum value "valid"
	InvoiceUpdateDetailsRequestShippingAddressValidationStatusValid string = "valid"
	// InvoiceUpdateDetailsRequestShippingAddressValidationStatusPartiallyValid captures enum value "partially_valid"
	InvoiceUpdateDetailsRequestShippingAddressValidationStatusPartiallyValid string = "partially_valid"
	// InvoiceUpdateDetailsRequestShippingAddressValidationStatusInvalid captures enum value "invalid"
	InvoiceUpdateDetailsRequestShippingAddressValidationStatusInvalid string = "invalid"
)

// prop value enum
func (m *InvoiceUpdateDetailsRequest) validateShippingAddressValidationStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, invoiceUpdateDetailsRequestTypeShippingAddressValidationStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceUpdateDetailsRequest) validateShippingAddressValidationStatus(formats strfmt.Registry) error {

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
func (m *InvoiceUpdateDetailsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceUpdateDetailsRequest) UnmarshalBinary(b []byte) error {
	var res InvoiceUpdateDetailsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}