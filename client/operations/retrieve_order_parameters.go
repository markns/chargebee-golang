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

// NewRetrieveOrderParams creates a new RetrieveOrderParams object
// with the default values initialized.
func NewRetrieveOrderParams() *RetrieveOrderParams {
	var ()
	return &RetrieveOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveOrderParamsWithTimeout creates a new RetrieveOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveOrderParamsWithTimeout(timeout time.Duration) *RetrieveOrderParams {
	var ()
	return &RetrieveOrderParams{

		timeout: timeout,
	}
}

// NewRetrieveOrderParamsWithContext creates a new RetrieveOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveOrderParamsWithContext(ctx context.Context) *RetrieveOrderParams {
	var ()
	return &RetrieveOrderParams{

		Context: ctx,
	}
}

// NewRetrieveOrderParamsWithHTTPClient creates a new RetrieveOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveOrderParamsWithHTTPClient(client *http.Client) *RetrieveOrderParams {
	var ()
	return &RetrieveOrderParams{
		HTTPClient: client,
	}
}

/*RetrieveOrderParams contains all the parameters to send to the API endpoint
for the retrieve order operation typically these are written to a http.Request
*/
type RetrieveOrderParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve order params
func (o *RetrieveOrderParams) WithTimeout(timeout time.Duration) *RetrieveOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve order params
func (o *RetrieveOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve order params
func (o *RetrieveOrderParams) WithContext(ctx context.Context) *RetrieveOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve order params
func (o *RetrieveOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve order params
func (o *RetrieveOrderParams) WithHTTPClient(client *http.Client) *RetrieveOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve order params
func (o *RetrieveOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the retrieve order params
func (o *RetrieveOrderParams) WithID(id string) *RetrieveOrderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the retrieve order params
func (o *RetrieveOrderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
