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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/markns/chargebee-golang/models"
)

// NewCreateForCustomerSubscriptionParams creates a new CreateForCustomerSubscriptionParams object
// with the default values initialized.
func NewCreateForCustomerSubscriptionParams() *CreateForCustomerSubscriptionParams {
	var ()
	return &CreateForCustomerSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateForCustomerSubscriptionParamsWithTimeout creates a new CreateForCustomerSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateForCustomerSubscriptionParamsWithTimeout(timeout time.Duration) *CreateForCustomerSubscriptionParams {
	var ()
	return &CreateForCustomerSubscriptionParams{

		timeout: timeout,
	}
}

// NewCreateForCustomerSubscriptionParamsWithContext creates a new CreateForCustomerSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateForCustomerSubscriptionParamsWithContext(ctx context.Context) *CreateForCustomerSubscriptionParams {
	var ()
	return &CreateForCustomerSubscriptionParams{

		Context: ctx,
	}
}

// NewCreateForCustomerSubscriptionParamsWithHTTPClient creates a new CreateForCustomerSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateForCustomerSubscriptionParamsWithHTTPClient(client *http.Client) *CreateForCustomerSubscriptionParams {
	var ()
	return &CreateForCustomerSubscriptionParams{
		HTTPClient: client,
	}
}

/*CreateForCustomerSubscriptionParams contains all the parameters to send to the API endpoint
for the create for customer subscription operation typically these are written to a http.Request
*/
type CreateForCustomerSubscriptionParams struct {

	/*SubscriptionCreateForCustomerRequest*/
	SubscriptionCreateForCustomerRequest *models.SubscriptionCreateForCustomerRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create for customer subscription params
func (o *CreateForCustomerSubscriptionParams) WithTimeout(timeout time.Duration) *CreateForCustomerSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create for customer subscription params
func (o *CreateForCustomerSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create for customer subscription params
func (o *CreateForCustomerSubscriptionParams) WithContext(ctx context.Context) *CreateForCustomerSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create for customer subscription params
func (o *CreateForCustomerSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create for customer subscription params
func (o *CreateForCustomerSubscriptionParams) WithHTTPClient(client *http.Client) *CreateForCustomerSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create for customer subscription params
func (o *CreateForCustomerSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubscriptionCreateForCustomerRequest adds the subscriptionCreateForCustomerRequest to the create for customer subscription params
func (o *CreateForCustomerSubscriptionParams) WithSubscriptionCreateForCustomerRequest(subscriptionCreateForCustomerRequest *models.SubscriptionCreateForCustomerRequest) *CreateForCustomerSubscriptionParams {
	o.SetSubscriptionCreateForCustomerRequest(subscriptionCreateForCustomerRequest)
	return o
}

// SetSubscriptionCreateForCustomerRequest adds the subscriptionCreateForCustomerRequest to the create for customer subscription params
func (o *CreateForCustomerSubscriptionParams) SetSubscriptionCreateForCustomerRequest(subscriptionCreateForCustomerRequest *models.SubscriptionCreateForCustomerRequest) {
	o.SubscriptionCreateForCustomerRequest = subscriptionCreateForCustomerRequest
}

// WithID adds the id to the create for customer subscription params
func (o *CreateForCustomerSubscriptionParams) WithID(id string) *CreateForCustomerSubscriptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create for customer subscription params
func (o *CreateForCustomerSubscriptionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateForCustomerSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SubscriptionCreateForCustomerRequest != nil {
		if err := r.SetBodyParam(o.SubscriptionCreateForCustomerRequest); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
