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
)

// NewDeleteCouponParams creates a new DeleteCouponParams object
// with the default values initialized.
func NewDeleteCouponParams() *DeleteCouponParams {
	var ()
	return &DeleteCouponParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCouponParamsWithTimeout creates a new DeleteCouponParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCouponParamsWithTimeout(timeout time.Duration) *DeleteCouponParams {
	var ()
	return &DeleteCouponParams{

		timeout: timeout,
	}
}

// NewDeleteCouponParamsWithContext creates a new DeleteCouponParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCouponParamsWithContext(ctx context.Context) *DeleteCouponParams {
	var ()
	return &DeleteCouponParams{

		Context: ctx,
	}
}

// NewDeleteCouponParamsWithHTTPClient creates a new DeleteCouponParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCouponParamsWithHTTPClient(client *http.Client) *DeleteCouponParams {
	var ()
	return &DeleteCouponParams{
		HTTPClient: client,
	}
}

/*DeleteCouponParams contains all the parameters to send to the API endpoint
for the delete coupon operation typically these are written to a http.Request
*/
type DeleteCouponParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete coupon params
func (o *DeleteCouponParams) WithTimeout(timeout time.Duration) *DeleteCouponParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete coupon params
func (o *DeleteCouponParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete coupon params
func (o *DeleteCouponParams) WithContext(ctx context.Context) *DeleteCouponParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete coupon params
func (o *DeleteCouponParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete coupon params
func (o *DeleteCouponParams) WithHTTPClient(client *http.Client) *DeleteCouponParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete coupon params
func (o *DeleteCouponParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete coupon params
func (o *DeleteCouponParams) WithID(id string) *DeleteCouponParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete coupon params
func (o *DeleteCouponParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCouponParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
