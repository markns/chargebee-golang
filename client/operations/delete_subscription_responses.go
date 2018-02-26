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

// DeleteSubscriptionReader is a Reader for the DeleteSubscription structure.
type DeleteSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSubscriptionOK creates a DeleteSubscriptionOK with default headers values
func NewDeleteSubscriptionOK() *DeleteSubscriptionOK {
	return &DeleteSubscriptionOK{}
}

/*DeleteSubscriptionOK handles this case with default header values.

deleteSubscription response
*/
type DeleteSubscriptionOK struct {
	Payload *models.SubscriptionResponse
}

func (o *DeleteSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{id}/delete][%d] deleteSubscriptionOK  %+v", 200, o.Payload)
}

func (o *DeleteSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
