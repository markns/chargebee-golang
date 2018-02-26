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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListCouponSetParams creates a new ListCouponSetParams object
// with the default values initialized.
func NewListCouponSetParams() *ListCouponSetParams {
	var ()
	return &ListCouponSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCouponSetParamsWithTimeout creates a new ListCouponSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCouponSetParamsWithTimeout(timeout time.Duration) *ListCouponSetParams {
	var ()
	return &ListCouponSetParams{

		timeout: timeout,
	}
}

// NewListCouponSetParamsWithContext creates a new ListCouponSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCouponSetParamsWithContext(ctx context.Context) *ListCouponSetParams {
	var ()
	return &ListCouponSetParams{

		Context: ctx,
	}
}

// NewListCouponSetParamsWithHTTPClient creates a new ListCouponSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCouponSetParamsWithHTTPClient(client *http.Client) *ListCouponSetParams {
	var ()
	return &ListCouponSetParams{
		HTTPClient: client,
	}
}

/*ListCouponSetParams contains all the parameters to send to the API endpoint
for the list coupon set operation typically these are written to a http.Request
*/
type ListCouponSetParams struct {

	/*Limit*/
	Limit *int32
	/*Offset*/
	Offset *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list coupon set params
func (o *ListCouponSetParams) WithTimeout(timeout time.Duration) *ListCouponSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list coupon set params
func (o *ListCouponSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list coupon set params
func (o *ListCouponSetParams) WithContext(ctx context.Context) *ListCouponSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list coupon set params
func (o *ListCouponSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list coupon set params
func (o *ListCouponSetParams) WithHTTPClient(client *http.Client) *ListCouponSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list coupon set params
func (o *ListCouponSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list coupon set params
func (o *ListCouponSetParams) WithLimit(limit *int32) *ListCouponSetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list coupon set params
func (o *ListCouponSetParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the list coupon set params
func (o *ListCouponSetParams) WithOffset(offset *string) *ListCouponSetParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list coupon set params
func (o *ListCouponSetParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *ListCouponSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}