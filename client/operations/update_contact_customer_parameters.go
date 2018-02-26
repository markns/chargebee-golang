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

// NewUpdateContactCustomerParams creates a new UpdateContactCustomerParams object
// with the default values initialized.
func NewUpdateContactCustomerParams() *UpdateContactCustomerParams {
	var ()
	return &UpdateContactCustomerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateContactCustomerParamsWithTimeout creates a new UpdateContactCustomerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateContactCustomerParamsWithTimeout(timeout time.Duration) *UpdateContactCustomerParams {
	var ()
	return &UpdateContactCustomerParams{

		timeout: timeout,
	}
}

// NewUpdateContactCustomerParamsWithContext creates a new UpdateContactCustomerParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateContactCustomerParamsWithContext(ctx context.Context) *UpdateContactCustomerParams {
	var ()
	return &UpdateContactCustomerParams{

		Context: ctx,
	}
}

// NewUpdateContactCustomerParamsWithHTTPClient creates a new UpdateContactCustomerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateContactCustomerParamsWithHTTPClient(client *http.Client) *UpdateContactCustomerParams {
	var ()
	return &UpdateContactCustomerParams{
		HTTPClient: client,
	}
}

/*UpdateContactCustomerParams contains all the parameters to send to the API endpoint
for the update contact customer operation typically these are written to a http.Request
*/
type UpdateContactCustomerParams struct {

	/*CustomerUpdateContactRequest*/
	CustomerUpdateContactRequest *models.CustomerUpdateContactRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update contact customer params
func (o *UpdateContactCustomerParams) WithTimeout(timeout time.Duration) *UpdateContactCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update contact customer params
func (o *UpdateContactCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update contact customer params
func (o *UpdateContactCustomerParams) WithContext(ctx context.Context) *UpdateContactCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update contact customer params
func (o *UpdateContactCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update contact customer params
func (o *UpdateContactCustomerParams) WithHTTPClient(client *http.Client) *UpdateContactCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update contact customer params
func (o *UpdateContactCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerUpdateContactRequest adds the customerUpdateContactRequest to the update contact customer params
func (o *UpdateContactCustomerParams) WithCustomerUpdateContactRequest(customerUpdateContactRequest *models.CustomerUpdateContactRequest) *UpdateContactCustomerParams {
	o.SetCustomerUpdateContactRequest(customerUpdateContactRequest)
	return o
}

// SetCustomerUpdateContactRequest adds the customerUpdateContactRequest to the update contact customer params
func (o *UpdateContactCustomerParams) SetCustomerUpdateContactRequest(customerUpdateContactRequest *models.CustomerUpdateContactRequest) {
	o.CustomerUpdateContactRequest = customerUpdateContactRequest
}

// WithID adds the id to the update contact customer params
func (o *UpdateContactCustomerParams) WithID(id string) *UpdateContactCustomerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update contact customer params
func (o *UpdateContactCustomerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateContactCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CustomerUpdateContactRequest == nil {
		o.CustomerUpdateContactRequest = new(models.CustomerUpdateContactRequest)
	}

	if err := r.SetBodyParam(o.CustomerUpdateContactRequest); err != nil {
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