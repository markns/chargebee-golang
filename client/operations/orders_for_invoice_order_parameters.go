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

// NewOrdersForInvoiceOrderParams creates a new OrdersForInvoiceOrderParams object
// with the default values initialized.
func NewOrdersForInvoiceOrderParams() *OrdersForInvoiceOrderParams {
	var ()
	return &OrdersForInvoiceOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrdersForInvoiceOrderParamsWithTimeout creates a new OrdersForInvoiceOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrdersForInvoiceOrderParamsWithTimeout(timeout time.Duration) *OrdersForInvoiceOrderParams {
	var ()
	return &OrdersForInvoiceOrderParams{

		timeout: timeout,
	}
}

// NewOrdersForInvoiceOrderParamsWithContext creates a new OrdersForInvoiceOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrdersForInvoiceOrderParamsWithContext(ctx context.Context) *OrdersForInvoiceOrderParams {
	var ()
	return &OrdersForInvoiceOrderParams{

		Context: ctx,
	}
}

// NewOrdersForInvoiceOrderParamsWithHTTPClient creates a new OrdersForInvoiceOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrdersForInvoiceOrderParamsWithHTTPClient(client *http.Client) *OrdersForInvoiceOrderParams {
	var ()
	return &OrdersForInvoiceOrderParams{
		HTTPClient: client,
	}
}

/*OrdersForInvoiceOrderParams contains all the parameters to send to the API endpoint
for the orders for invoice order operation typically these are written to a http.Request
*/
type OrdersForInvoiceOrderParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the orders for invoice order params
func (o *OrdersForInvoiceOrderParams) WithTimeout(timeout time.Duration) *OrdersForInvoiceOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the orders for invoice order params
func (o *OrdersForInvoiceOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the orders for invoice order params
func (o *OrdersForInvoiceOrderParams) WithContext(ctx context.Context) *OrdersForInvoiceOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the orders for invoice order params
func (o *OrdersForInvoiceOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the orders for invoice order params
func (o *OrdersForInvoiceOrderParams) WithHTTPClient(client *http.Client) *OrdersForInvoiceOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the orders for invoice order params
func (o *OrdersForInvoiceOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the orders for invoice order params
func (o *OrdersForInvoiceOrderParams) WithID(id string) *OrdersForInvoiceOrderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the orders for invoice order params
func (o *OrdersForInvoiceOrderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrdersForInvoiceOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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