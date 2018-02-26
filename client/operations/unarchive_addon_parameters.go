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

// NewUnarchiveAddonParams creates a new UnarchiveAddonParams object
// with the default values initialized.
func NewUnarchiveAddonParams() *UnarchiveAddonParams {
	var ()
	return &UnarchiveAddonParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUnarchiveAddonParamsWithTimeout creates a new UnarchiveAddonParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnarchiveAddonParamsWithTimeout(timeout time.Duration) *UnarchiveAddonParams {
	var ()
	return &UnarchiveAddonParams{

		timeout: timeout,
	}
}

// NewUnarchiveAddonParamsWithContext creates a new UnarchiveAddonParams object
// with the default values initialized, and the ability to set a context for a request
func NewUnarchiveAddonParamsWithContext(ctx context.Context) *UnarchiveAddonParams {
	var ()
	return &UnarchiveAddonParams{

		Context: ctx,
	}
}

// NewUnarchiveAddonParamsWithHTTPClient creates a new UnarchiveAddonParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUnarchiveAddonParamsWithHTTPClient(client *http.Client) *UnarchiveAddonParams {
	var ()
	return &UnarchiveAddonParams{
		HTTPClient: client,
	}
}

/*UnarchiveAddonParams contains all the parameters to send to the API endpoint
for the unarchive addon operation typically these are written to a http.Request
*/
type UnarchiveAddonParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the unarchive addon params
func (o *UnarchiveAddonParams) WithTimeout(timeout time.Duration) *UnarchiveAddonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unarchive addon params
func (o *UnarchiveAddonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unarchive addon params
func (o *UnarchiveAddonParams) WithContext(ctx context.Context) *UnarchiveAddonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unarchive addon params
func (o *UnarchiveAddonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unarchive addon params
func (o *UnarchiveAddonParams) WithHTTPClient(client *http.Client) *UnarchiveAddonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unarchive addon params
func (o *UnarchiveAddonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the unarchive addon params
func (o *UnarchiveAddonParams) WithID(id string) *UnarchiveAddonParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the unarchive addon params
func (o *UnarchiveAddonParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UnarchiveAddonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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