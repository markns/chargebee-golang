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

	"github.com/markns/chargebee-golang/models"
)

// NewTravelForwardTimeMachineParams creates a new TravelForwardTimeMachineParams object
// with the default values initialized.
func NewTravelForwardTimeMachineParams() *TravelForwardTimeMachineParams {
	var ()
	return &TravelForwardTimeMachineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTravelForwardTimeMachineParamsWithTimeout creates a new TravelForwardTimeMachineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTravelForwardTimeMachineParamsWithTimeout(timeout time.Duration) *TravelForwardTimeMachineParams {
	var ()
	return &TravelForwardTimeMachineParams{

		timeout: timeout,
	}
}

// NewTravelForwardTimeMachineParamsWithContext creates a new TravelForwardTimeMachineParams object
// with the default values initialized, and the ability to set a context for a request
func NewTravelForwardTimeMachineParamsWithContext(ctx context.Context) *TravelForwardTimeMachineParams {
	var ()
	return &TravelForwardTimeMachineParams{

		Context: ctx,
	}
}

// NewTravelForwardTimeMachineParamsWithHTTPClient creates a new TravelForwardTimeMachineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTravelForwardTimeMachineParamsWithHTTPClient(client *http.Client) *TravelForwardTimeMachineParams {
	var ()
	return &TravelForwardTimeMachineParams{
		HTTPClient: client,
	}
}

/*TravelForwardTimeMachineParams contains all the parameters to send to the API endpoint
for the travel forward time machine operation typically these are written to a http.Request
*/
type TravelForwardTimeMachineParams struct {

	/*TimeMachineTravelForwardRequest*/
	TimeMachineTravelForwardRequest *models.TimeMachineTravelForwardRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the travel forward time machine params
func (o *TravelForwardTimeMachineParams) WithTimeout(timeout time.Duration) *TravelForwardTimeMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the travel forward time machine params
func (o *TravelForwardTimeMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the travel forward time machine params
func (o *TravelForwardTimeMachineParams) WithContext(ctx context.Context) *TravelForwardTimeMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the travel forward time machine params
func (o *TravelForwardTimeMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the travel forward time machine params
func (o *TravelForwardTimeMachineParams) WithHTTPClient(client *http.Client) *TravelForwardTimeMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the travel forward time machine params
func (o *TravelForwardTimeMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimeMachineTravelForwardRequest adds the timeMachineTravelForwardRequest to the travel forward time machine params
func (o *TravelForwardTimeMachineParams) WithTimeMachineTravelForwardRequest(timeMachineTravelForwardRequest *models.TimeMachineTravelForwardRequest) *TravelForwardTimeMachineParams {
	o.SetTimeMachineTravelForwardRequest(timeMachineTravelForwardRequest)
	return o
}

// SetTimeMachineTravelForwardRequest adds the timeMachineTravelForwardRequest to the travel forward time machine params
func (o *TravelForwardTimeMachineParams) SetTimeMachineTravelForwardRequest(timeMachineTravelForwardRequest *models.TimeMachineTravelForwardRequest) {
	o.TimeMachineTravelForwardRequest = timeMachineTravelForwardRequest
}

// WithID adds the id to the travel forward time machine params
func (o *TravelForwardTimeMachineParams) WithID(id string) *TravelForwardTimeMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the travel forward time machine params
func (o *TravelForwardTimeMachineParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TravelForwardTimeMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TimeMachineTravelForwardRequest == nil {
		o.TimeMachineTravelForwardRequest = new(models.TimeMachineTravelForwardRequest)
	}

	if err := r.SetBodyParam(o.TimeMachineTravelForwardRequest); err != nil {
		return err
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
