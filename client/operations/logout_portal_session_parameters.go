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

// NewLogoutPortalSessionParams creates a new LogoutPortalSessionParams object
// with the default values initialized.
func NewLogoutPortalSessionParams() *LogoutPortalSessionParams {
	var ()
	return &LogoutPortalSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLogoutPortalSessionParamsWithTimeout creates a new LogoutPortalSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLogoutPortalSessionParamsWithTimeout(timeout time.Duration) *LogoutPortalSessionParams {
	var ()
	return &LogoutPortalSessionParams{

		timeout: timeout,
	}
}

// NewLogoutPortalSessionParamsWithContext creates a new LogoutPortalSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewLogoutPortalSessionParamsWithContext(ctx context.Context) *LogoutPortalSessionParams {
	var ()
	return &LogoutPortalSessionParams{

		Context: ctx,
	}
}

// NewLogoutPortalSessionParamsWithHTTPClient creates a new LogoutPortalSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLogoutPortalSessionParamsWithHTTPClient(client *http.Client) *LogoutPortalSessionParams {
	var ()
	return &LogoutPortalSessionParams{
		HTTPClient: client,
	}
}

/*LogoutPortalSessionParams contains all the parameters to send to the API endpoint
for the logout portal session operation typically these are written to a http.Request
*/
type LogoutPortalSessionParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the logout portal session params
func (o *LogoutPortalSessionParams) WithTimeout(timeout time.Duration) *LogoutPortalSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the logout portal session params
func (o *LogoutPortalSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the logout portal session params
func (o *LogoutPortalSessionParams) WithContext(ctx context.Context) *LogoutPortalSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the logout portal session params
func (o *LogoutPortalSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the logout portal session params
func (o *LogoutPortalSessionParams) WithHTTPClient(client *http.Client) *LogoutPortalSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the logout portal session params
func (o *LogoutPortalSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the logout portal session params
func (o *LogoutPortalSessionParams) WithID(id string) *LogoutPortalSessionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the logout portal session params
func (o *LogoutPortalSessionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LogoutPortalSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
