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

// NewCopyPlanParams creates a new CopyPlanParams object
// with the default values initialized.
func NewCopyPlanParams() *CopyPlanParams {
	var ()
	return &CopyPlanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCopyPlanParamsWithTimeout creates a new CopyPlanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCopyPlanParamsWithTimeout(timeout time.Duration) *CopyPlanParams {
	var ()
	return &CopyPlanParams{

		timeout: timeout,
	}
}

// NewCopyPlanParamsWithContext creates a new CopyPlanParams object
// with the default values initialized, and the ability to set a context for a request
func NewCopyPlanParamsWithContext(ctx context.Context) *CopyPlanParams {
	var ()
	return &CopyPlanParams{

		Context: ctx,
	}
}

// NewCopyPlanParamsWithHTTPClient creates a new CopyPlanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCopyPlanParamsWithHTTPClient(client *http.Client) *CopyPlanParams {
	var ()
	return &CopyPlanParams{
		HTTPClient: client,
	}
}

/*CopyPlanParams contains all the parameters to send to the API endpoint
for the copy plan operation typically these are written to a http.Request
*/
type CopyPlanParams struct {

	/*PlanCopyRequest*/
	PlanCopyRequest *models.PlanCopyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the copy plan params
func (o *CopyPlanParams) WithTimeout(timeout time.Duration) *CopyPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the copy plan params
func (o *CopyPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the copy plan params
func (o *CopyPlanParams) WithContext(ctx context.Context) *CopyPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the copy plan params
func (o *CopyPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the copy plan params
func (o *CopyPlanParams) WithHTTPClient(client *http.Client) *CopyPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the copy plan params
func (o *CopyPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlanCopyRequest adds the planCopyRequest to the copy plan params
func (o *CopyPlanParams) WithPlanCopyRequest(planCopyRequest *models.PlanCopyRequest) *CopyPlanParams {
	o.SetPlanCopyRequest(planCopyRequest)
	return o
}

// SetPlanCopyRequest adds the planCopyRequest to the copy plan params
func (o *CopyPlanParams) SetPlanCopyRequest(planCopyRequest *models.PlanCopyRequest) {
	o.PlanCopyRequest = planCopyRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CopyPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PlanCopyRequest == nil {
		o.PlanCopyRequest = new(models.PlanCopyRequest)
	}

	if err := r.SetBodyParam(o.PlanCopyRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
