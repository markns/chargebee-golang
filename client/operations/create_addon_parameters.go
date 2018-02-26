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

// NewCreateAddonParams creates a new CreateAddonParams object
// with the default values initialized.
func NewCreateAddonParams() *CreateAddonParams {
	var ()
	return &CreateAddonParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAddonParamsWithTimeout creates a new CreateAddonParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAddonParamsWithTimeout(timeout time.Duration) *CreateAddonParams {
	var ()
	return &CreateAddonParams{

		timeout: timeout,
	}
}

// NewCreateAddonParamsWithContext creates a new CreateAddonParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAddonParamsWithContext(ctx context.Context) *CreateAddonParams {
	var ()
	return &CreateAddonParams{

		Context: ctx,
	}
}

// NewCreateAddonParamsWithHTTPClient creates a new CreateAddonParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAddonParamsWithHTTPClient(client *http.Client) *CreateAddonParams {
	var ()
	return &CreateAddonParams{
		HTTPClient: client,
	}
}

/*CreateAddonParams contains all the parameters to send to the API endpoint
for the create addon operation typically these are written to a http.Request
*/
type CreateAddonParams struct {

	/*AddonCreateRequest*/
	AddonCreateRequest *models.AddonCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create addon params
func (o *CreateAddonParams) WithTimeout(timeout time.Duration) *CreateAddonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create addon params
func (o *CreateAddonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create addon params
func (o *CreateAddonParams) WithContext(ctx context.Context) *CreateAddonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create addon params
func (o *CreateAddonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create addon params
func (o *CreateAddonParams) WithHTTPClient(client *http.Client) *CreateAddonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create addon params
func (o *CreateAddonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddonCreateRequest adds the addonCreateRequest to the create addon params
func (o *CreateAddonParams) WithAddonCreateRequest(addonCreateRequest *models.AddonCreateRequest) *CreateAddonParams {
	o.SetAddonCreateRequest(addonCreateRequest)
	return o
}

// SetAddonCreateRequest adds the addonCreateRequest to the create addon params
func (o *CreateAddonParams) SetAddonCreateRequest(addonCreateRequest *models.AddonCreateRequest) {
	o.AddonCreateRequest = addonCreateRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAddonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AddonCreateRequest == nil {
		o.AddonCreateRequest = new(models.AddonCreateRequest)
	}

	if err := r.SetBodyParam(o.AddonCreateRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
