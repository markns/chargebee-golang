// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/markns/chargebee-golang/models"
)

// ChargeFutureRenewalsSubscriptionReader is a Reader for the ChargeFutureRenewalsSubscription structure.
type ChargeFutureRenewalsSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChargeFutureRenewalsSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChargeFutureRenewalsSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChargeFutureRenewalsSubscriptionOK creates a ChargeFutureRenewalsSubscriptionOK with default headers values
func NewChargeFutureRenewalsSubscriptionOK() *ChargeFutureRenewalsSubscriptionOK {
	return &ChargeFutureRenewalsSubscriptionOK{}
}

/*ChargeFutureRenewalsSubscriptionOK handles this case with default header values.

chargeFutureRenewalsSubscription response
*/
type ChargeFutureRenewalsSubscriptionOK struct {
	Payload *models.SubscriptionResponse
}

func (o *ChargeFutureRenewalsSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{id}/charge_future_renewals][%d] chargeFutureRenewalsSubscriptionOK  %+v", 200, o.Payload)
}

func (o *ChargeFutureRenewalsSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
