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

// CreateUsingPermanentTokenPaymentSourceReader is a Reader for the CreateUsingPermanentTokenPaymentSource structure.
type CreateUsingPermanentTokenPaymentSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUsingPermanentTokenPaymentSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateUsingPermanentTokenPaymentSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUsingPermanentTokenPaymentSourceOK creates a CreateUsingPermanentTokenPaymentSourceOK with default headers values
func NewCreateUsingPermanentTokenPaymentSourceOK() *CreateUsingPermanentTokenPaymentSourceOK {
	return &CreateUsingPermanentTokenPaymentSourceOK{}
}

/*CreateUsingPermanentTokenPaymentSourceOK handles this case with default header values.

createUsingPermanentTokenPaymentSource response
*/
type CreateUsingPermanentTokenPaymentSourceOK struct {
	Payload *models.PaymentSourceResponse
}

func (o *CreateUsingPermanentTokenPaymentSourceOK) Error() string {
	return fmt.Sprintf("[POST /payment_sources/create_using_permanent_token][%d] createUsingPermanentTokenPaymentSourceOK  %+v", 200, o.Payload)
}

func (o *CreateUsingPermanentTokenPaymentSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentSourceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}