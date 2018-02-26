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

// NewVoidInvoiceInvoiceParams creates a new VoidInvoiceInvoiceParams object
// with the default values initialized.
func NewVoidInvoiceInvoiceParams() *VoidInvoiceInvoiceParams {
	var ()
	return &VoidInvoiceInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVoidInvoiceInvoiceParamsWithTimeout creates a new VoidInvoiceInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVoidInvoiceInvoiceParamsWithTimeout(timeout time.Duration) *VoidInvoiceInvoiceParams {
	var ()
	return &VoidInvoiceInvoiceParams{

		timeout: timeout,
	}
}

// NewVoidInvoiceInvoiceParamsWithContext creates a new VoidInvoiceInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewVoidInvoiceInvoiceParamsWithContext(ctx context.Context) *VoidInvoiceInvoiceParams {
	var ()
	return &VoidInvoiceInvoiceParams{

		Context: ctx,
	}
}

// NewVoidInvoiceInvoiceParamsWithHTTPClient creates a new VoidInvoiceInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVoidInvoiceInvoiceParamsWithHTTPClient(client *http.Client) *VoidInvoiceInvoiceParams {
	var ()
	return &VoidInvoiceInvoiceParams{
		HTTPClient: client,
	}
}

/*VoidInvoiceInvoiceParams contains all the parameters to send to the API endpoint
for the void invoice invoice operation typically these are written to a http.Request
*/
type VoidInvoiceInvoiceParams struct {

	/*InvoiceVoidInvoiceRequest*/
	InvoiceVoidInvoiceRequest *models.InvoiceVoidInvoiceRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the void invoice invoice params
func (o *VoidInvoiceInvoiceParams) WithTimeout(timeout time.Duration) *VoidInvoiceInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the void invoice invoice params
func (o *VoidInvoiceInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the void invoice invoice params
func (o *VoidInvoiceInvoiceParams) WithContext(ctx context.Context) *VoidInvoiceInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the void invoice invoice params
func (o *VoidInvoiceInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the void invoice invoice params
func (o *VoidInvoiceInvoiceParams) WithHTTPClient(client *http.Client) *VoidInvoiceInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the void invoice invoice params
func (o *VoidInvoiceInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvoiceVoidInvoiceRequest adds the invoiceVoidInvoiceRequest to the void invoice invoice params
func (o *VoidInvoiceInvoiceParams) WithInvoiceVoidInvoiceRequest(invoiceVoidInvoiceRequest *models.InvoiceVoidInvoiceRequest) *VoidInvoiceInvoiceParams {
	o.SetInvoiceVoidInvoiceRequest(invoiceVoidInvoiceRequest)
	return o
}

// SetInvoiceVoidInvoiceRequest adds the invoiceVoidInvoiceRequest to the void invoice invoice params
func (o *VoidInvoiceInvoiceParams) SetInvoiceVoidInvoiceRequest(invoiceVoidInvoiceRequest *models.InvoiceVoidInvoiceRequest) {
	o.InvoiceVoidInvoiceRequest = invoiceVoidInvoiceRequest
}

// WithID adds the id to the void invoice invoice params
func (o *VoidInvoiceInvoiceParams) WithID(id string) *VoidInvoiceInvoiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the void invoice invoice params
func (o *VoidInvoiceInvoiceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VoidInvoiceInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InvoiceVoidInvoiceRequest == nil {
		o.InvoiceVoidInvoiceRequest = new(models.InvoiceVoidInvoiceRequest)
	}

	if err := r.SetBodyParam(o.InvoiceVoidInvoiceRequest); err != nil {
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