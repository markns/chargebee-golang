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

// CreateSubscriptionReader is a Reader for the CreateSubscription structure.
type CreateSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSubscriptionOK creates a CreateSubscriptionOK with default headers values
func NewCreateSubscriptionOK() *CreateSubscriptionOK {
	return &CreateSubscriptionOK{}
}

/*CreateSubscriptionOK handles this case with default header values.

createSubscription response
*/
type CreateSubscriptionOK struct {
	Payload *models.Subscription
}

func (o *CreateSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] createSubscriptionOK  %+v", 200, o.Payload)
}

func (o *CreateSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Subscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
