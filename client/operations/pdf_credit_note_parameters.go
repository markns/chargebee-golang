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

// NewPdfCreditNoteParams creates a new PdfCreditNoteParams object
// with the default values initialized.
func NewPdfCreditNoteParams() *PdfCreditNoteParams {
	var ()
	return &PdfCreditNoteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPdfCreditNoteParamsWithTimeout creates a new PdfCreditNoteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPdfCreditNoteParamsWithTimeout(timeout time.Duration) *PdfCreditNoteParams {
	var ()
	return &PdfCreditNoteParams{

		timeout: timeout,
	}
}

// NewPdfCreditNoteParamsWithContext creates a new PdfCreditNoteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPdfCreditNoteParamsWithContext(ctx context.Context) *PdfCreditNoteParams {
	var ()
	return &PdfCreditNoteParams{

		Context: ctx,
	}
}

// NewPdfCreditNoteParamsWithHTTPClient creates a new PdfCreditNoteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPdfCreditNoteParamsWithHTTPClient(client *http.Client) *PdfCreditNoteParams {
	var ()
	return &PdfCreditNoteParams{
		HTTPClient: client,
	}
}

/*PdfCreditNoteParams contains all the parameters to send to the API endpoint
for the pdf credit note operation typically these are written to a http.Request
*/
type PdfCreditNoteParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pdf credit note params
func (o *PdfCreditNoteParams) WithTimeout(timeout time.Duration) *PdfCreditNoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pdf credit note params
func (o *PdfCreditNoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pdf credit note params
func (o *PdfCreditNoteParams) WithContext(ctx context.Context) *PdfCreditNoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pdf credit note params
func (o *PdfCreditNoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pdf credit note params
func (o *PdfCreditNoteParams) WithHTTPClient(client *http.Client) *PdfCreditNoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pdf credit note params
func (o *PdfCreditNoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the pdf credit note params
func (o *PdfCreditNoteParams) WithID(id string) *PdfCreditNoteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the pdf credit note params
func (o *PdfCreditNoteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PdfCreditNoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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