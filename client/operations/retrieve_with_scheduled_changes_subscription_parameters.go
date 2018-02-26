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
)

// NewRetrieveWithScheduledChangesSubscriptionParams creates a new RetrieveWithScheduledChangesSubscriptionParams object
// with the default values initialized.
func NewRetrieveWithScheduledChangesSubscriptionParams() *RetrieveWithScheduledChangesSubscriptionParams {
	var ()
	return &RetrieveWithScheduledChangesSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveWithScheduledChangesSubscriptionParamsWithTimeout creates a new RetrieveWithScheduledChangesSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveWithScheduledChangesSubscriptionParamsWithTimeout(timeout time.Duration) *RetrieveWithScheduledChangesSubscriptionParams {
	var ()
	return &RetrieveWithScheduledChangesSubscriptionParams{

		timeout: timeout,
	}
}

// NewRetrieveWithScheduledChangesSubscriptionParamsWithContext creates a new RetrieveWithScheduledChangesSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveWithScheduledChangesSubscriptionParamsWithContext(ctx context.Context) *RetrieveWithScheduledChangesSubscriptionParams {
	var ()
	return &RetrieveWithScheduledChangesSubscriptionParams{

		Context: ctx,
	}
}

// NewRetrieveWithScheduledChangesSubscriptionParamsWithHTTPClient creates a new RetrieveWithScheduledChangesSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveWithScheduledChangesSubscriptionParamsWithHTTPClient(client *http.Client) *RetrieveWithScheduledChangesSubscriptionParams {
	var ()
	return &RetrieveWithScheduledChangesSubscriptionParams{
		HTTPClient: client,
	}
}

/*RetrieveWithScheduledChangesSubscriptionParams contains all the parameters to send to the API endpoint
for the retrieve with scheduled changes subscription operation typically these are written to a http.Request
*/
type RetrieveWithScheduledChangesSubscriptionParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve with scheduled changes subscription params
func (o *RetrieveWithScheduledChangesSubscriptionParams) WithTimeout(timeout time.Duration) *RetrieveWithScheduledChangesSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve with scheduled changes subscription params
func (o *RetrieveWithScheduledChangesSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve with scheduled changes subscription params
func (o *RetrieveWithScheduledChangesSubscriptionParams) WithContext(ctx context.Context) *RetrieveWithScheduledChangesSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve with scheduled changes subscription params
func (o *RetrieveWithScheduledChangesSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve with scheduled changes subscription params
func (o *RetrieveWithScheduledChangesSubscriptionParams) WithHTTPClient(client *http.Client) *RetrieveWithScheduledChangesSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve with scheduled changes subscription params
func (o *RetrieveWithScheduledChangesSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the retrieve with scheduled changes subscription params
func (o *RetrieveWithScheduledChangesSubscriptionParams) WithID(id string) *RetrieveWithScheduledChangesSubscriptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the retrieve with scheduled changes subscription params
func (o *RetrieveWithScheduledChangesSubscriptionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveWithScheduledChangesSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}