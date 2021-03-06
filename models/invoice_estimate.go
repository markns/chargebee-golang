// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InvoiceEstimate invoice estimate
// swagger:model InvoiceEstimate
type InvoiceEstimate struct {

	// amount due
	AmountDue int32 `json:"amount_due,omitempty"`

	// amount paid
	AmountPaid int32 `json:"amount_paid,omitempty"`

	// credits applied
	CreditsApplied int32 `json:"credits_applied,omitempty"`

	// currency code
	CurrencyCode string `json:"currency_code,omitempty"`

	// discounts
	Discounts []*Discount `json:"discounts"`

	// line item discounts
	LineItemDiscounts []*LineItemDiscount `json:"line_item_discounts"`

	// line item taxes
	LineItemTaxes []*LineItemTax `json:"line_item_taxes"`

	// line items
	LineItems []*LineItem `json:"line_items"`

	// price type
	PriceType string `json:"price_type,omitempty"`

	// recurring
	Recurring bool `json:"recurring,omitempty"`

	// sub total
	SubTotal int32 `json:"sub_total,omitempty"`

	// taxes
	Taxes []*Tax `json:"taxes"`

	// total
	Total int32 `json:"total,omitempty"`
}

// Validate validates this invoice estimate
func (m *InvoiceEstimate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscounts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLineItemDiscounts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLineItemTaxes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLineItems(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePriceType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTaxes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceEstimate) validateDiscounts(formats strfmt.Registry) error {

	if swag.IsZero(m.Discounts) { // not required
		return nil
	}

	for i := 0; i < len(m.Discounts); i++ {

		if swag.IsZero(m.Discounts[i]) { // not required
			continue
		}

		if m.Discounts[i] != nil {

			if err := m.Discounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("discounts" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *InvoiceEstimate) validateLineItemDiscounts(formats strfmt.Registry) error {

	if swag.IsZero(m.LineItemDiscounts) { // not required
		return nil
	}

	for i := 0; i < len(m.LineItemDiscounts); i++ {

		if swag.IsZero(m.LineItemDiscounts[i]) { // not required
			continue
		}

		if m.LineItemDiscounts[i] != nil {

			if err := m.LineItemDiscounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("line_item_discounts" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *InvoiceEstimate) validateLineItemTaxes(formats strfmt.Registry) error {

	if swag.IsZero(m.LineItemTaxes) { // not required
		return nil
	}

	for i := 0; i < len(m.LineItemTaxes); i++ {

		if swag.IsZero(m.LineItemTaxes[i]) { // not required
			continue
		}

		if m.LineItemTaxes[i] != nil {

			if err := m.LineItemTaxes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("line_item_taxes" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *InvoiceEstimate) validateLineItems(formats strfmt.Registry) error {

	if swag.IsZero(m.LineItems) { // not required
		return nil
	}

	for i := 0; i < len(m.LineItems); i++ {

		if swag.IsZero(m.LineItems[i]) { // not required
			continue
		}

		if m.LineItems[i] != nil {

			if err := m.LineItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("line_items" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

var invoiceEstimateTypePriceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tax_exclusive","tax_inclusive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceEstimateTypePriceTypePropEnum = append(invoiceEstimateTypePriceTypePropEnum, v)
	}
}

const (

	// InvoiceEstimatePriceTypeTaxExclusive captures enum value "tax_exclusive"
	InvoiceEstimatePriceTypeTaxExclusive string = "tax_exclusive"

	// InvoiceEstimatePriceTypeTaxInclusive captures enum value "tax_inclusive"
	InvoiceEstimatePriceTypeTaxInclusive string = "tax_inclusive"
)

// prop value enum
func (m *InvoiceEstimate) validatePriceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, invoiceEstimateTypePriceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceEstimate) validatePriceType(formats strfmt.Registry) error {

	if swag.IsZero(m.PriceType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePriceTypeEnum("price_type", "body", m.PriceType); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceEstimate) validateTaxes(formats strfmt.Registry) error {

	if swag.IsZero(m.Taxes) { // not required
		return nil
	}

	for i := 0; i < len(m.Taxes); i++ {

		if swag.IsZero(m.Taxes[i]) { // not required
			continue
		}

		if m.Taxes[i] != nil {

			if err := m.Taxes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taxes" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceEstimate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceEstimate) UnmarshalBinary(b []byte) error {
	var res InvoiceEstimate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
