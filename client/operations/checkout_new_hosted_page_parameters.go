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

// NewCheckoutNewHostedPageParams creates a new CheckoutNewHostedPageParams object
// with the default values initialized.
func NewCheckoutNewHostedPageParams() *CheckoutNewHostedPageParams {
	var ()
	return &CheckoutNewHostedPageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckoutNewHostedPageParamsWithTimeout creates a new CheckoutNewHostedPageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckoutNewHostedPageParamsWithTimeout(timeout time.Duration) *CheckoutNewHostedPageParams {
	var ()
	return &CheckoutNewHostedPageParams{

		timeout: timeout,
	}
}

// NewCheckoutNewHostedPageParamsWithContext creates a new CheckoutNewHostedPageParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckoutNewHostedPageParamsWithContext(ctx context.Context) *CheckoutNewHostedPageParams {
	var ()
	return &CheckoutNewHostedPageParams{

		Context: ctx,
	}
}

// NewCheckoutNewHostedPageParamsWithHTTPClient creates a new CheckoutNewHostedPageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckoutNewHostedPageParamsWithHTTPClient(client *http.Client) *CheckoutNewHostedPageParams {
	var ()
	return &CheckoutNewHostedPageParams{
		HTTPClient: client,
	}
}

/*CheckoutNewHostedPageParams contains all the parameters to send to the API endpoint
for the checkout new hosted page operation typically these are written to a http.Request
*/
type CheckoutNewHostedPageParams struct {

	/*HostedPageCheckoutNewRequest*/
	HostedPageCheckoutNewRequest *models.HostedPageCheckoutNewRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the checkout new hosted page params
func (o *CheckoutNewHostedPageParams) WithTimeout(timeout time.Duration) *CheckoutNewHostedPageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checkout new hosted page params
func (o *CheckoutNewHostedPageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checkout new hosted page params
func (o *CheckoutNewHostedPageParams) WithContext(ctx context.Context) *CheckoutNewHostedPageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checkout new hosted page params
func (o *CheckoutNewHostedPageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checkout new hosted page params
func (o *CheckoutNewHostedPageParams) WithHTTPClient(client *http.Client) *CheckoutNewHostedPageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checkout new hosted page params
func (o *CheckoutNewHostedPageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHostedPageCheckoutNewRequest adds the hostedPageCheckoutNewRequest to the checkout new hosted page params
func (o *CheckoutNewHostedPageParams) WithHostedPageCheckoutNewRequest(hostedPageCheckoutNewRequest *models.HostedPageCheckoutNewRequest) *CheckoutNewHostedPageParams {
	o.SetHostedPageCheckoutNewRequest(hostedPageCheckoutNewRequest)
	return o
}

// SetHostedPageCheckoutNewRequest adds the hostedPageCheckoutNewRequest to the checkout new hosted page params
func (o *CheckoutNewHostedPageParams) SetHostedPageCheckoutNewRequest(hostedPageCheckoutNewRequest *models.HostedPageCheckoutNewRequest) {
	o.HostedPageCheckoutNewRequest = hostedPageCheckoutNewRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CheckoutNewHostedPageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HostedPageCheckoutNewRequest == nil {
		o.HostedPageCheckoutNewRequest = new(models.HostedPageCheckoutNewRequest)
	}

	if err := r.SetBodyParam(o.HostedPageCheckoutNewRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}