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

// UpdateCardPaymentSourceReader is a Reader for the UpdateCardPaymentSource structure.
type UpdateCardPaymentSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCardPaymentSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateCardPaymentSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCardPaymentSourceOK creates a UpdateCardPaymentSourceOK with default headers values
func NewUpdateCardPaymentSourceOK() *UpdateCardPaymentSourceOK {
	return &UpdateCardPaymentSourceOK{}
}

/*UpdateCardPaymentSourceOK handles this case with default header values.

updateCardPaymentSource response
*/
type UpdateCardPaymentSourceOK struct {
	Payload *models.PaymentSource
}

func (o *UpdateCardPaymentSourceOK) Error() string {
	return fmt.Sprintf("[POST /payment_sources/{id}/update_card][%d] updateCardPaymentSourceOK  %+v", 200, o.Payload)
}

func (o *UpdateCardPaymentSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
