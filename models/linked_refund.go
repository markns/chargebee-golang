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

// LinkedRefund linked refund
// swagger:model LinkedRefund

type LinkedRefund struct {

	// txn amount
	TxnAmount int32 `json:"txn_amount,omitempty"`

	// txn date
	TxnDate int64 `json:"txn_date,omitempty"`

	// txn id
	TxnID string `json:"txn_id,omitempty"`

	// txn status
	TxnStatus string `json:"txn_status,omitempty"`
}

/* polymorph LinkedRefund txn_amount false */

/* polymorph LinkedRefund txn_date false */

/* polymorph LinkedRefund txn_id false */

/* polymorph LinkedRefund txn_status false */

// Validate validates this linked refund
func (m *LinkedRefund) Validate(formats strfmt.Registry) error {
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

var linkedRefundTypeTxnStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["future","in_trial","active","non_renewing","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		linkedRefundTypeTxnStatusPropEnum = append(linkedRefundTypeTxnStatusPropEnum, v)
	}
}

const (
	// LinkedRefundTxnStatusFuture captures enum value "future"
	LinkedRefundTxnStatusFuture string = "future"
	// LinkedRefundTxnStatusInTrial captures enum value "in_trial"
	LinkedRefundTxnStatusInTrial string = "in_trial"
	// LinkedRefundTxnStatusActive captures enum value "active"
	LinkedRefundTxnStatusActive string = "active"
	// LinkedRefundTxnStatusNonRenewing captures enum value "non_renewing"
	LinkedRefundTxnStatusNonRenewing string = "non_renewing"
	// LinkedRefundTxnStatusCancelled captures enum value "cancelled"
	LinkedRefundTxnStatusCancelled string = "cancelled"
)

// prop value enum
func (m *LinkedRefund) validateTxnStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, linkedRefundTypeTxnStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LinkedRefund) validateTxnStatus(formats strfmt.Registry) error {

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
func (m *LinkedRefund) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkedRefund) UnmarshalBinary(b []byte) error {
	var res LinkedRefund
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
