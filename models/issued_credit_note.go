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

// IssuedCreditNote issued credit note
// swagger:model IssuedCreditNote

type IssuedCreditNote struct {

	// cn date
	CnDate int64 `json:"cn_date,omitempty"`

	// cn id
	CnID string `json:"cn_id,omitempty"`

	// cn reason code
	CnReasonCode string `json:"cn_reason_code,omitempty"`

	// cn status
	CnStatus string `json:"cn_status,omitempty"`

	// cn total
	CnTotal int32 `json:"cn_total,omitempty"`
}

/* polymorph IssuedCreditNote cn_date false */

/* polymorph IssuedCreditNote cn_id false */

/* polymorph IssuedCreditNote cn_reason_code false */

/* polymorph IssuedCreditNote cn_status false */

/* polymorph IssuedCreditNote cn_total false */

// Validate validates this issued credit note
func (m *IssuedCreditNote) Validate(formats strfmt.Registry) error {
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

var issuedCreditNoteTypeCnReasonCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["write_off","subscription_change","subscription_cancellation","chargeback","product_unsatisfactory","service_unsatisfactory","order_change","order_cancellation","waiver","other","fraudulent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		issuedCreditNoteTypeCnReasonCodePropEnum = append(issuedCreditNoteTypeCnReasonCodePropEnum, v)
	}
}

const (
	// IssuedCreditNoteCnReasonCodeWriteOff captures enum value "write_off"
	IssuedCreditNoteCnReasonCodeWriteOff string = "write_off"
	// IssuedCreditNoteCnReasonCodeSubscriptionChange captures enum value "subscription_change"
	IssuedCreditNoteCnReasonCodeSubscriptionChange string = "subscription_change"
	// IssuedCreditNoteCnReasonCodeSubscriptionCancellation captures enum value "subscription_cancellation"
	IssuedCreditNoteCnReasonCodeSubscriptionCancellation string = "subscription_cancellation"
	// IssuedCreditNoteCnReasonCodeChargeback captures enum value "chargeback"
	IssuedCreditNoteCnReasonCodeChargeback string = "chargeback"
	// IssuedCreditNoteCnReasonCodeProductUnsatisfactory captures enum value "product_unsatisfactory"
	IssuedCreditNoteCnReasonCodeProductUnsatisfactory string = "product_unsatisfactory"
	// IssuedCreditNoteCnReasonCodeServiceUnsatisfactory captures enum value "service_unsatisfactory"
	IssuedCreditNoteCnReasonCodeServiceUnsatisfactory string = "service_unsatisfactory"
	// IssuedCreditNoteCnReasonCodeOrderChange captures enum value "order_change"
	IssuedCreditNoteCnReasonCodeOrderChange string = "order_change"
	// IssuedCreditNoteCnReasonCodeOrderCancellation captures enum value "order_cancellation"
	IssuedCreditNoteCnReasonCodeOrderCancellation string = "order_cancellation"
	// IssuedCreditNoteCnReasonCodeWaiver captures enum value "waiver"
	IssuedCreditNoteCnReasonCodeWaiver string = "waiver"
	// IssuedCreditNoteCnReasonCodeOther captures enum value "other"
	IssuedCreditNoteCnReasonCodeOther string = "other"
	// IssuedCreditNoteCnReasonCodeFraudulent captures enum value "fraudulent"
	IssuedCreditNoteCnReasonCodeFraudulent string = "fraudulent"
)

// prop value enum
func (m *IssuedCreditNote) validateCnReasonCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, issuedCreditNoteTypeCnReasonCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IssuedCreditNote) validateCnReasonCode(formats strfmt.Registry) error {

	if swag.IsZero(m.CnReasonCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateCnReasonCodeEnum("cn_reason_code", "body", m.CnReasonCode); err != nil {
		return err
	}

	return nil
}

var issuedCreditNoteTypeCnStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["future","in_trial","active","non_renewing","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		issuedCreditNoteTypeCnStatusPropEnum = append(issuedCreditNoteTypeCnStatusPropEnum, v)
	}
}

const (
	// IssuedCreditNoteCnStatusFuture captures enum value "future"
	IssuedCreditNoteCnStatusFuture string = "future"
	// IssuedCreditNoteCnStatusInTrial captures enum value "in_trial"
	IssuedCreditNoteCnStatusInTrial string = "in_trial"
	// IssuedCreditNoteCnStatusActive captures enum value "active"
	IssuedCreditNoteCnStatusActive string = "active"
	// IssuedCreditNoteCnStatusNonRenewing captures enum value "non_renewing"
	IssuedCreditNoteCnStatusNonRenewing string = "non_renewing"
	// IssuedCreditNoteCnStatusCancelled captures enum value "cancelled"
	IssuedCreditNoteCnStatusCancelled string = "cancelled"
)

// prop value enum
func (m *IssuedCreditNote) validateCnStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, issuedCreditNoteTypeCnStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IssuedCreditNote) validateCnStatus(formats strfmt.Registry) error {

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
func (m *IssuedCreditNote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssuedCreditNote) UnmarshalBinary(b []byte) error {
	var res IssuedCreditNote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
