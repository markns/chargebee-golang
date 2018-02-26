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

// NewRetrievePlanParams creates a new RetrievePlanParams object
// with the default values initialized.
func NewRetrievePlanParams() *RetrievePlanParams {
	var ()
	return &RetrievePlanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrievePlanParamsWithTimeout creates a new RetrievePlanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrievePlanParamsWithTimeout(timeout time.Duration) *RetrievePlanParams {
	var ()
	return &RetrievePlanParams{

		timeout: timeout,
	}
}

// NewRetrievePlanParamsWithContext creates a new RetrievePlanParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrievePlanParamsWithContext(ctx context.Context) *RetrievePlanParams {
	var ()
	return &RetrievePlanParams{

		Context: ctx,
	}
}

// NewRetrievePlanParamsWithHTTPClient creates a new RetrievePlanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrievePlanParamsWithHTTPClient(client *http.Client) *RetrievePlanParams {
	var ()
	return &RetrievePlanParams{
		HTTPClient: client,
	}
}

/*RetrievePlanParams contains all the parameters to send to the API endpoint
for the retrieve plan operation typically these are written to a http.Request
*/
type RetrievePlanParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve plan params
func (o *RetrievePlanParams) WithTimeout(timeout time.Duration) *RetrievePlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve plan params
func (o *RetrievePlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve plan params
func (o *RetrievePlanParams) WithContext(ctx context.Context) *RetrievePlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve plan params
func (o *RetrievePlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve plan params
func (o *RetrievePlanParams) WithHTTPClient(client *http.Client) *RetrievePlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve plan params
func (o *RetrievePlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the retrieve plan params
func (o *RetrievePlanParams) WithID(id string) *RetrievePlanParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the retrieve plan params
func (o *RetrievePlanParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RetrievePlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
