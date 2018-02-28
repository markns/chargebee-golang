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

// NewDeleteInvoiceParams creates a new DeleteInvoiceParams object
// with the default values initialized.
func NewDeleteInvoiceParams() *DeleteInvoiceParams {
	var ()
	return &DeleteInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInvoiceParamsWithTimeout creates a new DeleteInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInvoiceParamsWithTimeout(timeout time.Duration) *DeleteInvoiceParams {
	var ()
	return &DeleteInvoiceParams{

		timeout: timeout,
	}
}

// NewDeleteInvoiceParamsWithContext creates a new DeleteInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInvoiceParamsWithContext(ctx context.Context) *DeleteInvoiceParams {
	var ()
	return &DeleteInvoiceParams{

		Context: ctx,
	}
}

// NewDeleteInvoiceParamsWithHTTPClient creates a new DeleteInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInvoiceParamsWithHTTPClient(client *http.Client) *DeleteInvoiceParams {
	var ()
	return &DeleteInvoiceParams{
		HTTPClient: client,
	}
}

/*DeleteInvoiceParams contains all the parameters to send to the API endpoint
for the delete invoice operation typically these are written to a http.Request
*/
type DeleteInvoiceParams struct {

	/*InvoiceDeleteRequest*/
	InvoiceDeleteRequest *models.InvoiceDeleteRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete invoice params
func (o *DeleteInvoiceParams) WithTimeout(timeout time.Duration) *DeleteInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete invoice params
func (o *DeleteInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete invoice params
func (o *DeleteInvoiceParams) WithContext(ctx context.Context) *DeleteInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete invoice params
func (o *DeleteInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete invoice params
func (o *DeleteInvoiceParams) WithHTTPClient(client *http.Client) *DeleteInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete invoice params
func (o *DeleteInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvoiceDeleteRequest adds the invoiceDeleteRequest to the delete invoice params
func (o *DeleteInvoiceParams) WithInvoiceDeleteRequest(invoiceDeleteRequest *models.InvoiceDeleteRequest) *DeleteInvoiceParams {
	o.SetInvoiceDeleteRequest(invoiceDeleteRequest)
	return o
}

// SetInvoiceDeleteRequest adds the invoiceDeleteRequest to the delete invoice params
func (o *DeleteInvoiceParams) SetInvoiceDeleteRequest(invoiceDeleteRequest *models.InvoiceDeleteRequest) {
	o.InvoiceDeleteRequest = invoiceDeleteRequest
}

// WithID adds the id to the delete invoice params
func (o *DeleteInvoiceParams) WithID(id string) *DeleteInvoiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete invoice params
func (o *DeleteInvoiceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InvoiceDeleteRequest != nil {
		if err := r.SetBodyParam(o.InvoiceDeleteRequest); err != nil {
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
