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

// NewRenewalEstimateEstimateParams creates a new RenewalEstimateEstimateParams object
// with the default values initialized.
func NewRenewalEstimateEstimateParams() *RenewalEstimateEstimateParams {
	var ()
	return &RenewalEstimateEstimateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRenewalEstimateEstimateParamsWithTimeout creates a new RenewalEstimateEstimateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRenewalEstimateEstimateParamsWithTimeout(timeout time.Duration) *RenewalEstimateEstimateParams {
	var ()
	return &RenewalEstimateEstimateParams{

		timeout: timeout,
	}
}

// NewRenewalEstimateEstimateParamsWithContext creates a new RenewalEstimateEstimateParams object
// with the default values initialized, and the ability to set a context for a request
func NewRenewalEstimateEstimateParamsWithContext(ctx context.Context) *RenewalEstimateEstimateParams {
	var ()
	return &RenewalEstimateEstimateParams{

		Context: ctx,
	}
}

// NewRenewalEstimateEstimateParamsWithHTTPClient creates a new RenewalEstimateEstimateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRenewalEstimateEstimateParamsWithHTTPClient(client *http.Client) *RenewalEstimateEstimateParams {
	var ()
	return &RenewalEstimateEstimateParams{
		HTTPClient: client,
	}
}

/*RenewalEstimateEstimateParams contains all the parameters to send to the API endpoint
for the renewal estimate estimate operation typically these are written to a http.Request
*/
type RenewalEstimateEstimateParams struct {

	/*EstimateRenewalEstimateRequest*/
	EstimateRenewalEstimateRequest *models.EstimateRenewalEstimateRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the renewal estimate estimate params
func (o *RenewalEstimateEstimateParams) WithTimeout(timeout time.Duration) *RenewalEstimateEstimateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the renewal estimate estimate params
func (o *RenewalEstimateEstimateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the renewal estimate estimate params
func (o *RenewalEstimateEstimateParams) WithContext(ctx context.Context) *RenewalEstimateEstimateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the renewal estimate estimate params
func (o *RenewalEstimateEstimateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the renewal estimate estimate params
func (o *RenewalEstimateEstimateParams) WithHTTPClient(client *http.Client) *RenewalEstimateEstimateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the renewal estimate estimate params
func (o *RenewalEstimateEstimateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEstimateRenewalEstimateRequest adds the estimateRenewalEstimateRequest to the renewal estimate estimate params
func (o *RenewalEstimateEstimateParams) WithEstimateRenewalEstimateRequest(estimateRenewalEstimateRequest *models.EstimateRenewalEstimateRequest) *RenewalEstimateEstimateParams {
	o.SetEstimateRenewalEstimateRequest(estimateRenewalEstimateRequest)
	return o
}

// SetEstimateRenewalEstimateRequest adds the estimateRenewalEstimateRequest to the renewal estimate estimate params
func (o *RenewalEstimateEstimateParams) SetEstimateRenewalEstimateRequest(estimateRenewalEstimateRequest *models.EstimateRenewalEstimateRequest) {
	o.EstimateRenewalEstimateRequest = estimateRenewalEstimateRequest
}

// WithID adds the id to the renewal estimate estimate params
func (o *RenewalEstimateEstimateParams) WithID(id string) *RenewalEstimateEstimateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the renewal estimate estimate params
func (o *RenewalEstimateEstimateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RenewalEstimateEstimateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EstimateRenewalEstimateRequest != nil {
		if err := r.SetBodyParam(o.EstimateRenewalEstimateRequest); err != nil {
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
