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

	models "github.com/markns/chargebee-golang/models"
)

// NewAddContactCustomerParams creates a new AddContactCustomerParams object
// with the default values initialized.
func NewAddContactCustomerParams() *AddContactCustomerParams {
	var ()
	return &AddContactCustomerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddContactCustomerParamsWithTimeout creates a new AddContactCustomerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddContactCustomerParamsWithTimeout(timeout time.Duration) *AddContactCustomerParams {
	var ()
	return &AddContactCustomerParams{

		timeout: timeout,
	}
}

// NewAddContactCustomerParamsWithContext creates a new AddContactCustomerParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddContactCustomerParamsWithContext(ctx context.Context) *AddContactCustomerParams {
	var ()
	return &AddContactCustomerParams{

		Context: ctx,
	}
}

// NewAddContactCustomerParamsWithHTTPClient creates a new AddContactCustomerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddContactCustomerParamsWithHTTPClient(client *http.Client) *AddContactCustomerParams {
	var ()
	return &AddContactCustomerParams{
		HTTPClient: client,
	}
}

/*AddContactCustomerParams contains all the parameters to send to the API endpoint
for the add contact customer operation typically these are written to a http.Request
*/
type AddContactCustomerParams struct {

	/*CustomerAddContactRequest*/
	CustomerAddContactRequest *models.CustomerAddContactRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add contact customer params
func (o *AddContactCustomerParams) WithTimeout(timeout time.Duration) *AddContactCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add contact customer params
func (o *AddContactCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add contact customer params
func (o *AddContactCustomerParams) WithContext(ctx context.Context) *AddContactCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add contact customer params
func (o *AddContactCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add contact customer params
func (o *AddContactCustomerParams) WithHTTPClient(client *http.Client) *AddContactCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add contact customer params
func (o *AddContactCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerAddContactRequest adds the customerAddContactRequest to the add contact customer params
func (o *AddContactCustomerParams) WithCustomerAddContactRequest(customerAddContactRequest *models.CustomerAddContactRequest) *AddContactCustomerParams {
	o.SetCustomerAddContactRequest(customerAddContactRequest)
	return o
}

// SetCustomerAddContactRequest adds the customerAddContactRequest to the add contact customer params
func (o *AddContactCustomerParams) SetCustomerAddContactRequest(customerAddContactRequest *models.CustomerAddContactRequest) {
	o.CustomerAddContactRequest = customerAddContactRequest
}

// WithID adds the id to the add contact customer params
func (o *AddContactCustomerParams) WithID(id string) *AddContactCustomerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add contact customer params
func (o *AddContactCustomerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddContactCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CustomerAddContactRequest != nil {
		if err := r.SetBodyParam(o.CustomerAddContactRequest); err != nil {
			return err
		}
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
