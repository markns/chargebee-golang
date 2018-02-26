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

// NewListTransactionParams creates a new ListTransactionParams object
// with the default values initialized.
func NewListTransactionParams() *ListTransactionParams {
	var ()
	return &ListTransactionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTransactionParamsWithTimeout creates a new ListTransactionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTransactionParamsWithTimeout(timeout time.Duration) *ListTransactionParams {
	var ()
	return &ListTransactionParams{

		timeout: timeout,
	}
}

// NewListTransactionParamsWithContext creates a new ListTransactionParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTransactionParamsWithContext(ctx context.Context) *ListTransactionParams {
	var ()
	return &ListTransactionParams{

		Context: ctx,
	}
}

// NewListTransactionParamsWithHTTPClient creates a new ListTransactionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTransactionParamsWithHTTPClient(client *http.Client) *ListTransactionParams {
	var ()
	return &ListTransactionParams{
		HTTPClient: client,
	}
}

/*ListTransactionParams contains all the parameters to send to the API endpoint
for the list transaction operation typically these are written to a http.Request
*/
type ListTransactionParams struct {

	/*Limit*/
	Limit *int32
	/*Offset*/
	Offset *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list transaction params
func (o *ListTransactionParams) WithTimeout(timeout time.Duration) *ListTransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list transaction params
func (o *ListTransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list transaction params
func (o *ListTransactionParams) WithContext(ctx context.Context) *ListTransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list transaction params
func (o *ListTransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list transaction params
func (o *ListTransactionParams) WithHTTPClient(client *http.Client) *ListTransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list transaction params
func (o *ListTransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list transaction params
func (o *ListTransactionParams) WithLimit(limit *int32) *ListTransactionParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list transaction params
func (o *ListTransactionParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the list transaction params
func (o *ListTransactionParams) WithOffset(offset *string) *ListTransactionParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list transaction params
func (o *ListTransactionParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *ListTransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
