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

// NewCreateCreditNoteParams creates a new CreateCreditNoteParams object
// with the default values initialized.
func NewCreateCreditNoteParams() *CreateCreditNoteParams {
	var ()
	return &CreateCreditNoteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCreditNoteParamsWithTimeout creates a new CreateCreditNoteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCreditNoteParamsWithTimeout(timeout time.Duration) *CreateCreditNoteParams {
	var ()
	return &CreateCreditNoteParams{

		timeout: timeout,
	}
}

// NewCreateCreditNoteParamsWithContext creates a new CreateCreditNoteParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCreditNoteParamsWithContext(ctx context.Context) *CreateCreditNoteParams {
	var ()
	return &CreateCreditNoteParams{

		Context: ctx,
	}
}

// NewCreateCreditNoteParamsWithHTTPClient creates a new CreateCreditNoteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCreditNoteParamsWithHTTPClient(client *http.Client) *CreateCreditNoteParams {
	var ()
	return &CreateCreditNoteParams{
		HTTPClient: client,
	}
}

/*CreateCreditNoteParams contains all the parameters to send to the API endpoint
for the create credit note operation typically these are written to a http.Request
*/
type CreateCreditNoteParams struct {

	/*CreditNoteCreateRequest*/
	CreditNoteCreateRequest *models.CreditNoteCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create credit note params
func (o *CreateCreditNoteParams) WithTimeout(timeout time.Duration) *CreateCreditNoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create credit note params
func (o *CreateCreditNoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create credit note params
func (o *CreateCreditNoteParams) WithContext(ctx context.Context) *CreateCreditNoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create credit note params
func (o *CreateCreditNoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create credit note params
func (o *CreateCreditNoteParams) WithHTTPClient(client *http.Client) *CreateCreditNoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create credit note params
func (o *CreateCreditNoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreditNoteCreateRequest adds the creditNoteCreateRequest to the create credit note params
func (o *CreateCreditNoteParams) WithCreditNoteCreateRequest(creditNoteCreateRequest *models.CreditNoteCreateRequest) *CreateCreditNoteParams {
	o.SetCreditNoteCreateRequest(creditNoteCreateRequest)
	return o
}

// SetCreditNoteCreateRequest adds the creditNoteCreateRequest to the create credit note params
func (o *CreateCreditNoteParams) SetCreditNoteCreateRequest(creditNoteCreateRequest *models.CreditNoteCreateRequest) {
	o.CreditNoteCreateRequest = creditNoteCreateRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCreditNoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreditNoteCreateRequest != nil {
		if err := r.SetBodyParam(o.CreditNoteCreateRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
