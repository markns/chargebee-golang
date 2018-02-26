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

// NewRetrievePortalSessionParams creates a new RetrievePortalSessionParams object
// with the default values initialized.
func NewRetrievePortalSessionParams() *RetrievePortalSessionParams {
	var ()
	return &RetrievePortalSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrievePortalSessionParamsWithTimeout creates a new RetrievePortalSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrievePortalSessionParamsWithTimeout(timeout time.Duration) *RetrievePortalSessionParams {
	var ()
	return &RetrievePortalSessionParams{

		timeout: timeout,
	}
}

// NewRetrievePortalSessionParamsWithContext creates a new RetrievePortalSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrievePortalSessionParamsWithContext(ctx context.Context) *RetrievePortalSessionParams {
	var ()
	return &RetrievePortalSessionParams{

		Context: ctx,
	}
}

// NewRetrievePortalSessionParamsWithHTTPClient creates a new RetrievePortalSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrievePortalSessionParamsWithHTTPClient(client *http.Client) *RetrievePortalSessionParams {
	var ()
	return &RetrievePortalSessionParams{
		HTTPClient: client,
	}
}

/*RetrievePortalSessionParams contains all the parameters to send to the API endpoint
for the retrieve portal session operation typically these are written to a http.Request
*/
type RetrievePortalSessionParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve portal session params
func (o *RetrievePortalSessionParams) WithTimeout(timeout time.Duration) *RetrievePortalSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve portal session params
func (o *RetrievePortalSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve portal session params
func (o *RetrievePortalSessionParams) WithContext(ctx context.Context) *RetrievePortalSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve portal session params
func (o *RetrievePortalSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve portal session params
func (o *RetrievePortalSessionParams) WithHTTPClient(client *http.Client) *RetrievePortalSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve portal session params
func (o *RetrievePortalSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the retrieve portal session params
func (o *RetrievePortalSessionParams) WithID(id string) *RetrievePortalSessionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the retrieve portal session params
func (o *RetrievePortalSessionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RetrievePortalSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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