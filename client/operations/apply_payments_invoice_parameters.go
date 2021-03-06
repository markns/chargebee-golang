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

// NewApplyPaymentsInvoiceParams creates a new ApplyPaymentsInvoiceParams object
// with the default values initialized.
func NewApplyPaymentsInvoiceParams() *ApplyPaymentsInvoiceParams {
	var ()
	return &ApplyPaymentsInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewApplyPaymentsInvoiceParamsWithTimeout creates a new ApplyPaymentsInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewApplyPaymentsInvoiceParamsWithTimeout(timeout time.Duration) *ApplyPaymentsInvoiceParams {
	var ()
	return &ApplyPaymentsInvoiceParams{

		timeout: timeout,
	}
}

// NewApplyPaymentsInvoiceParamsWithContext creates a new ApplyPaymentsInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewApplyPaymentsInvoiceParamsWithContext(ctx context.Context) *ApplyPaymentsInvoiceParams {
	var ()
	return &ApplyPaymentsInvoiceParams{

		Context: ctx,
	}
}

// NewApplyPaymentsInvoiceParamsWithHTTPClient creates a new ApplyPaymentsInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewApplyPaymentsInvoiceParamsWithHTTPClient(client *http.Client) *ApplyPaymentsInvoiceParams {
	var ()
	return &ApplyPaymentsInvoiceParams{
		HTTPClient: client,
	}
}

/*ApplyPaymentsInvoiceParams contains all the parameters to send to the API endpoint
for the apply payments invoice operation typically these are written to a http.Request
*/
type ApplyPaymentsInvoiceParams struct {

	/*InvoiceApplyPaymentsRequest*/
	InvoiceApplyPaymentsRequest models.InvoiceApplyPaymentsRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the apply payments invoice params
func (o *ApplyPaymentsInvoiceParams) WithTimeout(timeout time.Duration) *ApplyPaymentsInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the apply payments invoice params
func (o *ApplyPaymentsInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the apply payments invoice params
func (o *ApplyPaymentsInvoiceParams) WithContext(ctx context.Context) *ApplyPaymentsInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the apply payments invoice params
func (o *ApplyPaymentsInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the apply payments invoice params
func (o *ApplyPaymentsInvoiceParams) WithHTTPClient(client *http.Client) *ApplyPaymentsInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the apply payments invoice params
func (o *ApplyPaymentsInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvoiceApplyPaymentsRequest adds the invoiceApplyPaymentsRequest to the apply payments invoice params
func (o *ApplyPaymentsInvoiceParams) WithInvoiceApplyPaymentsRequest(invoiceApplyPaymentsRequest models.InvoiceApplyPaymentsRequest) *ApplyPaymentsInvoiceParams {
	o.SetInvoiceApplyPaymentsRequest(invoiceApplyPaymentsRequest)
	return o
}

// SetInvoiceApplyPaymentsRequest adds the invoiceApplyPaymentsRequest to the apply payments invoice params
func (o *ApplyPaymentsInvoiceParams) SetInvoiceApplyPaymentsRequest(invoiceApplyPaymentsRequest models.InvoiceApplyPaymentsRequest) {
	o.InvoiceApplyPaymentsRequest = invoiceApplyPaymentsRequest
}

// WithID adds the id to the apply payments invoice params
func (o *ApplyPaymentsInvoiceParams) WithID(id string) *ApplyPaymentsInvoiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the apply payments invoice params
func (o *ApplyPaymentsInvoiceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ApplyPaymentsInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InvoiceApplyPaymentsRequest != nil {
		if err := r.SetBodyParam(o.InvoiceApplyPaymentsRequest); err != nil {
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
