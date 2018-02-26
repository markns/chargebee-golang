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

// NewRetrieveCommentParams creates a new RetrieveCommentParams object
// with the default values initialized.
func NewRetrieveCommentParams() *RetrieveCommentParams {
	var ()
	return &RetrieveCommentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveCommentParamsWithTimeout creates a new RetrieveCommentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveCommentParamsWithTimeout(timeout time.Duration) *RetrieveCommentParams {
	var ()
	return &RetrieveCommentParams{

		timeout: timeout,
	}
}

// NewRetrieveCommentParamsWithContext creates a new RetrieveCommentParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveCommentParamsWithContext(ctx context.Context) *RetrieveCommentParams {
	var ()
	return &RetrieveCommentParams{

		Context: ctx,
	}
}

// NewRetrieveCommentParamsWithHTTPClient creates a new RetrieveCommentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveCommentParamsWithHTTPClient(client *http.Client) *RetrieveCommentParams {
	var ()
	return &RetrieveCommentParams{
		HTTPClient: client,
	}
}

/*RetrieveCommentParams contains all the parameters to send to the API endpoint
for the retrieve comment operation typically these are written to a http.Request
*/
type RetrieveCommentParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve comment params
func (o *RetrieveCommentParams) WithTimeout(timeout time.Duration) *RetrieveCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve comment params
func (o *RetrieveCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve comment params
func (o *RetrieveCommentParams) WithContext(ctx context.Context) *RetrieveCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve comment params
func (o *RetrieveCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve comment params
func (o *RetrieveCommentParams) WithHTTPClient(client *http.Client) *RetrieveCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve comment params
func (o *RetrieveCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the retrieve comment params
func (o *RetrieveCommentParams) WithID(id string) *RetrieveCommentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the retrieve comment params
func (o *RetrieveCommentParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
