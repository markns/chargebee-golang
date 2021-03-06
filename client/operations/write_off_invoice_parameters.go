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

// NewWriteOffInvoiceParams creates a new WriteOffInvoiceParams object
// with the default values initialized.
func NewWriteOffInvoiceParams() *WriteOffInvoiceParams {
	var ()
	return &WriteOffInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWriteOffInvoiceParamsWithTimeout creates a new WriteOffInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWriteOffInvoiceParamsWithTimeout(timeout time.Duration) *WriteOffInvoiceParams {
	var ()
	return &WriteOffInvoiceParams{

		timeout: timeout,
	}
}

// NewWriteOffInvoiceParamsWithContext creates a new WriteOffInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewWriteOffInvoiceParamsWithContext(ctx context.Context) *WriteOffInvoiceParams {
	var ()
	return &WriteOffInvoiceParams{

		Context: ctx,
	}
}

// NewWriteOffInvoiceParamsWithHTTPClient creates a new WriteOffInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWriteOffInvoiceParamsWithHTTPClient(client *http.Client) *WriteOffInvoiceParams {
	var ()
	return &WriteOffInvoiceParams{
		HTTPClient: client,
	}
}

/*WriteOffInvoiceParams contains all the parameters to send to the API endpoint
for the write off invoice operation typically these are written to a http.Request
*/
type WriteOffInvoiceParams struct {

	/*InvoiceWriteOffRequest*/
	InvoiceWriteOffRequest *models.InvoiceWriteOffRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the write off invoice params
func (o *WriteOffInvoiceParams) WithTimeout(timeout time.Duration) *WriteOffInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the write off invoice params
func (o *WriteOffInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the write off invoice params
func (o *WriteOffInvoiceParams) WithContext(ctx context.Context) *WriteOffInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the write off invoice params
func (o *WriteOffInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the write off invoice params
func (o *WriteOffInvoiceParams) WithHTTPClient(client *http.Client) *WriteOffInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the write off invoice params
func (o *WriteOffInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvoiceWriteOffRequest adds the invoiceWriteOffRequest to the write off invoice params
func (o *WriteOffInvoiceParams) WithInvoiceWriteOffRequest(invoiceWriteOffRequest *models.InvoiceWriteOffRequest) *WriteOffInvoiceParams {
	o.SetInvoiceWriteOffRequest(invoiceWriteOffRequest)
	return o
}

// SetInvoiceWriteOffRequest adds the invoiceWriteOffRequest to the write off invoice params
func (o *WriteOffInvoiceParams) SetInvoiceWriteOffRequest(invoiceWriteOffRequest *models.InvoiceWriteOffRequest) {
	o.InvoiceWriteOffRequest = invoiceWriteOffRequest
}

// WithID adds the id to the write off invoice params
func (o *WriteOffInvoiceParams) WithID(id string) *WriteOffInvoiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the write off invoice params
func (o *WriteOffInvoiceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *WriteOffInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InvoiceWriteOffRequest != nil {
		if err := r.SetBodyParam(o.InvoiceWriteOffRequest); err != nil {
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
