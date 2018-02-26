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

// Allocation allocation
// swagger:model Allocation

type Allocation struct {

	// allocated amount
	AllocatedAmount int32 `json:"allocated_amount,omitempty"`

	// allocated at
	AllocatedAt int64 `json:"allocated_at,omitempty"`

	// invoice date
	InvoiceDate int64 `json:"invoice_date,omitempty"`

	// invoice id
	InvoiceID string `json:"invoice_id,omitempty"`

	// invoice status
	InvoiceStatus string `json:"invoice_status,omitempty"`
}

/* polymorph Allocation allocated_amount false */

/* polymorph Allocation allocated_at false */

/* polymorph Allocation invoice_date false */

/* polymorph Allocation invoice_id false */

/* polymorph Allocation invoice_status false */

// Validate validates this allocation
func (m *Allocation) Validate(formats strfmt.Registry) error {
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

var allocationTypeInvoiceStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["future","in_trial","active","non_renewing","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		allocationTypeInvoiceStatusPropEnum = append(allocationTypeInvoiceStatusPropEnum, v)
	}
}

const (
	// AllocationInvoiceStatusFuture captures enum value "future"
	AllocationInvoiceStatusFuture string = "future"
	// AllocationInvoiceStatusInTrial captures enum value "in_trial"
	AllocationInvoiceStatusInTrial string = "in_trial"
	// AllocationInvoiceStatusActive captures enum value "active"
	AllocationInvoiceStatusActive string = "active"
	// AllocationInvoiceStatusNonRenewing captures enum value "non_renewing"
	AllocationInvoiceStatusNonRenewing string = "non_renewing"
	// AllocationInvoiceStatusCancelled captures enum value "cancelled"
	AllocationInvoiceStatusCancelled string = "cancelled"
)

// prop value enum
func (m *Allocation) validateInvoiceStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, allocationTypeInvoiceStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Allocation) validateInvoiceStatus(formats strfmt.Registry) error {

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
func (m *Allocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Allocation) UnmarshalBinary(b []byte) error {
	var res Allocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
