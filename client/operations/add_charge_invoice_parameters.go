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

// NewAddChargeInvoiceParams creates a new AddChargeInvoiceParams object
// with the default values initialized.
func NewAddChargeInvoiceParams() *AddChargeInvoiceParams {
	var ()
	return &AddChargeInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddChargeInvoiceParamsWithTimeout creates a new AddChargeInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddChargeInvoiceParamsWithTimeout(timeout time.Duration) *AddChargeInvoiceParams {
	var ()
	return &AddChargeInvoiceParams{

		timeout: timeout,
	}
}

// NewAddChargeInvoiceParamsWithContext creates a new AddChargeInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddChargeInvoiceParamsWithContext(ctx context.Context) *AddChargeInvoiceParams {
	var ()
	return &AddChargeInvoiceParams{

		Context: ctx,
	}
}

// NewAddChargeInvoiceParamsWithHTTPClient creates a new AddChargeInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddChargeInvoiceParamsWithHTTPClient(client *http.Client) *AddChargeInvoiceParams {
	var ()
	return &AddChargeInvoiceParams{
		HTTPClient: client,
	}
}

/*AddChargeInvoiceParams contains all the parameters to send to the API endpoint
for the add charge invoice operation typically these are written to a http.Request
*/
type AddChargeInvoiceParams struct {

	/*InvoiceAddChargeRequest*/
	InvoiceAddChargeRequest *models.InvoiceAddChargeRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add charge invoice params
func (o *AddChargeInvoiceParams) WithTimeout(timeout time.Duration) *AddChargeInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add charge invoice params
func (o *AddChargeInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add charge invoice params
func (o *AddChargeInvoiceParams) WithContext(ctx context.Context) *AddChargeInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add charge invoice params
func (o *AddChargeInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add charge invoice params
func (o *AddChargeInvoiceParams) WithHTTPClient(client *http.Client) *AddChargeInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add charge invoice params
func (o *AddChargeInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvoiceAddChargeRequest adds the invoiceAddChargeRequest to the add charge invoice params
func (o *AddChargeInvoiceParams) WithInvoiceAddChargeRequest(invoiceAddChargeRequest *models.InvoiceAddChargeRequest) *AddChargeInvoiceParams {
	o.SetInvoiceAddChargeRequest(invoiceAddChargeRequest)
	return o
}

// SetInvoiceAddChargeRequest adds the invoiceAddChargeRequest to the add charge invoice params
func (o *AddChargeInvoiceParams) SetInvoiceAddChargeRequest(invoiceAddChargeRequest *models.InvoiceAddChargeRequest) {
	o.InvoiceAddChargeRequest = invoiceAddChargeRequest
}

// WithID adds the id to the add charge invoice params
func (o *AddChargeInvoiceParams) WithID(id string) *AddChargeInvoiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add charge invoice params
func (o *AddChargeInvoiceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddChargeInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InvoiceAddChargeRequest != nil {
		if err := r.SetBodyParam(o.InvoiceAddChargeRequest); err != nil {
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
