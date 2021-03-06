// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSubscriptionsForCustomerSubscriptionParams creates a new SubscriptionsForCustomerSubscriptionParams object
// with the default values initialized.
func NewSubscriptionsForCustomerSubscriptionParams() *SubscriptionsForCustomerSubscriptionParams {
	var ()
	return &SubscriptionsForCustomerSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionsForCustomerSubscriptionParamsWithTimeout creates a new SubscriptionsForCustomerSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubscriptionsForCustomerSubscriptionParamsWithTimeout(timeout time.Duration) *SubscriptionsForCustomerSubscriptionParams {
	var ()
	return &SubscriptionsForCustomerSubscriptionParams{

		timeout: timeout,
	}
}

// NewSubscriptionsForCustomerSubscriptionParamsWithContext creates a new SubscriptionsForCustomerSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubscriptionsForCustomerSubscriptionParamsWithContext(ctx context.Context) *SubscriptionsForCustomerSubscriptionParams {
	var ()
	return &SubscriptionsForCustomerSubscriptionParams{

		Context: ctx,
	}
}

// NewSubscriptionsForCustomerSubscriptionParamsWithHTTPClient creates a new SubscriptionsForCustomerSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubscriptionsForCustomerSubscriptionParamsWithHTTPClient(client *http.Client) *SubscriptionsForCustomerSubscriptionParams {
	var ()
	return &SubscriptionsForCustomerSubscriptionParams{
		HTTPClient: client,
	}
}

/*SubscriptionsForCustomerSubscriptionParams contains all the parameters to send to the API endpoint
for the subscriptions for customer subscription operation typically these are written to a http.Request
*/
type SubscriptionsForCustomerSubscriptionParams struct {

	/*ID*/
	ID string
	/*Limit*/
	Limit *int32
	/*Offset*/
	Offset *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the subscriptions for customer subscription params
func (o *SubscriptionsForCustomerSubscriptionParams) WithTimeout(timeout time.Duration) *SubscriptionsForCustomerSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscriptions for customer subscription params
func (o *SubscriptionsForCustomerSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscriptions for customer subscription params
func (o *SubscriptionsForCustomerSubscriptionParams) WithContext(ctx context.Context) *SubscriptionsForCustomerSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscriptions for customer subscription params
func (o *SubscriptionsForCustomerSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscriptions for customer subscription params
func (o *SubscriptionsForCustomerSubscriptionParams) WithHTTPClient(client *http.Client) *SubscriptionsForCustomerSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscriptions for customer subscription params
func (o *SubscriptionsForCustomerSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the subscriptions for customer subscription params
func (o *SubscriptionsForCustomerSubscriptionParams) WithID(id string) *SubscriptionsForCustomerSubscriptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the subscriptions for customer subscription params
func (o *SubscriptionsForCustomerSubscriptionParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the subscriptions for customer subscription params
func (o *SubscriptionsForCustomerSubscriptionParams) WithLimit(limit *int32) *SubscriptionsForCustomerSubscriptionParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the subscriptions for customer subscription params
func (o *SubscriptionsForCustomerSubscriptionParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the subscriptions for customer subscription params
func (o *SubscriptionsForCustomerSubscriptionParams) WithOffset(offset *string) *SubscriptionsForCustomerSubscriptionParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the subscriptions for customer subscription params
func (o *SubscriptionsForCustomerSubscriptionParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionsForCustomerSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
