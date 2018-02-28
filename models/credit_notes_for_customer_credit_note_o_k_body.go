// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreditNotesForCustomerCreditNoteOKBody credit notes for customer credit note o k body
// swagger:model creditNotesForCustomerCreditNoteOKBody
type CreditNotesForCustomerCreditNoteOKBody struct {

	// list
	List []*CreditNoteResponse `json:"list"`

	// next offset
	NextOffset string `json:"next_offset,omitempty"`
}

// Validate validates this credit notes for customer credit note o k body
func (m *CreditNotesForCustomerCreditNoteOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreditNotesForCustomerCreditNoteOKBody) validateList(formats strfmt.Registry) error {

	if swag.IsZero(m.List) { // not required
		return nil
	}

	for i := 0; i < len(m.List); i++ {

		if swag.IsZero(m.List[i]) { // not required
			continue
		}

		if m.List[i] != nil {

			if err := m.List[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreditNotesForCustomerCreditNoteOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreditNotesForCustomerCreditNoteOKBody) UnmarshalBinary(b []byte) error {
	var res CreditNotesForCustomerCreditNoteOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
