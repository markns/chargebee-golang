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

// NewCreateSubForCustomerEstimateEstimateParams creates a new CreateSubForCustomerEstimateEstimateParams object
// with the default values initialized.
func NewCreateSubForCustomerEstimateEstimateParams() *CreateSubForCustomerEstimateEstimateParams {
	var ()
	return &CreateSubForCustomerEstimateEstimateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSubForCustomerEstimateEstimateParamsWithTimeout creates a new CreateSubForCustomerEstimateEstimateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSubForCustomerEstimateEstimateParamsWithTimeout(timeout time.Duration) *CreateSubForCustomerEstimateEstimateParams {
	var ()
	return &CreateSubForCustomerEstimateEstimateParams{

		timeout: timeout,
	}
}

// NewCreateSubForCustomerEstimateEstimateParamsWithContext creates a new CreateSubForCustomerEstimateEstimateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSubForCustomerEstimateEstimateParamsWithContext(ctx context.Context) *CreateSubForCustomerEstimateEstimateParams {
	var ()
	return &CreateSubForCustomerEstimateEstimateParams{

		Context: ctx,
	}
}

// NewCreateSubForCustomerEstimateEstimateParamsWithHTTPClient creates a new CreateSubForCustomerEstimateEstimateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSubForCustomerEstimateEstimateParamsWithHTTPClient(client *http.Client) *CreateSubForCustomerEstimateEstimateParams {
	var ()
	return &CreateSubForCustomerEstimateEstimateParams{
		HTTPClient: client,
	}
}

/*CreateSubForCustomerEstimateEstimateParams contains all the parameters to send to the API endpoint
for the create sub for customer estimate estimate operation typically these are written to a http.Request
*/
type CreateSubForCustomerEstimateEstimateParams struct {

	/*EstimateCreateSubForCustomerEstimateRequest*/
	EstimateCreateSubForCustomerEstimateRequest *models.EstimateCreateSubForCustomerEstimateRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create sub for customer estimate estimate params
func (o *CreateSubForCustomerEstimateEstimateParams) WithTimeout(timeout time.Duration) *CreateSubForCustomerEstimateEstimateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create sub for customer estimate estimate params
func (o *CreateSubForCustomerEstimateEstimateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create sub for customer estimate estimate params
func (o *CreateSubForCustomerEstimateEstimateParams) WithContext(ctx context.Context) *CreateSubForCustomerEstimateEstimateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create sub for customer estimate estimate params
func (o *CreateSubForCustomerEstimateEstimateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create sub for customer estimate estimate params
func (o *CreateSubForCustomerEstimateEstimateParams) WithHTTPClient(client *http.Client) *CreateSubForCustomerEstimateEstimateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create sub for customer estimate estimate params
func (o *CreateSubForCustomerEstimateEstimateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEstimateCreateSubForCustomerEstimateRequest adds the estimateCreateSubForCustomerEstimateRequest to the create sub for customer estimate estimate params
func (o *CreateSubForCustomerEstimateEstimateParams) WithEstimateCreateSubForCustomerEstimateRequest(estimateCreateSubForCustomerEstimateRequest *models.EstimateCreateSubForCustomerEstimateRequest) *CreateSubForCustomerEstimateEstimateParams {
	o.SetEstimateCreateSubForCustomerEstimateRequest(estimateCreateSubForCustomerEstimateRequest)
	return o
}

// SetEstimateCreateSubForCustomerEstimateRequest adds the estimateCreateSubForCustomerEstimateRequest to the create sub for customer estimate estimate params
func (o *CreateSubForCustomerEstimateEstimateParams) SetEstimateCreateSubForCustomerEstimateRequest(estimateCreateSubForCustomerEstimateRequest *models.EstimateCreateSubForCustomerEstimateRequest) {
	o.EstimateCreateSubForCustomerEstimateRequest = estimateCreateSubForCustomerEstimateRequest
}

// WithID adds the id to the create sub for customer estimate estimate params
func (o *CreateSubForCustomerEstimateEstimateParams) WithID(id string) *CreateSubForCustomerEstimateEstimateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create sub for customer estimate estimate params
func (o *CreateSubForCustomerEstimateEstimateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSubForCustomerEstimateEstimateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EstimateCreateSubForCustomerEstimateRequest == nil {
		o.EstimateCreateSubForCustomerEstimateRequest = new(models.EstimateCreateSubForCustomerEstimateRequest)
	}

	if err := r.SetBodyParam(o.EstimateCreateSubForCustomerEstimateRequest); err != nil {
		return err
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