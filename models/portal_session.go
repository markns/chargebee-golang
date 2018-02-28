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

// PortalSession portal session
// swagger:model PortalSession
type PortalSession struct {

	// access url
	AccessURL string `json:"access_url,omitempty"`

	// created at
	CreatedAt int64 `json:"created_at,omitempty"`

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// expires at
	ExpiresAt int64 `json:"expires_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// linked customers
	LinkedCustomers []*LinkedCustomer `json:"linked_customers"`

	// login at
	LoginAt int64 `json:"login_at,omitempty"`

	// login ipaddress
	LoginIpaddress string `json:"login_ipaddress,omitempty"`

	// logout at
	LogoutAt int64 `json:"logout_at,omitempty"`

	// logout ipaddress
	LogoutIpaddress string `json:"logout_ipaddress,omitempty"`

	// redirect url
	RedirectURL string `json:"redirect_url,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this portal session
func (m *PortalSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinkedCustomers(formats); err != nil {
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

func (m *PortalSession) validateLinkedCustomers(formats strfmt.Registry) error {

	if swag.IsZero(m.LinkedCustomers) { // not required
		return nil
	}

	for i := 0; i < len(m.LinkedCustomers); i++ {

		if swag.IsZero(m.LinkedCustomers[i]) { // not required
			continue
		}

		if m.LinkedCustomers[i] != nil {

			if err := m.LinkedCustomers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("linked_customers" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

var portalSessionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["future","in_trial","active","non_renewing","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portalSessionTypeStatusPropEnum = append(portalSessionTypeStatusPropEnum, v)
	}
}

const (

	// PortalSessionStatusFuture captures enum value "future"
	PortalSessionStatusFuture string = "future"

	// PortalSessionStatusInTrial captures enum value "in_trial"
	PortalSessionStatusInTrial string = "in_trial"

	// PortalSessionStatusActive captures enum value "active"
	PortalSessionStatusActive string = "active"

	// PortalSessionStatusNonRenewing captures enum value "non_renewing"
	PortalSessionStatusNonRenewing string = "non_renewing"

	// PortalSessionStatusCancelled captures enum value "cancelled"
	PortalSessionStatusCancelled string = "cancelled"
)

// prop value enum
func (m *PortalSession) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, portalSessionTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PortalSession) validateStatus(formats strfmt.Registry) error {

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
func (m *PortalSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortalSession) UnmarshalBinary(b []byte) error {
	var res PortalSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
