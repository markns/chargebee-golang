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

// ResourceMigration resource migration
// swagger:model ResourceMigration
type ResourceMigration struct {

	// created at
	CreatedAt int64 `json:"created_at,omitempty"`

	// entity id
	EntityID string `json:"entity_id,omitempty"`

	// entity type
	EntityType string `json:"entity_type,omitempty"`

	// errors
	Errors string `json:"errors,omitempty"`

	// from site
	FromSite string `json:"from_site,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// updated at
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

// Validate validates this resource migration
func (m *ResourceMigration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
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

var resourceMigrationTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["customer","subscription","invoice","credit_note","transaction","plan","addon","coupon"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceMigrationTypeEntityTypePropEnum = append(resourceMigrationTypeEntityTypePropEnum, v)
	}
}

const (

	// ResourceMigrationEntityTypeCustomer captures enum value "customer"
	ResourceMigrationEntityTypeCustomer string = "customer"

	// ResourceMigrationEntityTypeSubscription captures enum value "subscription"
	ResourceMigrationEntityTypeSubscription string = "subscription"

	// ResourceMigrationEntityTypeInvoice captures enum value "invoice"
	ResourceMigrationEntityTypeInvoice string = "invoice"

	// ResourceMigrationEntityTypeCreditNote captures enum value "credit_note"
	ResourceMigrationEntityTypeCreditNote string = "credit_note"

	// ResourceMigrationEntityTypeTransaction captures enum value "transaction"
	ResourceMigrationEntityTypeTransaction string = "transaction"

	// ResourceMigrationEntityTypePlan captures enum value "plan"
	ResourceMigrationEntityTypePlan string = "plan"

	// ResourceMigrationEntityTypeAddon captures enum value "addon"
	ResourceMigrationEntityTypeAddon string = "addon"

	// ResourceMigrationEntityTypeCoupon captures enum value "coupon"
	ResourceMigrationEntityTypeCoupon string = "coupon"
)

// prop value enum
func (m *ResourceMigration) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceMigrationTypeEntityTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ResourceMigration) validateEntityType(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityTypeEnum("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

var resourceMigrationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["future","in_trial","active","non_renewing","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceMigrationTypeStatusPropEnum = append(resourceMigrationTypeStatusPropEnum, v)
	}
}

const (

	// ResourceMigrationStatusFuture captures enum value "future"
	ResourceMigrationStatusFuture string = "future"

	// ResourceMigrationStatusInTrial captures enum value "in_trial"
	ResourceMigrationStatusInTrial string = "in_trial"

	// ResourceMigrationStatusActive captures enum value "active"
	ResourceMigrationStatusActive string = "active"

	// ResourceMigrationStatusNonRenewing captures enum value "non_renewing"
	ResourceMigrationStatusNonRenewing string = "non_renewing"

	// ResourceMigrationStatusCancelled captures enum value "cancelled"
	ResourceMigrationStatusCancelled string = "cancelled"
)

// prop value enum
func (m *ResourceMigration) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceMigrationTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ResourceMigration) validateStatus(formats strfmt.Registry) error {

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
func (m *ResourceMigration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceMigration) UnmarshalBinary(b []byte) error {
	var res ResourceMigration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
