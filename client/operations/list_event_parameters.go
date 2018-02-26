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

// NewListEventParams creates a new ListEventParams object
// with the default values initialized.
func NewListEventParams() *ListEventParams {
	var ()
	return &ListEventParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListEventParamsWithTimeout creates a new ListEventParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListEventParamsWithTimeout(timeout time.Duration) *ListEventParams {
	var ()
	return &ListEventParams{

		timeout: timeout,
	}
}

// NewListEventParamsWithContext creates a new ListEventParams object
// with the default values initialized, and the ability to set a context for a request
func NewListEventParamsWithContext(ctx context.Context) *ListEventParams {
	var ()
	return &ListEventParams{

		Context: ctx,
	}
}

// NewListEventParamsWithHTTPClient creates a new ListEventParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListEventParamsWithHTTPClient(client *http.Client) *ListEventParams {
	var ()
	return &ListEventParams{
		HTTPClient: client,
	}
}

/*ListEventParams contains all the parameters to send to the API endpoint
for the list event operation typically these are written to a http.Request
*/
type ListEventParams struct {

	/*Limit*/
	Limit *int32
	/*Offset*/
	Offset *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list event params
func (o *ListEventParams) WithTimeout(timeout time.Duration) *ListEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list event params
func (o *ListEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list event params
func (o *ListEventParams) WithContext(ctx context.Context) *ListEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list event params
func (o *ListEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list event params
func (o *ListEventParams) WithHTTPClient(client *http.Client) *ListEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list event params
func (o *ListEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list event params
func (o *ListEventParams) WithLimit(limit *int32) *ListEventParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list event params
func (o *ListEventParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the list event params
func (o *ListEventParams) WithOffset(offset *string) *ListEventParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list event params
func (o *ListEventParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *ListEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
