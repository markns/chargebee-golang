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

// CancelSubscriptionReader is a Reader for the CancelSubscription structure.
type CancelSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCancelSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelSubscriptionOK creates a CancelSubscriptionOK with default headers values
func NewCancelSubscriptionOK() *CancelSubscriptionOK {
	return &CancelSubscriptionOK{}
}

/*CancelSubscriptionOK handles this case with default header values.

cancelSubscription response
*/
type CancelSubscriptionOK struct {
	Payload *models.SubscriptionResponse
}

func (o *CancelSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{id}/cancel][%d] cancelSubscriptionOK  %+v", 200, o.Payload)
}

func (o *CancelSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
