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

// NewRetrieveCardParams creates a new RetrieveCardParams object
// with the default values initialized.
func NewRetrieveCardParams() *RetrieveCardParams {
	var ()
	return &RetrieveCardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveCardParamsWithTimeout creates a new RetrieveCardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveCardParamsWithTimeout(timeout time.Duration) *RetrieveCardParams {
	var ()
	return &RetrieveCardParams{

		timeout: timeout,
	}
}

// NewRetrieveCardParamsWithContext creates a new RetrieveCardParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveCardParamsWithContext(ctx context.Context) *RetrieveCardParams {
	var ()
	return &RetrieveCardParams{

		Context: ctx,
	}
}

// NewRetrieveCardParamsWithHTTPClient creates a new RetrieveCardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveCardParamsWithHTTPClient(client *http.Client) *RetrieveCardParams {
	var ()
	return &RetrieveCardParams{
		HTTPClient: client,
	}
}

/*RetrieveCardParams contains all the parameters to send to the API endpoint
for the retrieve card operation typically these are written to a http.Request
*/
type RetrieveCardParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve card params
func (o *RetrieveCardParams) WithTimeout(timeout time.Duration) *RetrieveCardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve card params
func (o *RetrieveCardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve card params
func (o *RetrieveCardParams) WithContext(ctx context.Context) *RetrieveCardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve card params
func (o *RetrieveCardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve card params
func (o *RetrieveCardParams) WithHTTPClient(client *http.Client) *RetrieveCardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve card params
func (o *RetrieveCardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the retrieve card params
func (o *RetrieveCardParams) WithID(id string) *RetrieveCardParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the retrieve card params
func (o *RetrieveCardParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveCardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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