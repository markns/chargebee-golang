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

// NewUpdateAddonParams creates a new UpdateAddonParams object
// with the default values initialized.
func NewUpdateAddonParams() *UpdateAddonParams {
	var ()
	return &UpdateAddonParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAddonParamsWithTimeout creates a new UpdateAddonParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAddonParamsWithTimeout(timeout time.Duration) *UpdateAddonParams {
	var ()
	return &UpdateAddonParams{

		timeout: timeout,
	}
}

// NewUpdateAddonParamsWithContext creates a new UpdateAddonParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAddonParamsWithContext(ctx context.Context) *UpdateAddonParams {
	var ()
	return &UpdateAddonParams{

		Context: ctx,
	}
}

// NewUpdateAddonParamsWithHTTPClient creates a new UpdateAddonParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAddonParamsWithHTTPClient(client *http.Client) *UpdateAddonParams {
	var ()
	return &UpdateAddonParams{
		HTTPClient: client,
	}
}

/*UpdateAddonParams contains all the parameters to send to the API endpoint
for the update addon operation typically these are written to a http.Request
*/
type UpdateAddonParams struct {

	/*AddonUpdateRequest*/
	AddonUpdateRequest *models.AddonUpdateRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update addon params
func (o *UpdateAddonParams) WithTimeout(timeout time.Duration) *UpdateAddonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update addon params
func (o *UpdateAddonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update addon params
func (o *UpdateAddonParams) WithContext(ctx context.Context) *UpdateAddonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update addon params
func (o *UpdateAddonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update addon params
func (o *UpdateAddonParams) WithHTTPClient(client *http.Client) *UpdateAddonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update addon params
func (o *UpdateAddonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddonUpdateRequest adds the addonUpdateRequest to the update addon params
func (o *UpdateAddonParams) WithAddonUpdateRequest(addonUpdateRequest *models.AddonUpdateRequest) *UpdateAddonParams {
	o.SetAddonUpdateRequest(addonUpdateRequest)
	return o
}

// SetAddonUpdateRequest adds the addonUpdateRequest to the update addon params
func (o *UpdateAddonParams) SetAddonUpdateRequest(addonUpdateRequest *models.AddonUpdateRequest) {
	o.AddonUpdateRequest = addonUpdateRequest
}

// WithID adds the id to the update addon params
func (o *UpdateAddonParams) WithID(id string) *UpdateAddonParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update addon params
func (o *UpdateAddonParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAddonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AddonUpdateRequest == nil {
		o.AddonUpdateRequest = new(models.AddonUpdateRequest)
	}

	if err := r.SetBodyParam(o.AddonUpdateRequest); err != nil {
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