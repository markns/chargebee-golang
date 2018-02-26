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

// NewListCreditNoteParams creates a new ListCreditNoteParams object
// with the default values initialized.
func NewListCreditNoteParams() *ListCreditNoteParams {

	return &ListCreditNoteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCreditNoteParamsWithTimeout creates a new ListCreditNoteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCreditNoteParamsWithTimeout(timeout time.Duration) *ListCreditNoteParams {

	return &ListCreditNoteParams{

		timeout: timeout,
	}
}

// NewListCreditNoteParamsWithContext creates a new ListCreditNoteParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCreditNoteParamsWithContext(ctx context.Context) *ListCreditNoteParams {

	return &ListCreditNoteParams{

		Context: ctx,
	}
}

// NewListCreditNoteParamsWithHTTPClient creates a new ListCreditNoteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCreditNoteParamsWithHTTPClient(client *http.Client) *ListCreditNoteParams {

	return &ListCreditNoteParams{
		HTTPClient: client,
	}
}

/*ListCreditNoteParams contains all the parameters to send to the API endpoint
for the list credit note operation typically these are written to a http.Request
*/
type ListCreditNoteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list credit note params
func (o *ListCreditNoteParams) WithTimeout(timeout time.Duration) *ListCreditNoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list credit note params
func (o *ListCreditNoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list credit note params
func (o *ListCreditNoteParams) WithContext(ctx context.Context) *ListCreditNoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list credit note params
func (o *ListCreditNoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list credit note params
func (o *ListCreditNoteParams) WithHTTPClient(client *http.Client) *ListCreditNoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list credit note params
func (o *ListCreditNoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListCreditNoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
