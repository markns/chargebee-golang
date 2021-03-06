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

// Discount discount
// swagger:model Discount
type Discount struct {

	// amount
	Amount int32 `json:"amount,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// entity id
	EntityID string `json:"entity_id,omitempty"`

	// entity type
	EntityType string `json:"entity_type,omitempty"`
}

// Validate validates this discount
func (m *Discount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var discountTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["customer","subscription","invoice","credit_note","transaction","plan","addon","coupon"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		discountTypeEntityTypePropEnum = append(discountTypeEntityTypePropEnum, v)
	}
}

const (

	// DiscountEntityTypeCustomer captures enum value "customer"
	DiscountEntityTypeCustomer string = "customer"

	// DiscountEntityTypeSubscription captures enum value "subscription"
	DiscountEntityTypeSubscription string = "subscription"

	// DiscountEntityTypeInvoice captures enum value "invoice"
	DiscountEntityTypeInvoice string = "invoice"

	// DiscountEntityTypeCreditNote captures enum value "credit_note"
	DiscountEntityTypeCreditNote string = "credit_note"

	// DiscountEntityTypeTransaction captures enum value "transaction"
	DiscountEntityTypeTransaction string = "transaction"

	// DiscountEntityTypePlan captures enum value "plan"
	DiscountEntityTypePlan string = "plan"

	// DiscountEntityTypeAddon captures enum value "addon"
	DiscountEntityTypeAddon string = "addon"

	// DiscountEntityTypeCoupon captures enum value "coupon"
	DiscountEntityTypeCoupon string = "coupon"
)

// prop value enum
func (m *Discount) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, discountTypeEntityTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Discount) validateEntityType(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityTypeEnum("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Discount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Discount) UnmarshalBinary(b []byte) error {
	var res Discount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
