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

// CreateSubscriptionEstimateReader is a Reader for the CreateSubscriptionEstimate structure.
type CreateSubscriptionEstimateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSubscriptionEstimateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSubscriptionEstimateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSubscriptionEstimateOK creates a CreateSubscriptionEstimateOK with default headers values
func NewCreateSubscriptionEstimateOK() *CreateSubscriptionEstimateOK {
	return &CreateSubscriptionEstimateOK{}
}

/*CreateSubscriptionEstimateOK handles this case with default header values.

createSubscriptionEstimate response
*/
type CreateSubscriptionEstimateOK struct {
	Payload *models.EstimateResponse
}

func (o *CreateSubscriptionEstimateOK) Error() string {
	return fmt.Sprintf("[POST /estimates/create_subscription][%d] createSubscriptionEstimateOK  %+v", 200, o.Payload)
}

func (o *CreateSubscriptionEstimateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EstimateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
