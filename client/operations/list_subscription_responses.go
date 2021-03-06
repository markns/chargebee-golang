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

// ListSubscriptionReader is a Reader for the ListSubscription structure.
type ListSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSubscriptionOK creates a ListSubscriptionOK with default headers values
func NewListSubscriptionOK() *ListSubscriptionOK {
	return &ListSubscriptionOK{}
}

/*ListSubscriptionOK handles this case with default header values.

listSubscription response
*/
type ListSubscriptionOK struct {
	Payload *models.ListSubscriptionOKBody
}

func (o *ListSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions][%d] listSubscriptionOK  %+v", 200, o.Payload)
}

func (o *ListSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListSubscriptionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
