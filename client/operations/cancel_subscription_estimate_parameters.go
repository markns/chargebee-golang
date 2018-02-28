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

// NewCancelSubscriptionEstimateParams creates a new CancelSubscriptionEstimateParams object
// with the default values initialized.
func NewCancelSubscriptionEstimateParams() *CancelSubscriptionEstimateParams {
	var ()
	return &CancelSubscriptionEstimateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelSubscriptionEstimateParamsWithTimeout creates a new CancelSubscriptionEstimateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelSubscriptionEstimateParamsWithTimeout(timeout time.Duration) *CancelSubscriptionEstimateParams {
	var ()
	return &CancelSubscriptionEstimateParams{

		timeout: timeout,
	}
}

// NewCancelSubscriptionEstimateParamsWithContext creates a new CancelSubscriptionEstimateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelSubscriptionEstimateParamsWithContext(ctx context.Context) *CancelSubscriptionEstimateParams {
	var ()
	return &CancelSubscriptionEstimateParams{

		Context: ctx,
	}
}

// NewCancelSubscriptionEstimateParamsWithHTTPClient creates a new CancelSubscriptionEstimateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelSubscriptionEstimateParamsWithHTTPClient(client *http.Client) *CancelSubscriptionEstimateParams {
	var ()
	return &CancelSubscriptionEstimateParams{
		HTTPClient: client,
	}
}

/*CancelSubscriptionEstimateParams contains all the parameters to send to the API endpoint
for the cancel subscription estimate operation typically these are written to a http.Request
*/
type CancelSubscriptionEstimateParams struct {

	/*EstimateCancelSubscriptionRequest*/
	EstimateCancelSubscriptionRequest *models.EstimateCancelSubscriptionRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel subscription estimate params
func (o *CancelSubscriptionEstimateParams) WithTimeout(timeout time.Duration) *CancelSubscriptionEstimateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel subscription estimate params
func (o *CancelSubscriptionEstimateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel subscription estimate params
func (o *CancelSubscriptionEstimateParams) WithContext(ctx context.Context) *CancelSubscriptionEstimateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel subscription estimate params
func (o *CancelSubscriptionEstimateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel subscription estimate params
func (o *CancelSubscriptionEstimateParams) WithHTTPClient(client *http.Client) *CancelSubscriptionEstimateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel subscription estimate params
func (o *CancelSubscriptionEstimateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEstimateCancelSubscriptionRequest adds the estimateCancelSubscriptionRequest to the cancel subscription estimate params
func (o *CancelSubscriptionEstimateParams) WithEstimateCancelSubscriptionRequest(estimateCancelSubscriptionRequest *models.EstimateCancelSubscriptionRequest) *CancelSubscriptionEstimateParams {
	o.SetEstimateCancelSubscriptionRequest(estimateCancelSubscriptionRequest)
	return o
}

// SetEstimateCancelSubscriptionRequest adds the estimateCancelSubscriptionRequest to the cancel subscription estimate params
func (o *CancelSubscriptionEstimateParams) SetEstimateCancelSubscriptionRequest(estimateCancelSubscriptionRequest *models.EstimateCancelSubscriptionRequest) {
	o.EstimateCancelSubscriptionRequest = estimateCancelSubscriptionRequest
}

// WithID adds the id to the cancel subscription estimate params
func (o *CancelSubscriptionEstimateParams) WithID(id string) *CancelSubscriptionEstimateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cancel subscription estimate params
func (o *CancelSubscriptionEstimateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CancelSubscriptionEstimateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EstimateCancelSubscriptionRequest != nil {
		if err := r.SetBodyParam(o.EstimateCancelSubscriptionRequest); err != nil {
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
