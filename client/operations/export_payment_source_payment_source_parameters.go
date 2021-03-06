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

// NewExportPaymentSourcePaymentSourceParams creates a new ExportPaymentSourcePaymentSourceParams object
// with the default values initialized.
func NewExportPaymentSourcePaymentSourceParams() *ExportPaymentSourcePaymentSourceParams {
	var ()
	return &ExportPaymentSourcePaymentSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExportPaymentSourcePaymentSourceParamsWithTimeout creates a new ExportPaymentSourcePaymentSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExportPaymentSourcePaymentSourceParamsWithTimeout(timeout time.Duration) *ExportPaymentSourcePaymentSourceParams {
	var ()
	return &ExportPaymentSourcePaymentSourceParams{

		timeout: timeout,
	}
}

// NewExportPaymentSourcePaymentSourceParamsWithContext creates a new ExportPaymentSourcePaymentSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewExportPaymentSourcePaymentSourceParamsWithContext(ctx context.Context) *ExportPaymentSourcePaymentSourceParams {
	var ()
	return &ExportPaymentSourcePaymentSourceParams{

		Context: ctx,
	}
}

// NewExportPaymentSourcePaymentSourceParamsWithHTTPClient creates a new ExportPaymentSourcePaymentSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExportPaymentSourcePaymentSourceParamsWithHTTPClient(client *http.Client) *ExportPaymentSourcePaymentSourceParams {
	var ()
	return &ExportPaymentSourcePaymentSourceParams{
		HTTPClient: client,
	}
}

/*ExportPaymentSourcePaymentSourceParams contains all the parameters to send to the API endpoint
for the export payment source payment source operation typically these are written to a http.Request
*/
type ExportPaymentSourcePaymentSourceParams struct {

	/*PaymentSourceExportPaymentSourceRequest*/
	PaymentSourceExportPaymentSourceRequest *models.PaymentSourceExportPaymentSourceRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the export payment source payment source params
func (o *ExportPaymentSourcePaymentSourceParams) WithTimeout(timeout time.Duration) *ExportPaymentSourcePaymentSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export payment source payment source params
func (o *ExportPaymentSourcePaymentSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export payment source payment source params
func (o *ExportPaymentSourcePaymentSourceParams) WithContext(ctx context.Context) *ExportPaymentSourcePaymentSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export payment source payment source params
func (o *ExportPaymentSourcePaymentSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export payment source payment source params
func (o *ExportPaymentSourcePaymentSourceParams) WithHTTPClient(client *http.Client) *ExportPaymentSourcePaymentSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export payment source payment source params
func (o *ExportPaymentSourcePaymentSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaymentSourceExportPaymentSourceRequest adds the paymentSourceExportPaymentSourceRequest to the export payment source payment source params
func (o *ExportPaymentSourcePaymentSourceParams) WithPaymentSourceExportPaymentSourceRequest(paymentSourceExportPaymentSourceRequest *models.PaymentSourceExportPaymentSourceRequest) *ExportPaymentSourcePaymentSourceParams {
	o.SetPaymentSourceExportPaymentSourceRequest(paymentSourceExportPaymentSourceRequest)
	return o
}

// SetPaymentSourceExportPaymentSourceRequest adds the paymentSourceExportPaymentSourceRequest to the export payment source payment source params
func (o *ExportPaymentSourcePaymentSourceParams) SetPaymentSourceExportPaymentSourceRequest(paymentSourceExportPaymentSourceRequest *models.PaymentSourceExportPaymentSourceRequest) {
	o.PaymentSourceExportPaymentSourceRequest = paymentSourceExportPaymentSourceRequest
}

// WithID adds the id to the export payment source payment source params
func (o *ExportPaymentSourcePaymentSourceParams) WithID(id string) *ExportPaymentSourcePaymentSourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the export payment source payment source params
func (o *ExportPaymentSourcePaymentSourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExportPaymentSourcePaymentSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PaymentSourceExportPaymentSourceRequest != nil {
		if err := r.SetBodyParam(o.PaymentSourceExportPaymentSourceRequest); err != nil {
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
