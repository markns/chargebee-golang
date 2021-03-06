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

// LinkedInvoice linked invoice
// swagger:model LinkedInvoice
type LinkedInvoice struct {

	// applied amount
	AppliedAmount int32 `json:"applied_amount,omitempty"`

	// applied at
	AppliedAt int64 `json:"applied_at,omitempty"`

	// invoice date
	InvoiceDate int64 `json:"invoice_date,omitempty"`

	// invoice id
	InvoiceID string `json:"invoice_id,omitempty"`

	// invoice status
	InvoiceStatus string `json:"invoice_status,omitempty"`

	// invoice total
	InvoiceTotal int32 `json:"invoice_total,omitempty"`
}

// Validate validates this linked invoice
func (m *LinkedInvoice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvoiceStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var linkedInvoiceTypeInvoiceStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["future","in_trial","active","non_renewing","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		linkedInvoiceTypeInvoiceStatusPropEnum = append(linkedInvoiceTypeInvoiceStatusPropEnum, v)
	}
}

const (

	// LinkedInvoiceInvoiceStatusFuture captures enum value "future"
	LinkedInvoiceInvoiceStatusFuture string = "future"

	// LinkedInvoiceInvoiceStatusInTrial captures enum value "in_trial"
	LinkedInvoiceInvoiceStatusInTrial string = "in_trial"

	// LinkedInvoiceInvoiceStatusActive captures enum value "active"
	LinkedInvoiceInvoiceStatusActive string = "active"

	// LinkedInvoiceInvoiceStatusNonRenewing captures enum value "non_renewing"
	LinkedInvoiceInvoiceStatusNonRenewing string = "non_renewing"

	// LinkedInvoiceInvoiceStatusCancelled captures enum value "cancelled"
	LinkedInvoiceInvoiceStatusCancelled string = "cancelled"
)

// prop value enum
func (m *LinkedInvoice) validateInvoiceStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, linkedInvoiceTypeInvoiceStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LinkedInvoice) validateInvoiceStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.InvoiceStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateInvoiceStatusEnum("invoice_status", "body", m.InvoiceStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinkedInvoice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkedInvoice) UnmarshalBinary(b []byte) error {
	var res LinkedInvoice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
