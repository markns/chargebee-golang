// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CustomerAddContactRequest customer add contact request
// swagger:model CustomerAddContactRequest

type CustomerAddContactRequest struct {

	// contact email
	ContactEmail string `json:"contact[email],omitempty"`

	// contact enabled
	ContactEnabled bool `json:"contact[enabled],omitempty"`

	// contact first name
	ContactFirstName string `json:"contact[first_name],omitempty"`

	// contact id
	ContactID string `json:"contact[id],omitempty"`

	// contact label
	ContactLabel string `json:"contact[label],omitempty"`

	// contact last name
	ContactLastName string `json:"contact[last_name],omitempty"`

	// contact phone
	ContactPhone string `json:"contact[phone],omitempty"`

	// contact send account email
	ContactSendAccountEmail bool `json:"contact[send_account_email],omitempty"`

	// contact send billing email
	ContactSendBillingEmail bool `json:"contact[send_billing_email],omitempty"`
}

/* polymorph CustomerAddContactRequest contact[email] false */

/* polymorph CustomerAddContactRequest contact[enabled] false */

/* polymorph CustomerAddContactRequest contact[first_name] false */

/* polymorph CustomerAddContactRequest contact[id] false */

/* polymorph CustomerAddContactRequest contact[label] false */

/* polymorph CustomerAddContactRequest contact[last_name] false */

/* polymorph CustomerAddContactRequest contact[phone] false */

/* polymorph CustomerAddContactRequest contact[send_account_email] false */

/* polymorph CustomerAddContactRequest contact[send_billing_email] false */

// Validate validates this customer add contact request
func (m *CustomerAddContactRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CustomerAddContactRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerAddContactRequest) UnmarshalBinary(b []byte) error {
	var res CustomerAddContactRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
