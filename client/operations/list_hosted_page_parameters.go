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

// NewListHostedPageParams creates a new ListHostedPageParams object
// with the default values initialized.
func NewListHostedPageParams() *ListHostedPageParams {

	return &ListHostedPageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListHostedPageParamsWithTimeout creates a new ListHostedPageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListHostedPageParamsWithTimeout(timeout time.Duration) *ListHostedPageParams {

	return &ListHostedPageParams{

		timeout: timeout,
	}
}

// NewListHostedPageParamsWithContext creates a new ListHostedPageParams object
// with the default values initialized, and the ability to set a context for a request
func NewListHostedPageParamsWithContext(ctx context.Context) *ListHostedPageParams {

	return &ListHostedPageParams{

		Context: ctx,
	}
}

// NewListHostedPageParamsWithHTTPClient creates a new ListHostedPageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListHostedPageParamsWithHTTPClient(client *http.Client) *ListHostedPageParams {

	return &ListHostedPageParams{
		HTTPClient: client,
	}
}

/*ListHostedPageParams contains all the parameters to send to the API endpoint
for the list hosted page operation typically these are written to a http.Request
*/
type ListHostedPageParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list hosted page params
func (o *ListHostedPageParams) WithTimeout(timeout time.Duration) *ListHostedPageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list hosted page params
func (o *ListHostedPageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list hosted page params
func (o *ListHostedPageParams) WithContext(ctx context.Context) *ListHostedPageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list hosted page params
func (o *ListHostedPageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list hosted page params
func (o *ListHostedPageParams) WithHTTPClient(client *http.Client) *ListHostedPageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list hosted page params
func (o *ListHostedPageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListHostedPageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
