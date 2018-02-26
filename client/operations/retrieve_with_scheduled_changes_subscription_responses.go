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

// RetrieveWithScheduledChangesSubscriptionReader is a Reader for the RetrieveWithScheduledChangesSubscription structure.
type RetrieveWithScheduledChangesSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveWithScheduledChangesSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetrieveWithScheduledChangesSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveWithScheduledChangesSubscriptionOK creates a RetrieveWithScheduledChangesSubscriptionOK with default headers values
func NewRetrieveWithScheduledChangesSubscriptionOK() *RetrieveWithScheduledChangesSubscriptionOK {
	return &RetrieveWithScheduledChangesSubscriptionOK{}
}

/*RetrieveWithScheduledChangesSubscriptionOK handles this case with default header values.

retrieveWithScheduledChangesSubscription response
*/
type RetrieveWithScheduledChangesSubscriptionOK struct {
	Payload *models.SubscriptionResponse
}

func (o *RetrieveWithScheduledChangesSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{id}/retrieve_with_scheduled_changes][%d] retrieveWithScheduledChangesSubscriptionOK  %+v", 200, o.Payload)
}

func (o *RetrieveWithScheduledChangesSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
