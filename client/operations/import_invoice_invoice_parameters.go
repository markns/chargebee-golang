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

// NewImportInvoiceInvoiceParams creates a new ImportInvoiceInvoiceParams object
// with the default values initialized.
func NewImportInvoiceInvoiceParams() *ImportInvoiceInvoiceParams {
	var ()
	return &ImportInvoiceInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImportInvoiceInvoiceParamsWithTimeout creates a new ImportInvoiceInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImportInvoiceInvoiceParamsWithTimeout(timeout time.Duration) *ImportInvoiceInvoiceParams {
	var ()
	return &ImportInvoiceInvoiceParams{

		timeout: timeout,
	}
}

// NewImportInvoiceInvoiceParamsWithContext creates a new ImportInvoiceInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewImportInvoiceInvoiceParamsWithContext(ctx context.Context) *ImportInvoiceInvoiceParams {
	var ()
	return &ImportInvoiceInvoiceParams{

		Context: ctx,
	}
}

// NewImportInvoiceInvoiceParamsWithHTTPClient creates a new ImportInvoiceInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImportInvoiceInvoiceParamsWithHTTPClient(client *http.Client) *ImportInvoiceInvoiceParams {
	var ()
	return &ImportInvoiceInvoiceParams{
		HTTPClient: client,
	}
}

/*ImportInvoiceInvoiceParams contains all the parameters to send to the API endpoint
for the import invoice invoice operation typically these are written to a http.Request
*/
type ImportInvoiceInvoiceParams struct {

	/*InvoiceImportInvoiceRequest*/
	InvoiceImportInvoiceRequest *models.InvoiceImportInvoiceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the import invoice invoice params
func (o *ImportInvoiceInvoiceParams) WithTimeout(timeout time.Duration) *ImportInvoiceInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import invoice invoice params
func (o *ImportInvoiceInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import invoice invoice params
func (o *ImportInvoiceInvoiceParams) WithContext(ctx context.Context) *ImportInvoiceInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import invoice invoice params
func (o *ImportInvoiceInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import invoice invoice params
func (o *ImportInvoiceInvoiceParams) WithHTTPClient(client *http.Client) *ImportInvoiceInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import invoice invoice params
func (o *ImportInvoiceInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvoiceImportInvoiceRequest adds the invoiceImportInvoiceRequest to the import invoice invoice params
func (o *ImportInvoiceInvoiceParams) WithInvoiceImportInvoiceRequest(invoiceImportInvoiceRequest *models.InvoiceImportInvoiceRequest) *ImportInvoiceInvoiceParams {
	o.SetInvoiceImportInvoiceRequest(invoiceImportInvoiceRequest)
	return o
}

// SetInvoiceImportInvoiceRequest adds the invoiceImportInvoiceRequest to the import invoice invoice params
func (o *ImportInvoiceInvoiceParams) SetInvoiceImportInvoiceRequest(invoiceImportInvoiceRequest *models.InvoiceImportInvoiceRequest) {
	o.InvoiceImportInvoiceRequest = invoiceImportInvoiceRequest
}

// WriteToRequest writes these params to a swagger request
func (o *ImportInvoiceInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InvoiceImportInvoiceRequest != nil {
		if err := r.SetBodyParam(o.InvoiceImportInvoiceRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
