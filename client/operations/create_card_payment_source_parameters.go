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

// NewCreateCardPaymentSourceParams creates a new CreateCardPaymentSourceParams object
// with the default values initialized.
func NewCreateCardPaymentSourceParams() *CreateCardPaymentSourceParams {
	var ()
	return &CreateCardPaymentSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCardPaymentSourceParamsWithTimeout creates a new CreateCardPaymentSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCardPaymentSourceParamsWithTimeout(timeout time.Duration) *CreateCardPaymentSourceParams {
	var ()
	return &CreateCardPaymentSourceParams{

		timeout: timeout,
	}
}

// NewCreateCardPaymentSourceParamsWithContext creates a new CreateCardPaymentSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCardPaymentSourceParamsWithContext(ctx context.Context) *CreateCardPaymentSourceParams {
	var ()
	return &CreateCardPaymentSourceParams{

		Context: ctx,
	}
}

// NewCreateCardPaymentSourceParamsWithHTTPClient creates a new CreateCardPaymentSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCardPaymentSourceParamsWithHTTPClient(client *http.Client) *CreateCardPaymentSourceParams {
	var ()
	return &CreateCardPaymentSourceParams{
		HTTPClient: client,
	}
}

/*CreateCardPaymentSourceParams contains all the parameters to send to the API endpoint
for the create card payment source operation typically these are written to a http.Request
*/
type CreateCardPaymentSourceParams struct {

	/*PaymentSourceCreateCardRequest*/
	PaymentSourceCreateCardRequest *models.PaymentSourceCreateCardRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create card payment source params
func (o *CreateCardPaymentSourceParams) WithTimeout(timeout time.Duration) *CreateCardPaymentSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create card payment source params
func (o *CreateCardPaymentSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create card payment source params
func (o *CreateCardPaymentSourceParams) WithContext(ctx context.Context) *CreateCardPaymentSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create card payment source params
func (o *CreateCardPaymentSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create card payment source params
func (o *CreateCardPaymentSourceParams) WithHTTPClient(client *http.Client) *CreateCardPaymentSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create card payment source params
func (o *CreateCardPaymentSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaymentSourceCreateCardRequest adds the paymentSourceCreateCardRequest to the create card payment source params
func (o *CreateCardPaymentSourceParams) WithPaymentSourceCreateCardRequest(paymentSourceCreateCardRequest *models.PaymentSourceCreateCardRequest) *CreateCardPaymentSourceParams {
	o.SetPaymentSourceCreateCardRequest(paymentSourceCreateCardRequest)
	return o
}

// SetPaymentSourceCreateCardRequest adds the paymentSourceCreateCardRequest to the create card payment source params
func (o *CreateCardPaymentSourceParams) SetPaymentSourceCreateCardRequest(paymentSourceCreateCardRequest *models.PaymentSourceCreateCardRequest) {
	o.PaymentSourceCreateCardRequest = paymentSourceCreateCardRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCardPaymentSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PaymentSourceCreateCardRequest != nil {
		if err := r.SetBodyParam(o.PaymentSourceCreateCardRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
