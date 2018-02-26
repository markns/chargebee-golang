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

	"github.com/markns/chargebee-golang/models"
)

// NewActivatePortalSessionParams creates a new ActivatePortalSessionParams object
// with the default values initialized.
func NewActivatePortalSessionParams() *ActivatePortalSessionParams {
	var ()
	return &ActivatePortalSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewActivatePortalSessionParamsWithTimeout creates a new ActivatePortalSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewActivatePortalSessionParamsWithTimeout(timeout time.Duration) *ActivatePortalSessionParams {
	var ()
	return &ActivatePortalSessionParams{

		timeout: timeout,
	}
}

// NewActivatePortalSessionParamsWithContext creates a new ActivatePortalSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewActivatePortalSessionParamsWithContext(ctx context.Context) *ActivatePortalSessionParams {
	var ()
	return &ActivatePortalSessionParams{

		Context: ctx,
	}
}

// NewActivatePortalSessionParamsWithHTTPClient creates a new ActivatePortalSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewActivatePortalSessionParamsWithHTTPClient(client *http.Client) *ActivatePortalSessionParams {
	var ()
	return &ActivatePortalSessionParams{
		HTTPClient: client,
	}
}

/*ActivatePortalSessionParams contains all the parameters to send to the API endpoint
for the activate portal session operation typically these are written to a http.Request
*/
type ActivatePortalSessionParams struct {

	/*PortalSessionActivateRequest*/
	PortalSessionActivateRequest *models.PortalSessionActivateRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the activate portal session params
func (o *ActivatePortalSessionParams) WithTimeout(timeout time.Duration) *ActivatePortalSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the activate portal session params
func (o *ActivatePortalSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the activate portal session params
func (o *ActivatePortalSessionParams) WithContext(ctx context.Context) *ActivatePortalSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the activate portal session params
func (o *ActivatePortalSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the activate portal session params
func (o *ActivatePortalSessionParams) WithHTTPClient(client *http.Client) *ActivatePortalSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the activate portal session params
func (o *ActivatePortalSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPortalSessionActivateRequest adds the portalSessionActivateRequest to the activate portal session params
func (o *ActivatePortalSessionParams) WithPortalSessionActivateRequest(portalSessionActivateRequest *models.PortalSessionActivateRequest) *ActivatePortalSessionParams {
	o.SetPortalSessionActivateRequest(portalSessionActivateRequest)
	return o
}

// SetPortalSessionActivateRequest adds the portalSessionActivateRequest to the activate portal session params
func (o *ActivatePortalSessionParams) SetPortalSessionActivateRequest(portalSessionActivateRequest *models.PortalSessionActivateRequest) {
	o.PortalSessionActivateRequest = portalSessionActivateRequest
}

// WithID adds the id to the activate portal session params
func (o *ActivatePortalSessionParams) WithID(id string) *ActivatePortalSessionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the activate portal session params
func (o *ActivatePortalSessionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ActivatePortalSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PortalSessionActivateRequest == nil {
		o.PortalSessionActivateRequest = new(models.PortalSessionActivateRequest)
	}

	if err := r.SetBodyParam(o.PortalSessionActivateRequest); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
