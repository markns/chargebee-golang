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

// NewUnarchiveCouponParams creates a new UnarchiveCouponParams object
// with the default values initialized.
func NewUnarchiveCouponParams() *UnarchiveCouponParams {
	var ()
	return &UnarchiveCouponParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUnarchiveCouponParamsWithTimeout creates a new UnarchiveCouponParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnarchiveCouponParamsWithTimeout(timeout time.Duration) *UnarchiveCouponParams {
	var ()
	return &UnarchiveCouponParams{

		timeout: timeout,
	}
}

// NewUnarchiveCouponParamsWithContext creates a new UnarchiveCouponParams object
// with the default values initialized, and the ability to set a context for a request
func NewUnarchiveCouponParamsWithContext(ctx context.Context) *UnarchiveCouponParams {
	var ()
	return &UnarchiveCouponParams{

		Context: ctx,
	}
}

// NewUnarchiveCouponParamsWithHTTPClient creates a new UnarchiveCouponParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUnarchiveCouponParamsWithHTTPClient(client *http.Client) *UnarchiveCouponParams {
	var ()
	return &UnarchiveCouponParams{
		HTTPClient: client,
	}
}

/*UnarchiveCouponParams contains all the parameters to send to the API endpoint
for the unarchive coupon operation typically these are written to a http.Request
*/
type UnarchiveCouponParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the unarchive coupon params
func (o *UnarchiveCouponParams) WithTimeout(timeout time.Duration) *UnarchiveCouponParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unarchive coupon params
func (o *UnarchiveCouponParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unarchive coupon params
func (o *UnarchiveCouponParams) WithContext(ctx context.Context) *UnarchiveCouponParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unarchive coupon params
func (o *UnarchiveCouponParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unarchive coupon params
func (o *UnarchiveCouponParams) WithHTTPClient(client *http.Client) *UnarchiveCouponParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unarchive coupon params
func (o *UnarchiveCouponParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the unarchive coupon params
func (o *UnarchiveCouponParams) WithID(id string) *UnarchiveCouponParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the unarchive coupon params
func (o *UnarchiveCouponParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UnarchiveCouponParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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