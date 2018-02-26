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

// NewDeletePlanParams creates a new DeletePlanParams object
// with the default values initialized.
func NewDeletePlanParams() *DeletePlanParams {
	var ()
	return &DeletePlanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePlanParamsWithTimeout creates a new DeletePlanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePlanParamsWithTimeout(timeout time.Duration) *DeletePlanParams {
	var ()
	return &DeletePlanParams{

		timeout: timeout,
	}
}

// NewDeletePlanParamsWithContext creates a new DeletePlanParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePlanParamsWithContext(ctx context.Context) *DeletePlanParams {
	var ()
	return &DeletePlanParams{

		Context: ctx,
	}
}

// NewDeletePlanParamsWithHTTPClient creates a new DeletePlanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePlanParamsWithHTTPClient(client *http.Client) *DeletePlanParams {
	var ()
	return &DeletePlanParams{
		HTTPClient: client,
	}
}

/*DeletePlanParams contains all the parameters to send to the API endpoint
for the delete plan operation typically these are written to a http.Request
*/
type DeletePlanParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete plan params
func (o *DeletePlanParams) WithTimeout(timeout time.Duration) *DeletePlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete plan params
func (o *DeletePlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete plan params
func (o *DeletePlanParams) WithContext(ctx context.Context) *DeletePlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete plan params
func (o *DeletePlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete plan params
func (o *DeletePlanParams) WithHTTPClient(client *http.Client) *DeletePlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete plan params
func (o *DeletePlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete plan params
func (o *DeletePlanParams) WithID(id string) *DeletePlanParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete plan params
func (o *DeletePlanParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
