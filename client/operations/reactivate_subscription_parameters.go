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

// NewReactivateSubscriptionParams creates a new ReactivateSubscriptionParams object
// with the default values initialized.
func NewReactivateSubscriptionParams() *ReactivateSubscriptionParams {
	var ()
	return &ReactivateSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReactivateSubscriptionParamsWithTimeout creates a new ReactivateSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReactivateSubscriptionParamsWithTimeout(timeout time.Duration) *ReactivateSubscriptionParams {
	var ()
	return &ReactivateSubscriptionParams{

		timeout: timeout,
	}
}

// NewReactivateSubscriptionParamsWithContext creates a new ReactivateSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewReactivateSubscriptionParamsWithContext(ctx context.Context) *ReactivateSubscriptionParams {
	var ()
	return &ReactivateSubscriptionParams{

		Context: ctx,
	}
}

// NewReactivateSubscriptionParamsWithHTTPClient creates a new ReactivateSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReactivateSubscriptionParamsWithHTTPClient(client *http.Client) *ReactivateSubscriptionParams {
	var ()
	return &ReactivateSubscriptionParams{
		HTTPClient: client,
	}
}

/*ReactivateSubscriptionParams contains all the parameters to send to the API endpoint
for the reactivate subscription operation typically these are written to a http.Request
*/
type ReactivateSubscriptionParams struct {

	/*SubscriptionReactivateRequest*/
	SubscriptionReactivateRequest *models.SubscriptionReactivateRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reactivate subscription params
func (o *ReactivateSubscriptionParams) WithTimeout(timeout time.Duration) *ReactivateSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reactivate subscription params
func (o *ReactivateSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reactivate subscription params
func (o *ReactivateSubscriptionParams) WithContext(ctx context.Context) *ReactivateSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reactivate subscription params
func (o *ReactivateSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reactivate subscription params
func (o *ReactivateSubscriptionParams) WithHTTPClient(client *http.Client) *ReactivateSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reactivate subscription params
func (o *ReactivateSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubscriptionReactivateRequest adds the subscriptionReactivateRequest to the reactivate subscription params
func (o *ReactivateSubscriptionParams) WithSubscriptionReactivateRequest(subscriptionReactivateRequest *models.SubscriptionReactivateRequest) *ReactivateSubscriptionParams {
	o.SetSubscriptionReactivateRequest(subscriptionReactivateRequest)
	return o
}

// SetSubscriptionReactivateRequest adds the subscriptionReactivateRequest to the reactivate subscription params
func (o *ReactivateSubscriptionParams) SetSubscriptionReactivateRequest(subscriptionReactivateRequest *models.SubscriptionReactivateRequest) {
	o.SubscriptionReactivateRequest = subscriptionReactivateRequest
}

// WithID adds the id to the reactivate subscription params
func (o *ReactivateSubscriptionParams) WithID(id string) *ReactivateSubscriptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the reactivate subscription params
func (o *ReactivateSubscriptionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReactivateSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SubscriptionReactivateRequest != nil {
		if err := r.SetBodyParam(o.SubscriptionReactivateRequest); err != nil {
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
