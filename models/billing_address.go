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

// BillingAddress billing address
// swagger:model BillingAddress
type BillingAddress struct {

	// city
	City string `json:"city,omitempty"`

	// company
	Company string `json:"company,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// line1
	Line1 string `json:"line1,omitempty"`

	// line2
	Line2 string `json:"line2,omitempty"`

	// line3
	Line3 string `json:"line3,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// state code
	StateCode string `json:"state_code,omitempty"`

	// validation status
	ValidationStatus string `json:"validation_status,omitempty"`

	// zip
	Zip string `json:"zip,omitempty"`
}

// Validate validates this billing address
func (m *BillingAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidationStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var billingAddressTypeValidationStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["not_validated","valid","partially_valid","invalid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingAddressTypeValidationStatusPropEnum = append(billingAddressTypeValidationStatusPropEnum, v)
	}
}

const (

	// BillingAddressValidationStatusNotValidated captures enum value "not_validated"
	BillingAddressValidationStatusNotValidated string = "not_validated"

	// BillingAddressValidationStatusValid captures enum value "valid"
	BillingAddressValidationStatusValid string = "valid"

	// BillingAddressValidationStatusPartiallyValid captures enum value "partially_valid"
	BillingAddressValidationStatusPartiallyValid string = "partially_valid"

	// BillingAddressValidationStatusInvalid captures enum value "invalid"
	BillingAddressValidationStatusInvalid string = "invalid"
)

// prop value enum
func (m *BillingAddress) validateValidationStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, billingAddressTypeValidationStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BillingAddress) validateValidationStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateValidationStatusEnum("validation_status", "body", m.ValidationStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingAddress) UnmarshalBinary(b []byte) error {
	var res BillingAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
