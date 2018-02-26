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

// NewCreateUsingPermanentTokenPaymentSourceParams creates a new CreateUsingPermanentTokenPaymentSourceParams object
// with the default values initialized.
func NewCreateUsingPermanentTokenPaymentSourceParams() *CreateUsingPermanentTokenPaymentSourceParams {
	var ()
	return &CreateUsingPermanentTokenPaymentSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUsingPermanentTokenPaymentSourceParamsWithTimeout creates a new CreateUsingPermanentTokenPaymentSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUsingPermanentTokenPaymentSourceParamsWithTimeout(timeout time.Duration) *CreateUsingPermanentTokenPaymentSourceParams {
	var ()
	return &CreateUsingPermanentTokenPaymentSourceParams{

		timeout: timeout,
	}
}

// NewCreateUsingPermanentTokenPaymentSourceParamsWithContext creates a new CreateUsingPermanentTokenPaymentSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateUsingPermanentTokenPaymentSourceParamsWithContext(ctx context.Context) *CreateUsingPermanentTokenPaymentSourceParams {
	var ()
	return &CreateUsingPermanentTokenPaymentSourceParams{

		Context: ctx,
	}
}

// NewCreateUsingPermanentTokenPaymentSourceParamsWithHTTPClient creates a new CreateUsingPermanentTokenPaymentSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateUsingPermanentTokenPaymentSourceParamsWithHTTPClient(client *http.Client) *CreateUsingPermanentTokenPaymentSourceParams {
	var ()
	return &CreateUsingPermanentTokenPaymentSourceParams{
		HTTPClient: client,
	}
}

/*CreateUsingPermanentTokenPaymentSourceParams contains all the parameters to send to the API endpoint
for the create using permanent token payment source operation typically these are written to a http.Request
*/
type CreateUsingPermanentTokenPaymentSourceParams struct {

	/*PaymentSourceCreateUsingPermanentTokenRequest*/
	PaymentSourceCreateUsingPermanentTokenRequest *models.PaymentSourceCreateUsingPermanentTokenRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create using permanent token payment source params
func (o *CreateUsingPermanentTokenPaymentSourceParams) WithTimeout(timeout time.Duration) *CreateUsingPermanentTokenPaymentSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create using permanent token payment source params
func (o *CreateUsingPermanentTokenPaymentSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create using permanent token payment source params
func (o *CreateUsingPermanentTokenPaymentSourceParams) WithContext(ctx context.Context) *CreateUsingPermanentTokenPaymentSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create using permanent token payment source params
func (o *CreateUsingPermanentTokenPaymentSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create using permanent token payment source params
func (o *CreateUsingPermanentTokenPaymentSourceParams) WithHTTPClient(client *http.Client) *CreateUsingPermanentTokenPaymentSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create using permanent token payment source params
func (o *CreateUsingPermanentTokenPaymentSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaymentSourceCreateUsingPermanentTokenRequest adds the paymentSourceCreateUsingPermanentTokenRequest to the create using permanent token payment source params
func (o *CreateUsingPermanentTokenPaymentSourceParams) WithPaymentSourceCreateUsingPermanentTokenRequest(paymentSourceCreateUsingPermanentTokenRequest *models.PaymentSourceCreateUsingPermanentTokenRequest) *CreateUsingPermanentTokenPaymentSourceParams {
	o.SetPaymentSourceCreateUsingPermanentTokenRequest(paymentSourceCreateUsingPermanentTokenRequest)
	return o
}

// SetPaymentSourceCreateUsingPermanentTokenRequest adds the paymentSourceCreateUsingPermanentTokenRequest to the create using permanent token payment source params
func (o *CreateUsingPermanentTokenPaymentSourceParams) SetPaymentSourceCreateUsingPermanentTokenRequest(paymentSourceCreateUsingPermanentTokenRequest *models.PaymentSourceCreateUsingPermanentTokenRequest) {
	o.PaymentSourceCreateUsingPermanentTokenRequest = paymentSourceCreateUsingPermanentTokenRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUsingPermanentTokenPaymentSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PaymentSourceCreateUsingPermanentTokenRequest == nil {
		o.PaymentSourceCreateUsingPermanentTokenRequest = new(models.PaymentSourceCreateUsingPermanentTokenRequest)
	}

	if err := r.SetBodyParam(o.PaymentSourceCreateUsingPermanentTokenRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
