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

// AddressUpdateRequest address update request
// swagger:model AddressUpdateRequest

type AddressUpdateRequest struct {

	// addr
	Addr string `json:"addr,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// company
	Company string `json:"company,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// extended addr
	ExtendedAddr string `json:"extended_addr,omitempty"`

	// extended addr2
	ExtendedAddr2 string `json:"extended_addr2,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// state code
	StateCode string `json:"state_code,omitempty"`

	// subscription id
	SubscriptionID string `json:"subscription_id,omitempty"`

	// validation status
	ValidationStatus string `json:"validation_status,omitempty"`

	// zip
	Zip string `json:"zip,omitempty"`
}

/* polymorph AddressUpdateRequest addr false */

/* polymorph AddressUpdateRequest city false */

/* polymorph AddressUpdateRequest company false */

/* polymorph AddressUpdateRequest country false */

/* polymorph AddressUpdateRequest email false */

/* polymorph AddressUpdateRequest extended_addr false */

/* polymorph AddressUpdateRequest extended_addr2 false */

/* polymorph AddressUpdateRequest first_name false */

/* polymorph AddressUpdateRequest label false */

/* polymorph AddressUpdateRequest last_name false */

/* polymorph AddressUpdateRequest phone false */

/* polymorph AddressUpdateRequest state false */

/* polymorph AddressUpdateRequest state_code false */

/* polymorph AddressUpdateRequest subscription_id false */

/* polymorph AddressUpdateRequest validation_status false */

/* polymorph AddressUpdateRequest zip false */

// Validate validates this address update request
func (m *AddressUpdateRequest) Validate(formats strfmt.Registry) error {
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

var addressUpdateRequestTypeValidationStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["not_validated","valid","partially_valid","invalid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addressUpdateRequestTypeValidationStatusPropEnum = append(addressUpdateRequestTypeValidationStatusPropEnum, v)
	}
}

const (
	// AddressUpdateRequestValidationStatusNotValidated captures enum value "not_validated"
	AddressUpdateRequestValidationStatusNotValidated string = "not_validated"
	// AddressUpdateRequestValidationStatusValid captures enum value "valid"
	AddressUpdateRequestValidationStatusValid string = "valid"
	// AddressUpdateRequestValidationStatusPartiallyValid captures enum value "partially_valid"
	AddressUpdateRequestValidationStatusPartiallyValid string = "partially_valid"
	// AddressUpdateRequestValidationStatusInvalid captures enum value "invalid"
	AddressUpdateRequestValidationStatusInvalid string = "invalid"
)

// prop value enum
func (m *AddressUpdateRequest) validateValidationStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addressUpdateRequestTypeValidationStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AddressUpdateRequest) validateValidationStatus(formats strfmt.Registry) error {

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
func (m *AddressUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddressUpdateRequest) UnmarshalBinary(b []byte) error {
	var res AddressUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
