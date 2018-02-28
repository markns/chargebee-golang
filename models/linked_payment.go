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

// LinkedPayment linked payment
// swagger:model LinkedPayment
type LinkedPayment struct {

	// applied amount
	AppliedAmount int32 `json:"applied_amount,omitempty"`

	// applied at
	AppliedAt int64 `json:"applied_at,omitempty"`

	// txn amount
	TxnAmount int32 `json:"txn_amount,omitempty"`

	// txn date
	TxnDate int64 `json:"txn_date,omitempty"`

	// txn id
	TxnID string `json:"txn_id,omitempty"`

	// txn status
	TxnStatus string `json:"txn_status,omitempty"`
}

// Validate validates this linked payment
func (m *LinkedPayment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTxnStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var linkedPaymentTypeTxnStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["future","in_trial","active","non_renewing","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		linkedPaymentTypeTxnStatusPropEnum = append(linkedPaymentTypeTxnStatusPropEnum, v)
	}
}

const (

	// LinkedPaymentTxnStatusFuture captures enum value "future"
	LinkedPaymentTxnStatusFuture string = "future"

	// LinkedPaymentTxnStatusInTrial captures enum value "in_trial"
	LinkedPaymentTxnStatusInTrial string = "in_trial"

	// LinkedPaymentTxnStatusActive captures enum value "active"
	LinkedPaymentTxnStatusActive string = "active"

	// LinkedPaymentTxnStatusNonRenewing captures enum value "non_renewing"
	LinkedPaymentTxnStatusNonRenewing string = "non_renewing"

	// LinkedPaymentTxnStatusCancelled captures enum value "cancelled"
	LinkedPaymentTxnStatusCancelled string = "cancelled"
)

// prop value enum
func (m *LinkedPayment) validateTxnStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, linkedPaymentTypeTxnStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LinkedPayment) validateTxnStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.TxnStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateTxnStatusEnum("txn_status", "body", m.TxnStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinkedPayment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkedPayment) UnmarshalBinary(b []byte) error {
	var res LinkedPayment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
