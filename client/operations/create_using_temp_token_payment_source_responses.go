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

// CreateUsingTempTokenPaymentSourceReader is a Reader for the CreateUsingTempTokenPaymentSource structure.
type CreateUsingTempTokenPaymentSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUsingTempTokenPaymentSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateUsingTempTokenPaymentSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUsingTempTokenPaymentSourceOK creates a CreateUsingTempTokenPaymentSourceOK with default headers values
func NewCreateUsingTempTokenPaymentSourceOK() *CreateUsingTempTokenPaymentSourceOK {
	return &CreateUsingTempTokenPaymentSourceOK{}
}

/*CreateUsingTempTokenPaymentSourceOK handles this case with default header values.

createUsingTempTokenPaymentSource response
*/
type CreateUsingTempTokenPaymentSourceOK struct {
	Payload *models.PaymentSourceResponse
}

func (o *CreateUsingTempTokenPaymentSourceOK) Error() string {
	return fmt.Sprintf("[POST /payment_sources/create_using_temp_token][%d] createUsingTempTokenPaymentSourceOK  %+v", 200, o.Payload)
}

func (o *CreateUsingTempTokenPaymentSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentSourceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
