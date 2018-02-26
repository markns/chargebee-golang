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

// NewAcknowledgeHostedPageParams creates a new AcknowledgeHostedPageParams object
// with the default values initialized.
func NewAcknowledgeHostedPageParams() *AcknowledgeHostedPageParams {
	var ()
	return &AcknowledgeHostedPageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAcknowledgeHostedPageParamsWithTimeout creates a new AcknowledgeHostedPageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAcknowledgeHostedPageParamsWithTimeout(timeout time.Duration) *AcknowledgeHostedPageParams {
	var ()
	return &AcknowledgeHostedPageParams{

		timeout: timeout,
	}
}

// NewAcknowledgeHostedPageParamsWithContext creates a new AcknowledgeHostedPageParams object
// with the default values initialized, and the ability to set a context for a request
func NewAcknowledgeHostedPageParamsWithContext(ctx context.Context) *AcknowledgeHostedPageParams {
	var ()
	return &AcknowledgeHostedPageParams{

		Context: ctx,
	}
}

// NewAcknowledgeHostedPageParamsWithHTTPClient creates a new AcknowledgeHostedPageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAcknowledgeHostedPageParamsWithHTTPClient(client *http.Client) *AcknowledgeHostedPageParams {
	var ()
	return &AcknowledgeHostedPageParams{
		HTTPClient: client,
	}
}

/*AcknowledgeHostedPageParams contains all the parameters to send to the API endpoint
for the acknowledge hosted page operation typically these are written to a http.Request
*/
type AcknowledgeHostedPageParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the acknowledge hosted page params
func (o *AcknowledgeHostedPageParams) WithTimeout(timeout time.Duration) *AcknowledgeHostedPageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the acknowledge hosted page params
func (o *AcknowledgeHostedPageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the acknowledge hosted page params
func (o *AcknowledgeHostedPageParams) WithContext(ctx context.Context) *AcknowledgeHostedPageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the acknowledge hosted page params
func (o *AcknowledgeHostedPageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the acknowledge hosted page params
func (o *AcknowledgeHostedPageParams) WithHTTPClient(client *http.Client) *AcknowledgeHostedPageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the acknowledge hosted page params
func (o *AcknowledgeHostedPageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the acknowledge hosted page params
func (o *AcknowledgeHostedPageParams) WithID(id string) *AcknowledgeHostedPageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the acknowledge hosted page params
func (o *AcknowledgeHostedPageParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AcknowledgeHostedPageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
