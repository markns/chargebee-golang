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

// UpdateCardForCustomerCardReader is a Reader for the UpdateCardForCustomerCard structure.
type UpdateCardForCustomerCardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCardForCustomerCardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateCardForCustomerCardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCardForCustomerCardOK creates a UpdateCardForCustomerCardOK with default headers values
func NewUpdateCardForCustomerCardOK() *UpdateCardForCustomerCardOK {
	return &UpdateCardForCustomerCardOK{}
}

/*UpdateCardForCustomerCardOK handles this case with default header values.

updateCardForCustomerCard response
*/
type UpdateCardForCustomerCardOK struct {
	Payload *models.CardResponse
}

func (o *UpdateCardForCustomerCardOK) Error() string {
	return fmt.Sprintf("[POST /customers/{id}/credit_card][%d] updateCardForCustomerCardOK  %+v", 200, o.Payload)
}

func (o *UpdateCardForCustomerCardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CardResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
