// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/markns/chargebee-golang/models"
)

// ChargeAddonAtTermEndSubscriptionReader is a Reader for the ChargeAddonAtTermEndSubscription structure.
type ChargeAddonAtTermEndSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChargeAddonAtTermEndSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChargeAddonAtTermEndSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChargeAddonAtTermEndSubscriptionOK creates a ChargeAddonAtTermEndSubscriptionOK with default headers values
func NewChargeAddonAtTermEndSubscriptionOK() *ChargeAddonAtTermEndSubscriptionOK {
	return &ChargeAddonAtTermEndSubscriptionOK{}
}

/*ChargeAddonAtTermEndSubscriptionOK handles this case with default header values.

chargeAddonAtTermEndSubscription response
*/
type ChargeAddonAtTermEndSubscriptionOK struct {
	Payload *models.SubscriptionResponse
}

func (o *ChargeAddonAtTermEndSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{id}/charge_addon_at_term_end][%d] chargeAddonAtTermEndSubscriptionOK  %+v", 200, o.Payload)
}

func (o *ChargeAddonAtTermEndSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}