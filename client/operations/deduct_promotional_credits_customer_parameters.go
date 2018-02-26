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

// NewDeductPromotionalCreditsCustomerParams creates a new DeductPromotionalCreditsCustomerParams object
// with the default values initialized.
func NewDeductPromotionalCreditsCustomerParams() *DeductPromotionalCreditsCustomerParams {
	var ()
	return &DeductPromotionalCreditsCustomerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeductPromotionalCreditsCustomerParamsWithTimeout creates a new DeductPromotionalCreditsCustomerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeductPromotionalCreditsCustomerParamsWithTimeout(timeout time.Duration) *DeductPromotionalCreditsCustomerParams {
	var ()
	return &DeductPromotionalCreditsCustomerParams{

		timeout: timeout,
	}
}

// NewDeductPromotionalCreditsCustomerParamsWithContext creates a new DeductPromotionalCreditsCustomerParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeductPromotionalCreditsCustomerParamsWithContext(ctx context.Context) *DeductPromotionalCreditsCustomerParams {
	var ()
	return &DeductPromotionalCreditsCustomerParams{

		Context: ctx,
	}
}

// NewDeductPromotionalCreditsCustomerParamsWithHTTPClient creates a new DeductPromotionalCreditsCustomerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeductPromotionalCreditsCustomerParamsWithHTTPClient(client *http.Client) *DeductPromotionalCreditsCustomerParams {
	var ()
	return &DeductPromotionalCreditsCustomerParams{
		HTTPClient: client,
	}
}

/*DeductPromotionalCreditsCustomerParams contains all the parameters to send to the API endpoint
for the deduct promotional credits customer operation typically these are written to a http.Request
*/
type DeductPromotionalCreditsCustomerParams struct {

	/*CustomerDeductPromotionalCreditsRequest*/
	CustomerDeductPromotionalCreditsRequest *models.CustomerDeductPromotionalCreditsRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the deduct promotional credits customer params
func (o *DeductPromotionalCreditsCustomerParams) WithTimeout(timeout time.Duration) *DeductPromotionalCreditsCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deduct promotional credits customer params
func (o *DeductPromotionalCreditsCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deduct promotional credits customer params
func (o *DeductPromotionalCreditsCustomerParams) WithContext(ctx context.Context) *DeductPromotionalCreditsCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deduct promotional credits customer params
func (o *DeductPromotionalCreditsCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deduct promotional credits customer params
func (o *DeductPromotionalCreditsCustomerParams) WithHTTPClient(client *http.Client) *DeductPromotionalCreditsCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deduct promotional credits customer params
func (o *DeductPromotionalCreditsCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerDeductPromotionalCreditsRequest adds the customerDeductPromotionalCreditsRequest to the deduct promotional credits customer params
func (o *DeductPromotionalCreditsCustomerParams) WithCustomerDeductPromotionalCreditsRequest(customerDeductPromotionalCreditsRequest *models.CustomerDeductPromotionalCreditsRequest) *DeductPromotionalCreditsCustomerParams {
	o.SetCustomerDeductPromotionalCreditsRequest(customerDeductPromotionalCreditsRequest)
	return o
}

// SetCustomerDeductPromotionalCreditsRequest adds the customerDeductPromotionalCreditsRequest to the deduct promotional credits customer params
func (o *DeductPromotionalCreditsCustomerParams) SetCustomerDeductPromotionalCreditsRequest(customerDeductPromotionalCreditsRequest *models.CustomerDeductPromotionalCreditsRequest) {
	o.CustomerDeductPromotionalCreditsRequest = customerDeductPromotionalCreditsRequest
}

// WithID adds the id to the deduct promotional credits customer params
func (o *DeductPromotionalCreditsCustomerParams) WithID(id string) *DeductPromotionalCreditsCustomerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the deduct promotional credits customer params
func (o *DeductPromotionalCreditsCustomerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeductPromotionalCreditsCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CustomerDeductPromotionalCreditsRequest == nil {
		o.CustomerDeductPromotionalCreditsRequest = new(models.CustomerDeductPromotionalCreditsRequest)
	}

	if err := r.SetBodyParam(o.CustomerDeductPromotionalCreditsRequest); err != nil {
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