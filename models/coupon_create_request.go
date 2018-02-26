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

// CouponCreateRequest coupon create request
// swagger:model CouponCreateRequest

type CouponCreateRequest struct {

	// addon constraint
	AddonConstraint string `json:"addon_constraint,omitempty"`

	// addon ids
	AddonIds string `json:"addon_ids,omitempty"`

	// apply on
	ApplyOn string `json:"apply_on,omitempty"`

	// currency code
	CurrencyCode string `json:"currency_code,omitempty"`

	// discount amount
	DiscountAmount int32 `json:"discount_amount,omitempty"`

	// discount percentage
	DiscountPercentage float64 `json:"discount_percentage,omitempty"`

	// discount quantity
	DiscountQuantity int32 `json:"discount_quantity,omitempty"`

	// discount type
	DiscountType string `json:"discount_type,omitempty"`

	// duration month
	DurationMonth int32 `json:"duration_month,omitempty"`

	// duration type
	DurationType string `json:"duration_type,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// invoice name
	InvoiceName string `json:"invoice_name,omitempty"`

	// invoice notes
	InvoiceNotes string `json:"invoice_notes,omitempty"`

	// max redemptions
	MaxRedemptions int32 `json:"max_redemptions,omitempty"`

	// meta data
	MetaData string `json:"meta_data,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// plan constraint
	PlanConstraint string `json:"plan_constraint,omitempty"`

	// plan ids
	PlanIds string `json:"plan_ids,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// valid till
	ValidTill int64 `json:"valid_till,omitempty"`
}

/* polymorph CouponCreateRequest addon_constraint false */

/* polymorph CouponCreateRequest addon_ids false */

/* polymorph CouponCreateRequest apply_on false */

/* polymorph CouponCreateRequest currency_code false */

/* polymorph CouponCreateRequest discount_amount false */

/* polymorph CouponCreateRequest discount_percentage false */

/* polymorph CouponCreateRequest discount_quantity false */

/* polymorph CouponCreateRequest discount_type false */

/* polymorph CouponCreateRequest duration_month false */

/* polymorph CouponCreateRequest duration_type false */

/* polymorph CouponCreateRequest id false */

/* polymorph CouponCreateRequest invoice_name false */

/* polymorph CouponCreateRequest invoice_notes false */

/* polymorph CouponCreateRequest max_redemptions false */

/* polymorph CouponCreateRequest meta_data false */

/* polymorph CouponCreateRequest name false */

/* polymorph CouponCreateRequest plan_constraint false */

/* polymorph CouponCreateRequest plan_ids false */

/* polymorph CouponCreateRequest status false */

/* polymorph CouponCreateRequest valid_till false */

// Validate validates this coupon create request
func (m *CouponCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddonConstraint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateApplyOn(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDiscountType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDurationType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlanConstraint(formats); err != nil {
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

var couponCreateRequestTypeAddonConstraintPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","all","specific","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		couponCreateRequestTypeAddonConstraintPropEnum = append(couponCreateRequestTypeAddonConstraintPropEnum, v)
	}
}

const (
	// CouponCreateRequestAddonConstraintNone captures enum value "none"
	CouponCreateRequestAddonConstraintNone string = "none"
	// CouponCreateRequestAddonConstraintAll captures enum value "all"
	CouponCreateRequestAddonConstraintAll string = "all"
	// CouponCreateRequestAddonConstraintSpecific captures enum value "specific"
	CouponCreateRequestAddonConstraintSpecific string = "specific"
	// CouponCreateRequestAddonConstraintNotApplicable captures enum value "not_applicable"
	CouponCreateRequestAddonConstraintNotApplicable string = "not_applicable"
)

// prop value enum
func (m *CouponCreateRequest) validateAddonConstraintEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, couponCreateRequestTypeAddonConstraintPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CouponCreateRequest) validateAddonConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.AddonConstraint) { // not required
		return nil
	}

	// value enum
	if err := m.validateAddonConstraintEnum("addon_constraint", "body", m.AddonConstraint); err != nil {
		return err
	}

	return nil
}

var couponCreateRequestTypeApplyOnPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["invoice_amount","specified_items_total","each_specified_item","each_unit_of_specified_items"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		couponCreateRequestTypeApplyOnPropEnum = append(couponCreateRequestTypeApplyOnPropEnum, v)
	}
}

const (
	// CouponCreateRequestApplyOnInvoiceAmount captures enum value "invoice_amount"
	CouponCreateRequestApplyOnInvoiceAmount string = "invoice_amount"
	// CouponCreateRequestApplyOnSpecifiedItemsTotal captures enum value "specified_items_total"
	CouponCreateRequestApplyOnSpecifiedItemsTotal string = "specified_items_total"
	// CouponCreateRequestApplyOnEachSpecifiedItem captures enum value "each_specified_item"
	CouponCreateRequestApplyOnEachSpecifiedItem string = "each_specified_item"
	// CouponCreateRequestApplyOnEachUnitOfSpecifiedItems captures enum value "each_unit_of_specified_items"
	CouponCreateRequestApplyOnEachUnitOfSpecifiedItems string = "each_unit_of_specified_items"
)

// prop value enum
func (m *CouponCreateRequest) validateApplyOnEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, couponCreateRequestTypeApplyOnPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CouponCreateRequest) validateApplyOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ApplyOn) { // not required
		return nil
	}

	// value enum
	if err := m.validateApplyOnEnum("apply_on", "body", m.ApplyOn); err != nil {
		return err
	}

	return nil
}

var couponCreateRequestTypeDiscountTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["item_level_coupon","document_level_coupon","promotional_credits","prorated_credits"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		couponCreateRequestTypeDiscountTypePropEnum = append(couponCreateRequestTypeDiscountTypePropEnum, v)
	}
}

const (
	// CouponCreateRequestDiscountTypeItemLevelCoupon captures enum value "item_level_coupon"
	CouponCreateRequestDiscountTypeItemLevelCoupon string = "item_level_coupon"
	// CouponCreateRequestDiscountTypeDocumentLevelCoupon captures enum value "document_level_coupon"
	CouponCreateRequestDiscountTypeDocumentLevelCoupon string = "document_level_coupon"
	// CouponCreateRequestDiscountTypePromotionalCredits captures enum value "promotional_credits"
	CouponCreateRequestDiscountTypePromotionalCredits string = "promotional_credits"
	// CouponCreateRequestDiscountTypeProratedCredits captures enum value "prorated_credits"
	CouponCreateRequestDiscountTypeProratedCredits string = "prorated_credits"
)

// prop value enum
func (m *CouponCreateRequest) validateDiscountTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, couponCreateRequestTypeDiscountTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CouponCreateRequest) validateDiscountType(formats strfmt.Registry) error {

	if swag.IsZero(m.DiscountType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDiscountTypeEnum("discount_type", "body", m.DiscountType); err != nil {
		return err
	}

	return nil
}

var couponCreateRequestTypeDurationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["one_time","forever","limited_period"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		couponCreateRequestTypeDurationTypePropEnum = append(couponCreateRequestTypeDurationTypePropEnum, v)
	}
}

const (
	// CouponCreateRequestDurationTypeOneTime captures enum value "one_time"
	CouponCreateRequestDurationTypeOneTime string = "one_time"
	// CouponCreateRequestDurationTypeForever captures enum value "forever"
	CouponCreateRequestDurationTypeForever string = "forever"
	// CouponCreateRequestDurationTypeLimitedPeriod captures enum value "limited_period"
	CouponCreateRequestDurationTypeLimitedPeriod string = "limited_period"
)

// prop value enum
func (m *CouponCreateRequest) validateDurationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, couponCreateRequestTypeDurationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CouponCreateRequest) validateDurationType(formats strfmt.Registry) error {

	if swag.IsZero(m.DurationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationTypeEnum("duration_type", "body", m.DurationType); err != nil {
		return err
	}

	return nil
}

var couponCreateRequestTypePlanConstraintPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","all","specific","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		couponCreateRequestTypePlanConstraintPropEnum = append(couponCreateRequestTypePlanConstraintPropEnum, v)
	}
}

const (
	// CouponCreateRequestPlanConstraintNone captures enum value "none"
	CouponCreateRequestPlanConstraintNone string = "none"
	// CouponCreateRequestPlanConstraintAll captures enum value "all"
	CouponCreateRequestPlanConstraintAll string = "all"
	// CouponCreateRequestPlanConstraintSpecific captures enum value "specific"
	CouponCreateRequestPlanConstraintSpecific string = "specific"
	// CouponCreateRequestPlanConstraintNotApplicable captures enum value "not_applicable"
	CouponCreateRequestPlanConstraintNotApplicable string = "not_applicable"
)

// prop value enum
func (m *CouponCreateRequest) validatePlanConstraintEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, couponCreateRequestTypePlanConstraintPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CouponCreateRequest) validatePlanConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.PlanConstraint) { // not required
		return nil
	}

	// value enum
	if err := m.validatePlanConstraintEnum("plan_constraint", "body", m.PlanConstraint); err != nil {
		return err
	}

	return nil
}

var couponCreateRequestTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["in_progress","success","voided","failure","timeout","needs_attention"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		couponCreateRequestTypeStatusPropEnum = append(couponCreateRequestTypeStatusPropEnum, v)
	}
}

const (
	// CouponCreateRequestStatusInProgress captures enum value "in_progress"
	CouponCreateRequestStatusInProgress string = "in_progress"
	// CouponCreateRequestStatusSuccess captures enum value "success"
	CouponCreateRequestStatusSuccess string = "success"
	// CouponCreateRequestStatusVoided captures enum value "voided"
	CouponCreateRequestStatusVoided string = "voided"
	// CouponCreateRequestStatusFailure captures enum value "failure"
	CouponCreateRequestStatusFailure string = "failure"
	// CouponCreateRequestStatusTimeout captures enum value "timeout"
	CouponCreateRequestStatusTimeout string = "timeout"
	// CouponCreateRequestStatusNeedsAttention captures enum value "needs_attention"
	CouponCreateRequestStatusNeedsAttention string = "needs_attention"
)

// prop value enum
func (m *CouponCreateRequest) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, couponCreateRequestTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CouponCreateRequest) validateStatus(formats strfmt.Registry) error {

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
func (m *CouponCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CouponCreateRequest) UnmarshalBinary(b []byte) error {
	var res CouponCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
