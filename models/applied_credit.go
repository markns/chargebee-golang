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

// AppliedCredit applied credit
// swagger:model AppliedCredit

type AppliedCredit struct {

	// applied amount
	AppliedAmount int32 `json:"applied_amount,omitempty"`

	// applied at
	AppliedAt int64 `json:"applied_at,omitempty"`

	// cn date
	CnDate int64 `json:"cn_date,omitempty"`

	// cn id
	CnID string `json:"cn_id,omitempty"`

	// cn reason code
	CnReasonCode string `json:"cn_reason_code,omitempty"`

	// cn status
	CnStatus string `json:"cn_status,omitempty"`
}

/* polymorph AppliedCredit applied_amount false */

/* polymorph AppliedCredit applied_at false */

/* polymorph AppliedCredit cn_date false */

/* polymorph AppliedCredit cn_id false */

/* polymorph AppliedCredit cn_reason_code false */

/* polymorph AppliedCredit cn_status false */

// Validate validates this applied credit
func (m *AppliedCredit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCnReasonCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCnStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var appliedCreditTypeCnReasonCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["write_off","subscription_change","subscription_cancellation","chargeback","product_unsatisfactory","service_unsatisfactory","order_change","order_cancellation","waiver","other","fraudulent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appliedCreditTypeCnReasonCodePropEnum = append(appliedCreditTypeCnReasonCodePropEnum, v)
	}
}

const (
	// AppliedCreditCnReasonCodeWriteOff captures enum value "write_off"
	AppliedCreditCnReasonCodeWriteOff string = "write_off"
	// AppliedCreditCnReasonCodeSubscriptionChange captures enum value "subscription_change"
	AppliedCreditCnReasonCodeSubscriptionChange string = "subscription_change"
	// AppliedCreditCnReasonCodeSubscriptionCancellation captures enum value "subscription_cancellation"
	AppliedCreditCnReasonCodeSubscriptionCancellation string = "subscription_cancellation"
	// AppliedCreditCnReasonCodeChargeback captures enum value "chargeback"
	AppliedCreditCnReasonCodeChargeback string = "chargeback"
	// AppliedCreditCnReasonCodeProductUnsatisfactory captures enum value "product_unsatisfactory"
	AppliedCreditCnReasonCodeProductUnsatisfactory string = "product_unsatisfactory"
	// AppliedCreditCnReasonCodeServiceUnsatisfactory captures enum value "service_unsatisfactory"
	AppliedCreditCnReasonCodeServiceUnsatisfactory string = "service_unsatisfactory"
	// AppliedCreditCnReasonCodeOrderChange captures enum value "order_change"
	AppliedCreditCnReasonCodeOrderChange string = "order_change"
	// AppliedCreditCnReasonCodeOrderCancellation captures enum value "order_cancellation"
	AppliedCreditCnReasonCodeOrderCancellation string = "order_cancellation"
	// AppliedCreditCnReasonCodeWaiver captures enum value "waiver"
	AppliedCreditCnReasonCodeWaiver string = "waiver"
	// AppliedCreditCnReasonCodeOther captures enum value "other"
	AppliedCreditCnReasonCodeOther string = "other"
	// AppliedCreditCnReasonCodeFraudulent captures enum value "fraudulent"
	AppliedCreditCnReasonCodeFraudulent string = "fraudulent"
)

// prop value enum
func (m *AppliedCredit) validateCnReasonCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, appliedCreditTypeCnReasonCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AppliedCredit) validateCnReasonCode(formats strfmt.Registry) error {

	if swag.IsZero(m.CnReasonCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateCnReasonCodeEnum("cn_reason_code", "body", m.CnReasonCode); err != nil {
		return err
	}

	return nil
}

var appliedCreditTypeCnStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["in_progress","success","voided","failure","timeout","needs_attention"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appliedCreditTypeCnStatusPropEnum = append(appliedCreditTypeCnStatusPropEnum, v)
	}
}

const (
	// AppliedCreditCnStatusInProgress captures enum value "in_progress"
	AppliedCreditCnStatusInProgress string = "in_progress"
	// AppliedCreditCnStatusSuccess captures enum value "success"
	AppliedCreditCnStatusSuccess string = "success"
	// AppliedCreditCnStatusVoided captures enum value "voided"
	AppliedCreditCnStatusVoided string = "voided"
	// AppliedCreditCnStatusFailure captures enum value "failure"
	AppliedCreditCnStatusFailure string = "failure"
	// AppliedCreditCnStatusTimeout captures enum value "timeout"
	AppliedCreditCnStatusTimeout string = "timeout"
	// AppliedCreditCnStatusNeedsAttention captures enum value "needs_attention"
	AppliedCreditCnStatusNeedsAttention string = "needs_attention"
)

// prop value enum
func (m *AppliedCredit) validateCnStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, appliedCreditTypeCnStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AppliedCredit) validateCnStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.CnStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateCnStatusEnum("cn_status", "body", m.CnStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppliedCredit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppliedCredit) UnmarshalBinary(b []byte) error {
	var res AppliedCredit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}