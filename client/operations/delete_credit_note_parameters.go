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

// NewDeleteCreditNoteParams creates a new DeleteCreditNoteParams object
// with the default values initialized.
func NewDeleteCreditNoteParams() *DeleteCreditNoteParams {
	var ()
	return &DeleteCreditNoteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCreditNoteParamsWithTimeout creates a new DeleteCreditNoteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCreditNoteParamsWithTimeout(timeout time.Duration) *DeleteCreditNoteParams {
	var ()
	return &DeleteCreditNoteParams{

		timeout: timeout,
	}
}

// NewDeleteCreditNoteParamsWithContext creates a new DeleteCreditNoteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCreditNoteParamsWithContext(ctx context.Context) *DeleteCreditNoteParams {
	var ()
	return &DeleteCreditNoteParams{

		Context: ctx,
	}
}

// NewDeleteCreditNoteParamsWithHTTPClient creates a new DeleteCreditNoteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCreditNoteParamsWithHTTPClient(client *http.Client) *DeleteCreditNoteParams {
	var ()
	return &DeleteCreditNoteParams{
		HTTPClient: client,
	}
}

/*DeleteCreditNoteParams contains all the parameters to send to the API endpoint
for the delete credit note operation typically these are written to a http.Request
*/
type DeleteCreditNoteParams struct {

	/*CreditNoteDeleteRequest*/
	CreditNoteDeleteRequest *models.CreditNoteDeleteRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete credit note params
func (o *DeleteCreditNoteParams) WithTimeout(timeout time.Duration) *DeleteCreditNoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete credit note params
func (o *DeleteCreditNoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete credit note params
func (o *DeleteCreditNoteParams) WithContext(ctx context.Context) *DeleteCreditNoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete credit note params
func (o *DeleteCreditNoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete credit note params
func (o *DeleteCreditNoteParams) WithHTTPClient(client *http.Client) *DeleteCreditNoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete credit note params
func (o *DeleteCreditNoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreditNoteDeleteRequest adds the creditNoteDeleteRequest to the delete credit note params
func (o *DeleteCreditNoteParams) WithCreditNoteDeleteRequest(creditNoteDeleteRequest *models.CreditNoteDeleteRequest) *DeleteCreditNoteParams {
	o.SetCreditNoteDeleteRequest(creditNoteDeleteRequest)
	return o
}

// SetCreditNoteDeleteRequest adds the creditNoteDeleteRequest to the delete credit note params
func (o *DeleteCreditNoteParams) SetCreditNoteDeleteRequest(creditNoteDeleteRequest *models.CreditNoteDeleteRequest) {
	o.CreditNoteDeleteRequest = creditNoteDeleteRequest
}

// WithID adds the id to the delete credit note params
func (o *DeleteCreditNoteParams) WithID(id string) *DeleteCreditNoteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete credit note params
func (o *DeleteCreditNoteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCreditNoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreditNoteDeleteRequest != nil {
		if err := r.SetBodyParam(o.CreditNoteDeleteRequest); err != nil {
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
