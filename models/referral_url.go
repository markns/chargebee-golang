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

// ReferralURL referral Url
// swagger:model ReferralUrl
type ReferralURL struct {

	// created at
	CreatedAt int64 `json:"created_at,omitempty"`

	// external customer id
	ExternalCustomerID string `json:"external_customer_id,omitempty"`

	// referral account id
	ReferralAccountID string `json:"referral_account_id,omitempty"`

	// referral campaign id
	ReferralCampaignID string `json:"referral_campaign_id,omitempty"`

	// referral external campaign id
	ReferralExternalCampaignID string `json:"referral_external_campaign_id,omitempty"`

	// referral sharing url
	ReferralSharingURL string `json:"referral_sharing_url,omitempty"`

	// referral system
	ReferralSystem string `json:"referral_system,omitempty"`

	// updated at
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

// Validate validates this referral Url
func (m *ReferralURL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReferralSystem(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var referralUrlTypeReferralSystemPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["referral_candy","referral_saasquatch","friendbuy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		referralUrlTypeReferralSystemPropEnum = append(referralUrlTypeReferralSystemPropEnum, v)
	}
}

const (

	// ReferralURLReferralSystemReferralCandy captures enum value "referral_candy"
	ReferralURLReferralSystemReferralCandy string = "referral_candy"

	// ReferralURLReferralSystemReferralSaasquatch captures enum value "referral_saasquatch"
	ReferralURLReferralSystemReferralSaasquatch string = "referral_saasquatch"

	// ReferralURLReferralSystemFriendbuy captures enum value "friendbuy"
	ReferralURLReferralSystemFriendbuy string = "friendbuy"
)

// prop value enum
func (m *ReferralURL) validateReferralSystemEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, referralUrlTypeReferralSystemPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReferralURL) validateReferralSystem(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferralSystem) { // not required
		return nil
	}

	// value enum
	if err := m.validateReferralSystemEnum("referral_system", "body", m.ReferralSystem); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReferralURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReferralURL) UnmarshalBinary(b []byte) error {
	var res ReferralURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
