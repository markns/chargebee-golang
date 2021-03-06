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

// NewDeletePaymentSourceParams creates a new DeletePaymentSourceParams object
// with the default values initialized.
func NewDeletePaymentSourceParams() *DeletePaymentSourceParams {
	var ()
	return &DeletePaymentSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePaymentSourceParamsWithTimeout creates a new DeletePaymentSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePaymentSourceParamsWithTimeout(timeout time.Duration) *DeletePaymentSourceParams {
	var ()
	return &DeletePaymentSourceParams{

		timeout: timeout,
	}
}

// NewDeletePaymentSourceParamsWithContext creates a new DeletePaymentSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePaymentSourceParamsWithContext(ctx context.Context) *DeletePaymentSourceParams {
	var ()
	return &DeletePaymentSourceParams{

		Context: ctx,
	}
}

// NewDeletePaymentSourceParamsWithHTTPClient creates a new DeletePaymentSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePaymentSourceParamsWithHTTPClient(client *http.Client) *DeletePaymentSourceParams {
	var ()
	return &DeletePaymentSourceParams{
		HTTPClient: client,
	}
}

/*DeletePaymentSourceParams contains all the parameters to send to the API endpoint
for the delete payment source operation typically these are written to a http.Request
*/
type DeletePaymentSourceParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete payment source params
func (o *DeletePaymentSourceParams) WithTimeout(timeout time.Duration) *DeletePaymentSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete payment source params
func (o *DeletePaymentSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete payment source params
func (o *DeletePaymentSourceParams) WithContext(ctx context.Context) *DeletePaymentSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete payment source params
func (o *DeletePaymentSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete payment source params
func (o *DeletePaymentSourceParams) WithHTTPClient(client *http.Client) *DeletePaymentSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete payment source params
func (o *DeletePaymentSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete payment source params
func (o *DeletePaymentSourceParams) WithID(id string) *DeletePaymentSourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete payment source params
func (o *DeletePaymentSourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePaymentSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
