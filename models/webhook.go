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

// Webhook webhook
// swagger:model Webhook
type Webhook struct {

	// id
	ID string `json:"id,omitempty"`

	// webhook status
	WebhookStatus string `json:"webhook_status,omitempty"`
}

// Validate validates this webhook
func (m *Webhook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWebhookStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var webhookTypeWebhookStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["not_configured","scheduled","succeeded","re_scheduled","failed","skipped","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webhookTypeWebhookStatusPropEnum = append(webhookTypeWebhookStatusPropEnum, v)
	}
}

const (

	// WebhookWebhookStatusNotConfigured captures enum value "not_configured"
	WebhookWebhookStatusNotConfigured string = "not_configured"

	// WebhookWebhookStatusScheduled captures enum value "scheduled"
	WebhookWebhookStatusScheduled string = "scheduled"

	// WebhookWebhookStatusSucceeded captures enum value "succeeded"
	WebhookWebhookStatusSucceeded string = "succeeded"

	// WebhookWebhookStatusReScheduled captures enum value "re_scheduled"
	WebhookWebhookStatusReScheduled string = "re_scheduled"

	// WebhookWebhookStatusFailed captures enum value "failed"
	WebhookWebhookStatusFailed string = "failed"

	// WebhookWebhookStatusSkipped captures enum value "skipped"
	WebhookWebhookStatusSkipped string = "skipped"

	// WebhookWebhookStatusNotApplicable captures enum value "not_applicable"
	WebhookWebhookStatusNotApplicable string = "not_applicable"
)

// prop value enum
func (m *Webhook) validateWebhookStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, webhookTypeWebhookStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Webhook) validateWebhookStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.WebhookStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateWebhookStatusEnum("webhook_status", "body", m.WebhookStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Webhook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Webhook) UnmarshalBinary(b []byte) error {
	var res Webhook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
