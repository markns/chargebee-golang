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

// NewCreateCouponSetParams creates a new CreateCouponSetParams object
// with the default values initialized.
func NewCreateCouponSetParams() *CreateCouponSetParams {
	var ()
	return &CreateCouponSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCouponSetParamsWithTimeout creates a new CreateCouponSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCouponSetParamsWithTimeout(timeout time.Duration) *CreateCouponSetParams {
	var ()
	return &CreateCouponSetParams{

		timeout: timeout,
	}
}

// NewCreateCouponSetParamsWithContext creates a new CreateCouponSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCouponSetParamsWithContext(ctx context.Context) *CreateCouponSetParams {
	var ()
	return &CreateCouponSetParams{

		Context: ctx,
	}
}

// NewCreateCouponSetParamsWithHTTPClient creates a new CreateCouponSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCouponSetParamsWithHTTPClient(client *http.Client) *CreateCouponSetParams {
	var ()
	return &CreateCouponSetParams{
		HTTPClient: client,
	}
}

/*CreateCouponSetParams contains all the parameters to send to the API endpoint
for the create coupon set operation typically these are written to a http.Request
*/
type CreateCouponSetParams struct {

	/*CouponSetCreateRequest*/
	CouponSetCreateRequest *models.CouponSetCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create coupon set params
func (o *CreateCouponSetParams) WithTimeout(timeout time.Duration) *CreateCouponSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create coupon set params
func (o *CreateCouponSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create coupon set params
func (o *CreateCouponSetParams) WithContext(ctx context.Context) *CreateCouponSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create coupon set params
func (o *CreateCouponSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create coupon set params
func (o *CreateCouponSetParams) WithHTTPClient(client *http.Client) *CreateCouponSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create coupon set params
func (o *CreateCouponSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCouponSetCreateRequest adds the couponSetCreateRequest to the create coupon set params
func (o *CreateCouponSetParams) WithCouponSetCreateRequest(couponSetCreateRequest *models.CouponSetCreateRequest) *CreateCouponSetParams {
	o.SetCouponSetCreateRequest(couponSetCreateRequest)
	return o
}

// SetCouponSetCreateRequest adds the couponSetCreateRequest to the create coupon set params
func (o *CreateCouponSetParams) SetCouponSetCreateRequest(couponSetCreateRequest *models.CouponSetCreateRequest) {
	o.CouponSetCreateRequest = couponSetCreateRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCouponSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CouponSetCreateRequest == nil {
		o.CouponSetCreateRequest = new(models.CouponSetCreateRequest)
	}

	if err := r.SetBodyParam(o.CouponSetCreateRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}