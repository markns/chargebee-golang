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

// CreateCardPaymentSourceReader is a Reader for the CreateCardPaymentSource structure.
type CreateCardPaymentSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCardPaymentSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateCardPaymentSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateCardPaymentSourceOK creates a CreateCardPaymentSourceOK with default headers values
func NewCreateCardPaymentSourceOK() *CreateCardPaymentSourceOK {
	return &CreateCardPaymentSourceOK{}
}

/*CreateCardPaymentSourceOK handles this case with default header values.

createCardPaymentSource response
*/
type CreateCardPaymentSourceOK struct {
	Payload *models.PaymentSourceResponse
}

func (o *CreateCardPaymentSourceOK) Error() string {
	return fmt.Sprintf("[POST /payment_sources/create_card][%d] createCardPaymentSourceOK  %+v", 200, o.Payload)
}

func (o *CreateCardPaymentSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentSourceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
