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

// NewRemoveScheduledChangesSubscriptionParams creates a new RemoveScheduledChangesSubscriptionParams object
// with the default values initialized.
func NewRemoveScheduledChangesSubscriptionParams() *RemoveScheduledChangesSubscriptionParams {
	var ()
	return &RemoveScheduledChangesSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveScheduledChangesSubscriptionParamsWithTimeout creates a new RemoveScheduledChangesSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveScheduledChangesSubscriptionParamsWithTimeout(timeout time.Duration) *RemoveScheduledChangesSubscriptionParams {
	var ()
	return &RemoveScheduledChangesSubscriptionParams{

		timeout: timeout,
	}
}

// NewRemoveScheduledChangesSubscriptionParamsWithContext creates a new RemoveScheduledChangesSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveScheduledChangesSubscriptionParamsWithContext(ctx context.Context) *RemoveScheduledChangesSubscriptionParams {
	var ()
	return &RemoveScheduledChangesSubscriptionParams{

		Context: ctx,
	}
}

// NewRemoveScheduledChangesSubscriptionParamsWithHTTPClient creates a new RemoveScheduledChangesSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveScheduledChangesSubscriptionParamsWithHTTPClient(client *http.Client) *RemoveScheduledChangesSubscriptionParams {
	var ()
	return &RemoveScheduledChangesSubscriptionParams{
		HTTPClient: client,
	}
}

/*RemoveScheduledChangesSubscriptionParams contains all the parameters to send to the API endpoint
for the remove scheduled changes subscription operation typically these are written to a http.Request
*/
type RemoveScheduledChangesSubscriptionParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove scheduled changes subscription params
func (o *RemoveScheduledChangesSubscriptionParams) WithTimeout(timeout time.Duration) *RemoveScheduledChangesSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove scheduled changes subscription params
func (o *RemoveScheduledChangesSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove scheduled changes subscription params
func (o *RemoveScheduledChangesSubscriptionParams) WithContext(ctx context.Context) *RemoveScheduledChangesSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove scheduled changes subscription params
func (o *RemoveScheduledChangesSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove scheduled changes subscription params
func (o *RemoveScheduledChangesSubscriptionParams) WithHTTPClient(client *http.Client) *RemoveScheduledChangesSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove scheduled changes subscription params
func (o *RemoveScheduledChangesSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the remove scheduled changes subscription params
func (o *RemoveScheduledChangesSubscriptionParams) WithID(id string) *RemoveScheduledChangesSubscriptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove scheduled changes subscription params
func (o *RemoveScheduledChangesSubscriptionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveScheduledChangesSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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