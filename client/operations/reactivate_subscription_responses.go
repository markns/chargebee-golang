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

// ReactivateSubscriptionReader is a Reader for the ReactivateSubscription structure.
type ReactivateSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReactivateSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReactivateSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReactivateSubscriptionOK creates a ReactivateSubscriptionOK with default headers values
func NewReactivateSubscriptionOK() *ReactivateSubscriptionOK {
	return &ReactivateSubscriptionOK{}
}

/*ReactivateSubscriptionOK handles this case with default header values.

reactivateSubscription response
*/
type ReactivateSubscriptionOK struct {
	Payload *models.Subscription
}

func (o *ReactivateSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{id}/reactivate][%d] reactivateSubscriptionOK  %+v", 200, o.Payload)
}

func (o *ReactivateSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Subscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
