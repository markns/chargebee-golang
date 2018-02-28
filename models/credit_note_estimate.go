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

// CreditNoteEstimate credit note estimate
// swagger:model CreditNoteEstimate
type CreditNoteEstimate struct {

	// amount allocated
	AmountAllocated int32 `json:"amount_allocated,omitempty"`

	// amount available
	AmountAvailable int32 `json:"amount_available,omitempty"`

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

	// reference invoice id
	ReferenceInvoiceID string `json:"reference_invoice_id,omitempty"`

	// sub total
	SubTotal int32 `json:"sub_total,omitempty"`

	// taxes
	Taxes []*Tax `json:"taxes"`

	// total
	Total int32 `json:"total,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this credit note estimate
func (m *CreditNoteEstimate) Validate(formats strfmt.Registry) error {
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

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreditNoteEstimate) validateDiscounts(formats strfmt.Registry) error {

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

func (m *CreditNoteEstimate) validateLineItemDiscounts(formats strfmt.Registry) error {

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

func (m *CreditNoteEstimate) validateLineItemTaxes(formats strfmt.Registry) error {

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

func (m *CreditNoteEstimate) validateLineItems(formats strfmt.Registry) error {

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

var creditNoteEstimateTypePriceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tax_exclusive","tax_inclusive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		creditNoteEstimateTypePriceTypePropEnum = append(creditNoteEstimateTypePriceTypePropEnum, v)
	}
}

const (

	// CreditNoteEstimatePriceTypeTaxExclusive captures enum value "tax_exclusive"
	CreditNoteEstimatePriceTypeTaxExclusive string = "tax_exclusive"

	// CreditNoteEstimatePriceTypeTaxInclusive captures enum value "tax_inclusive"
	CreditNoteEstimatePriceTypeTaxInclusive string = "tax_inclusive"
)

// prop value enum
func (m *CreditNoteEstimate) validatePriceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, creditNoteEstimateTypePriceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreditNoteEstimate) validatePriceType(formats strfmt.Registry) error {

	if swag.IsZero(m.PriceType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePriceTypeEnum("price_type", "body", m.PriceType); err != nil {
		return err
	}

	return nil
}

func (m *CreditNoteEstimate) validateTaxes(formats strfmt.Registry) error {

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

var creditNoteEstimateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["card","paypal_express_checkout","amazon_payments","direct_debit","generic","alipay","unionpay","apple_pay"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		creditNoteEstimateTypeTypePropEnum = append(creditNoteEstimateTypeTypePropEnum, v)
	}
}

const (

	// CreditNoteEstimateTypeCard captures enum value "card"
	CreditNoteEstimateTypeCard string = "card"

	// CreditNoteEstimateTypePaypalExpressCheckout captures enum value "paypal_express_checkout"
	CreditNoteEstimateTypePaypalExpressCheckout string = "paypal_express_checkout"

	// CreditNoteEstimateTypeAmazonPayments captures enum value "amazon_payments"
	CreditNoteEstimateTypeAmazonPayments string = "amazon_payments"

	// CreditNoteEstimateTypeDirectDebit captures enum value "direct_debit"
	CreditNoteEstimateTypeDirectDebit string = "direct_debit"

	// CreditNoteEstimateTypeGeneric captures enum value "generic"
	CreditNoteEstimateTypeGeneric string = "generic"

	// CreditNoteEstimateTypeAlipay captures enum value "alipay"
	CreditNoteEstimateTypeAlipay string = "alipay"

	// CreditNoteEstimateTypeUnionpay captures enum value "unionpay"
	CreditNoteEstimateTypeUnionpay string = "unionpay"

	// CreditNoteEstimateTypeApplePay captures enum value "apple_pay"
	CreditNoteEstimateTypeApplePay string = "apple_pay"
)

// prop value enum
func (m *CreditNoteEstimate) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, creditNoteEstimateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreditNoteEstimate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreditNoteEstimate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreditNoteEstimate) UnmarshalBinary(b []byte) error {
	var res CreditNoteEstimate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
