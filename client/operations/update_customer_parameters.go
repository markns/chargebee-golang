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

// NewUpdateCustomerParams creates a new UpdateCustomerParams object
// with the default values initialized.
func NewUpdateCustomerParams() *UpdateCustomerParams {
	var ()
	return &UpdateCustomerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCustomerParamsWithTimeout creates a new UpdateCustomerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCustomerParamsWithTimeout(timeout time.Duration) *UpdateCustomerParams {
	var ()
	return &UpdateCustomerParams{

		timeout: timeout,
	}
}

// NewUpdateCustomerParamsWithContext creates a new UpdateCustomerParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCustomerParamsWithContext(ctx context.Context) *UpdateCustomerParams {
	var ()
	return &UpdateCustomerParams{

		Context: ctx,
	}
}

// NewUpdateCustomerParamsWithHTTPClient creates a new UpdateCustomerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCustomerParamsWithHTTPClient(client *http.Client) *UpdateCustomerParams {
	var ()
	return &UpdateCustomerParams{
		HTTPClient: client,
	}
}

/*UpdateCustomerParams contains all the parameters to send to the API endpoint
for the update customer operation typically these are written to a http.Request
*/
type UpdateCustomerParams struct {

	/*CustomerUpdateRequest*/
	CustomerUpdateRequest *models.CustomerUpdateRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update customer params
func (o *UpdateCustomerParams) WithTimeout(timeout time.Duration) *UpdateCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update customer params
func (o *UpdateCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update customer params
func (o *UpdateCustomerParams) WithContext(ctx context.Context) *UpdateCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update customer params
func (o *UpdateCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update customer params
func (o *UpdateCustomerParams) WithHTTPClient(client *http.Client) *UpdateCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update customer params
func (o *UpdateCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerUpdateRequest adds the customerUpdateRequest to the update customer params
func (o *UpdateCustomerParams) WithCustomerUpdateRequest(customerUpdateRequest *models.CustomerUpdateRequest) *UpdateCustomerParams {
	o.SetCustomerUpdateRequest(customerUpdateRequest)
	return o
}

// SetCustomerUpdateRequest adds the customerUpdateRequest to the update customer params
func (o *UpdateCustomerParams) SetCustomerUpdateRequest(customerUpdateRequest *models.CustomerUpdateRequest) {
	o.CustomerUpdateRequest = customerUpdateRequest
}

// WithID adds the id to the update customer params
func (o *UpdateCustomerParams) WithID(id string) *UpdateCustomerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update customer params
func (o *UpdateCustomerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CustomerUpdateRequest == nil {
		o.CustomerUpdateRequest = new(models.CustomerUpdateRequest)
	}

	if err := r.SetBodyParam(o.CustomerUpdateRequest); err != nil {
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