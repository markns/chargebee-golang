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

// NewUpdateCouponSetParams creates a new UpdateCouponSetParams object
// with the default values initialized.
func NewUpdateCouponSetParams() *UpdateCouponSetParams {
	var ()
	return &UpdateCouponSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCouponSetParamsWithTimeout creates a new UpdateCouponSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCouponSetParamsWithTimeout(timeout time.Duration) *UpdateCouponSetParams {
	var ()
	return &UpdateCouponSetParams{

		timeout: timeout,
	}
}

// NewUpdateCouponSetParamsWithContext creates a new UpdateCouponSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCouponSetParamsWithContext(ctx context.Context) *UpdateCouponSetParams {
	var ()
	return &UpdateCouponSetParams{

		Context: ctx,
	}
}

// NewUpdateCouponSetParamsWithHTTPClient creates a new UpdateCouponSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCouponSetParamsWithHTTPClient(client *http.Client) *UpdateCouponSetParams {
	var ()
	return &UpdateCouponSetParams{
		HTTPClient: client,
	}
}

/*UpdateCouponSetParams contains all the parameters to send to the API endpoint
for the update coupon set operation typically these are written to a http.Request
*/
type UpdateCouponSetParams struct {

	/*CouponSetUpdateRequest*/
	CouponSetUpdateRequest *models.CouponSetUpdateRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update coupon set params
func (o *UpdateCouponSetParams) WithTimeout(timeout time.Duration) *UpdateCouponSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update coupon set params
func (o *UpdateCouponSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update coupon set params
func (o *UpdateCouponSetParams) WithContext(ctx context.Context) *UpdateCouponSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update coupon set params
func (o *UpdateCouponSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update coupon set params
func (o *UpdateCouponSetParams) WithHTTPClient(client *http.Client) *UpdateCouponSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update coupon set params
func (o *UpdateCouponSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCouponSetUpdateRequest adds the couponSetUpdateRequest to the update coupon set params
func (o *UpdateCouponSetParams) WithCouponSetUpdateRequest(couponSetUpdateRequest *models.CouponSetUpdateRequest) *UpdateCouponSetParams {
	o.SetCouponSetUpdateRequest(couponSetUpdateRequest)
	return o
}

// SetCouponSetUpdateRequest adds the couponSetUpdateRequest to the update coupon set params
func (o *UpdateCouponSetParams) SetCouponSetUpdateRequest(couponSetUpdateRequest *models.CouponSetUpdateRequest) {
	o.CouponSetUpdateRequest = couponSetUpdateRequest
}

// WithID adds the id to the update coupon set params
func (o *UpdateCouponSetParams) WithID(id string) *UpdateCouponSetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update coupon set params
func (o *UpdateCouponSetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCouponSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CouponSetUpdateRequest != nil {
		if err := r.SetBodyParam(o.CouponSetUpdateRequest); err != nil {
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
