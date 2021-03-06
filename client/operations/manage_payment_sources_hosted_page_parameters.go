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

// NewManagePaymentSourcesHostedPageParams creates a new ManagePaymentSourcesHostedPageParams object
// with the default values initialized.
func NewManagePaymentSourcesHostedPageParams() *ManagePaymentSourcesHostedPageParams {
	var ()
	return &ManagePaymentSourcesHostedPageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewManagePaymentSourcesHostedPageParamsWithTimeout creates a new ManagePaymentSourcesHostedPageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewManagePaymentSourcesHostedPageParamsWithTimeout(timeout time.Duration) *ManagePaymentSourcesHostedPageParams {
	var ()
	return &ManagePaymentSourcesHostedPageParams{

		timeout: timeout,
	}
}

// NewManagePaymentSourcesHostedPageParamsWithContext creates a new ManagePaymentSourcesHostedPageParams object
// with the default values initialized, and the ability to set a context for a request
func NewManagePaymentSourcesHostedPageParamsWithContext(ctx context.Context) *ManagePaymentSourcesHostedPageParams {
	var ()
	return &ManagePaymentSourcesHostedPageParams{

		Context: ctx,
	}
}

// NewManagePaymentSourcesHostedPageParamsWithHTTPClient creates a new ManagePaymentSourcesHostedPageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewManagePaymentSourcesHostedPageParamsWithHTTPClient(client *http.Client) *ManagePaymentSourcesHostedPageParams {
	var ()
	return &ManagePaymentSourcesHostedPageParams{
		HTTPClient: client,
	}
}

/*ManagePaymentSourcesHostedPageParams contains all the parameters to send to the API endpoint
for the manage payment sources hosted page operation typically these are written to a http.Request
*/
type ManagePaymentSourcesHostedPageParams struct {

	/*HostedPageManagePaymentSourcesRequest*/
	HostedPageManagePaymentSourcesRequest *models.HostedPageManagePaymentSourcesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the manage payment sources hosted page params
func (o *ManagePaymentSourcesHostedPageParams) WithTimeout(timeout time.Duration) *ManagePaymentSourcesHostedPageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the manage payment sources hosted page params
func (o *ManagePaymentSourcesHostedPageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the manage payment sources hosted page params
func (o *ManagePaymentSourcesHostedPageParams) WithContext(ctx context.Context) *ManagePaymentSourcesHostedPageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the manage payment sources hosted page params
func (o *ManagePaymentSourcesHostedPageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the manage payment sources hosted page params
func (o *ManagePaymentSourcesHostedPageParams) WithHTTPClient(client *http.Client) *ManagePaymentSourcesHostedPageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the manage payment sources hosted page params
func (o *ManagePaymentSourcesHostedPageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHostedPageManagePaymentSourcesRequest adds the hostedPageManagePaymentSourcesRequest to the manage payment sources hosted page params
func (o *ManagePaymentSourcesHostedPageParams) WithHostedPageManagePaymentSourcesRequest(hostedPageManagePaymentSourcesRequest *models.HostedPageManagePaymentSourcesRequest) *ManagePaymentSourcesHostedPageParams {
	o.SetHostedPageManagePaymentSourcesRequest(hostedPageManagePaymentSourcesRequest)
	return o
}

// SetHostedPageManagePaymentSourcesRequest adds the hostedPageManagePaymentSourcesRequest to the manage payment sources hosted page params
func (o *ManagePaymentSourcesHostedPageParams) SetHostedPageManagePaymentSourcesRequest(hostedPageManagePaymentSourcesRequest *models.HostedPageManagePaymentSourcesRequest) {
	o.HostedPageManagePaymentSourcesRequest = hostedPageManagePaymentSourcesRequest
}

// WriteToRequest writes these params to a swagger request
func (o *ManagePaymentSourcesHostedPageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HostedPageManagePaymentSourcesRequest != nil {
		if err := r.SetBodyParam(o.HostedPageManagePaymentSourcesRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
