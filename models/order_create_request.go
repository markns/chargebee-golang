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

// OrderCreateRequest order create request
// swagger:model OrderCreateRequest
type OrderCreateRequest struct {

	// batch id
	BatchID string `json:"batch_id,omitempty"`

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

	// tracking id
	TrackingID string `json:"tracking_id,omitempty"`
}

// Validate validates this order create request
func (m *OrderCreateRequest) Validate(formats strfmt.Registry) error {
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

var orderCreateRequestTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["future","in_trial","active","non_renewing","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderCreateRequestTypeStatusPropEnum = append(orderCreateRequestTypeStatusPropEnum, v)
	}
}

const (

	// OrderCreateRequestStatusFuture captures enum value "future"
	OrderCreateRequestStatusFuture string = "future"

	// OrderCreateRequestStatusInTrial captures enum value "in_trial"
	OrderCreateRequestStatusInTrial string = "in_trial"

	// OrderCreateRequestStatusActive captures enum value "active"
	OrderCreateRequestStatusActive string = "active"

	// OrderCreateRequestStatusNonRenewing captures enum value "non_renewing"
	OrderCreateRequestStatusNonRenewing string = "non_renewing"

	// OrderCreateRequestStatusCancelled captures enum value "cancelled"
	OrderCreateRequestStatusCancelled string = "cancelled"
)

// prop value enum
func (m *OrderCreateRequest) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, orderCreateRequestTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *OrderCreateRequest) validateStatus(formats strfmt.Registry) error {

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
func (m *OrderCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderCreateRequest) UnmarshalBinary(b []byte) error {
	var res OrderCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
