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

// NewListAddonParams creates a new ListAddonParams object
// with the default values initialized.
func NewListAddonParams() *ListAddonParams {
	var ()
	return &ListAddonParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAddonParamsWithTimeout creates a new ListAddonParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAddonParamsWithTimeout(timeout time.Duration) *ListAddonParams {
	var ()
	return &ListAddonParams{

		timeout: timeout,
	}
}

// NewListAddonParamsWithContext creates a new ListAddonParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAddonParamsWithContext(ctx context.Context) *ListAddonParams {
	var ()
	return &ListAddonParams{

		Context: ctx,
	}
}

// NewListAddonParamsWithHTTPClient creates a new ListAddonParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAddonParamsWithHTTPClient(client *http.Client) *ListAddonParams {
	var ()
	return &ListAddonParams{
		HTTPClient: client,
	}
}

/*ListAddonParams contains all the parameters to send to the API endpoint
for the list addon operation typically these are written to a http.Request
*/
type ListAddonParams struct {

	/*Limit*/
	Limit *int32
	/*Offset*/
	Offset *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list addon params
func (o *ListAddonParams) WithTimeout(timeout time.Duration) *ListAddonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list addon params
func (o *ListAddonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list addon params
func (o *ListAddonParams) WithContext(ctx context.Context) *ListAddonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list addon params
func (o *ListAddonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list addon params
func (o *ListAddonParams) WithHTTPClient(client *http.Client) *ListAddonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list addon params
func (o *ListAddonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list addon params
func (o *ListAddonParams) WithLimit(limit *int32) *ListAddonParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list addon params
func (o *ListAddonParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the list addon params
func (o *ListAddonParams) WithOffset(offset *string) *ListAddonParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list addon params
func (o *ListAddonParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *ListAddonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
