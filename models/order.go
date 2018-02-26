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

// Order order
// swagger:model Order

type Order struct {

	// batch id
	BatchID string `json:"batch_id,omitempty"`

	// created at
	CreatedAt int64 `json:"created_at,omitempty"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// fulfillment status
	FulfillmentStatus string `json:"fulfillment_status,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// invoice id
	InvoiceID string `json:"invoice_id,omitempty"`

	// note
	Note string `json:"note,omitempty"`

	// reference id
	ReferenceID string `json:"reference_id,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status update at
	StatusUpdateAt int64 `json:"status_update_at,omitempty"`

	// tracking id
	TrackingID string `json:"tracking_id,omitempty"`
}

/* polymorph Order batch_id false */

/* polymorph Order created_at false */

/* polymorph Order created_by false */

/* polymorph Order fulfillment_status false */

/* polymorph Order id false */

/* polymorph Order invoice_id false */

/* polymorph Order note false */

/* polymorph Order reference_id false */

/* polymorph Order status false */

/* polymorph Order status_update_at false */

/* polymorph Order tracking_id false */

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var orderTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["future","in_trial","active","non_renewing","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeStatusPropEnum = append(orderTypeStatusPropEnum, v)
	}
}

const (
	// OrderStatusFuture captures enum value "future"
	OrderStatusFuture string = "future"
	// OrderStatusInTrial captures enum value "in_trial"
	OrderStatusInTrial string = "in_trial"
	// OrderStatusActive captures enum value "active"
	OrderStatusActive string = "active"
	// OrderStatusNonRenewing captures enum value "non_renewing"
	OrderStatusNonRenewing string = "non_renewing"
	// OrderStatusCancelled captures enum value "cancelled"
	OrderStatusCancelled string = "cancelled"
)

// prop value enum
func (m *Order) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, orderTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateStatus(formats strfmt.Registry) error {

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
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
