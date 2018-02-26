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

// NewRetrieveCustomerParams creates a new RetrieveCustomerParams object
// with the default values initialized.
func NewRetrieveCustomerParams() *RetrieveCustomerParams {
	var ()
	return &RetrieveCustomerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveCustomerParamsWithTimeout creates a new RetrieveCustomerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveCustomerParamsWithTimeout(timeout time.Duration) *RetrieveCustomerParams {
	var ()
	return &RetrieveCustomerParams{

		timeout: timeout,
	}
}

// NewRetrieveCustomerParamsWithContext creates a new RetrieveCustomerParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveCustomerParamsWithContext(ctx context.Context) *RetrieveCustomerParams {
	var ()
	return &RetrieveCustomerParams{

		Context: ctx,
	}
}

// NewRetrieveCustomerParamsWithHTTPClient creates a new RetrieveCustomerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveCustomerParamsWithHTTPClient(client *http.Client) *RetrieveCustomerParams {
	var ()
	return &RetrieveCustomerParams{
		HTTPClient: client,
	}
}

/*RetrieveCustomerParams contains all the parameters to send to the API endpoint
for the retrieve customer operation typically these are written to a http.Request
*/
type RetrieveCustomerParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve customer params
func (o *RetrieveCustomerParams) WithTimeout(timeout time.Duration) *RetrieveCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve customer params
func (o *RetrieveCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve customer params
func (o *RetrieveCustomerParams) WithContext(ctx context.Context) *RetrieveCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve customer params
func (o *RetrieveCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve customer params
func (o *RetrieveCustomerParams) WithHTTPClient(client *http.Client) *RetrieveCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve customer params
func (o *RetrieveCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the retrieve customer params
func (o *RetrieveCustomerParams) WithID(id string) *RetrieveCustomerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the retrieve customer params
func (o *RetrieveCustomerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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