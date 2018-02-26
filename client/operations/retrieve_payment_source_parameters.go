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

// NewRetrievePaymentSourceParams creates a new RetrievePaymentSourceParams object
// with the default values initialized.
func NewRetrievePaymentSourceParams() *RetrievePaymentSourceParams {
	var ()
	return &RetrievePaymentSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrievePaymentSourceParamsWithTimeout creates a new RetrievePaymentSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrievePaymentSourceParamsWithTimeout(timeout time.Duration) *RetrievePaymentSourceParams {
	var ()
	return &RetrievePaymentSourceParams{

		timeout: timeout,
	}
}

// NewRetrievePaymentSourceParamsWithContext creates a new RetrievePaymentSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrievePaymentSourceParamsWithContext(ctx context.Context) *RetrievePaymentSourceParams {
	var ()
	return &RetrievePaymentSourceParams{

		Context: ctx,
	}
}

// NewRetrievePaymentSourceParamsWithHTTPClient creates a new RetrievePaymentSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrievePaymentSourceParamsWithHTTPClient(client *http.Client) *RetrievePaymentSourceParams {
	var ()
	return &RetrievePaymentSourceParams{
		HTTPClient: client,
	}
}

/*RetrievePaymentSourceParams contains all the parameters to send to the API endpoint
for the retrieve payment source operation typically these are written to a http.Request
*/
type RetrievePaymentSourceParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve payment source params
func (o *RetrievePaymentSourceParams) WithTimeout(timeout time.Duration) *RetrievePaymentSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve payment source params
func (o *RetrievePaymentSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve payment source params
func (o *RetrievePaymentSourceParams) WithContext(ctx context.Context) *RetrievePaymentSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve payment source params
func (o *RetrievePaymentSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve payment source params
func (o *RetrievePaymentSourceParams) WithHTTPClient(client *http.Client) *RetrievePaymentSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve payment source params
func (o *RetrievePaymentSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the retrieve payment source params
func (o *RetrievePaymentSourceParams) WithID(id string) *RetrievePaymentSourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the retrieve payment source params
func (o *RetrievePaymentSourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RetrievePaymentSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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