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

// ImportSubscriptionSubscriptionReader is a Reader for the ImportSubscriptionSubscription structure.
type ImportSubscriptionSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportSubscriptionSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewImportSubscriptionSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImportSubscriptionSubscriptionOK creates a ImportSubscriptionSubscriptionOK with default headers values
func NewImportSubscriptionSubscriptionOK() *ImportSubscriptionSubscriptionOK {
	return &ImportSubscriptionSubscriptionOK{}
}

/*ImportSubscriptionSubscriptionOK handles this case with default header values.

importSubscriptionSubscription response
*/
type ImportSubscriptionSubscriptionOK struct {
	Payload *models.SubscriptionResponse
}

func (o *ImportSubscriptionSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/import_subscription][%d] importSubscriptionSubscriptionOK  %+v", 200, o.Payload)
}

func (o *ImportSubscriptionSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
