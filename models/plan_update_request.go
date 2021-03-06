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

// PlanUpdateRequest plan update request
// swagger:model PlanUpdateRequest
type PlanUpdateRequest struct {

	// accounting category1
	AccountingCategory1 string `json:"accounting_category1,omitempty"`

	// accounting category2
	AccountingCategory2 string `json:"accounting_category2,omitempty"`

	// accounting code
	AccountingCode string `json:"accounting_code,omitempty"`

	// billing cycles
	BillingCycles int32 `json:"billing_cycles,omitempty"`

	// charge model
	ChargeModel string `json:"charge_model,omitempty"`

	// currency code
	CurrencyCode string `json:"currency_code,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// downgrade penalty
	DowngradePenalty float64 `json:"downgrade_penalty,omitempty"`

	// enabled in hosted pages
	EnabledInHostedPages bool `json:"enabled_in_hosted_pages,omitempty"`

	// enabled in portal
	EnabledInPortal bool `json:"enabled_in_portal,omitempty"`

	// free quantity
	FreeQuantity int32 `json:"free_quantity,omitempty"`

	// invoice name
	InvoiceName string `json:"invoice_name,omitempty"`

	// invoice notes
	InvoiceNotes string `json:"invoice_notes,omitempty"`

	// meta data
	MetaData interface{} `json:"meta_data,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// period
	Period int32 `json:"period,omitempty"`

	// period unit
	PeriodUnit string `json:"period_unit,omitempty"`

	// price
	Price int32 `json:"price,omitempty"`

	// redirect url
	RedirectURL string `json:"redirect_url,omitempty"`

	// setup cost
	SetupCost int32 `json:"setup_cost,omitempty"`

	// sku
	Sku string `json:"sku,omitempty"`

	// tax code
	TaxCode string `json:"tax_code,omitempty"`

	// tax profile id
	TaxProfileID string `json:"tax_profile_id,omitempty"`

	// taxable
	Taxable bool `json:"taxable,omitempty"`

	// trial period
	TrialPeriod int32 `json:"trial_period,omitempty"`

	// trial period unit
	TrialPeriodUnit string `json:"trial_period_unit,omitempty"`
}

// Validate validates this plan update request
func (m *PlanUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChargeModel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePeriodUnit(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTrialPeriodUnit(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var planUpdateRequestTypeChargeModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["full_charge","prorate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		planUpdateRequestTypeChargeModelPropEnum = append(planUpdateRequestTypeChargeModelPropEnum, v)
	}
}

const (

	// PlanUpdateRequestChargeModelFullCharge captures enum value "full_charge"
	PlanUpdateRequestChargeModelFullCharge string = "full_charge"

	// PlanUpdateRequestChargeModelProrate captures enum value "prorate"
	PlanUpdateRequestChargeModelProrate string = "prorate"
)

// prop value enum
func (m *PlanUpdateRequest) validateChargeModelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, planUpdateRequestTypeChargeModelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PlanUpdateRequest) validateChargeModel(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargeModel) { // not required
		return nil
	}

	// value enum
	if err := m.validateChargeModelEnum("charge_model", "body", m.ChargeModel); err != nil {
		return err
	}

	return nil
}

var planUpdateRequestTypePeriodUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["week","month","year"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		planUpdateRequestTypePeriodUnitPropEnum = append(planUpdateRequestTypePeriodUnitPropEnum, v)
	}
}

const (

	// PlanUpdateRequestPeriodUnitWeek captures enum value "week"
	PlanUpdateRequestPeriodUnitWeek string = "week"

	// PlanUpdateRequestPeriodUnitMonth captures enum value "month"
	PlanUpdateRequestPeriodUnitMonth string = "month"

	// PlanUpdateRequestPeriodUnitYear captures enum value "year"
	PlanUpdateRequestPeriodUnitYear string = "year"
)

// prop value enum
func (m *PlanUpdateRequest) validatePeriodUnitEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, planUpdateRequestTypePeriodUnitPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PlanUpdateRequest) validatePeriodUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.PeriodUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validatePeriodUnitEnum("period_unit", "body", m.PeriodUnit); err != nil {
		return err
	}

	return nil
}

var planUpdateRequestTypeTrialPeriodUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["day","month"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		planUpdateRequestTypeTrialPeriodUnitPropEnum = append(planUpdateRequestTypeTrialPeriodUnitPropEnum, v)
	}
}

const (

	// PlanUpdateRequestTrialPeriodUnitDay captures enum value "day"
	PlanUpdateRequestTrialPeriodUnitDay string = "day"

	// PlanUpdateRequestTrialPeriodUnitMonth captures enum value "month"
	PlanUpdateRequestTrialPeriodUnitMonth string = "month"
)

// prop value enum
func (m *PlanUpdateRequest) validateTrialPeriodUnitEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, planUpdateRequestTypeTrialPeriodUnitPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PlanUpdateRequest) validateTrialPeriodUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.TrialPeriodUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateTrialPeriodUnitEnum("trial_period_unit", "body", m.TrialPeriodUnit); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanUpdateRequest) UnmarshalBinary(b []byte) error {
	var res PlanUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
