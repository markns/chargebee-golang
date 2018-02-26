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

// DeletePaymentSourceReader is a Reader for the DeletePaymentSource structure.
type DeletePaymentSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePaymentSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeletePaymentSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePaymentSourceOK creates a DeletePaymentSourceOK with default headers values
func NewDeletePaymentSourceOK() *DeletePaymentSourceOK {
	return &DeletePaymentSourceOK{}
}

/*DeletePaymentSourceOK handles this case with default header values.

deletePaymentSource response
*/
type DeletePaymentSourceOK struct {
	Payload *models.PaymentSourceResponse
}

func (o *DeletePaymentSourceOK) Error() string {
	return fmt.Sprintf("[POST /payment_sources/{id}/delete][%d] deletePaymentSourceOK  %+v", 200, o.Payload)
}

func (o *DeletePaymentSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentSourceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
