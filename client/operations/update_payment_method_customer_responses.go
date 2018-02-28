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

// UpdatePaymentMethodCustomerReader is a Reader for the UpdatePaymentMethodCustomer structure.
type UpdatePaymentMethodCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePaymentMethodCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdatePaymentMethodCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePaymentMethodCustomerOK creates a UpdatePaymentMethodCustomerOK with default headers values
func NewUpdatePaymentMethodCustomerOK() *UpdatePaymentMethodCustomerOK {
	return &UpdatePaymentMethodCustomerOK{}
}

/*UpdatePaymentMethodCustomerOK handles this case with default header values.

updatePaymentMethodCustomer response
*/
type UpdatePaymentMethodCustomerOK struct {
	Payload *models.CustomerResponse
}

func (o *UpdatePaymentMethodCustomerOK) Error() string {
	return fmt.Sprintf("[POST /customers/{id}/update_payment_method][%d] updatePaymentMethodCustomerOK  %+v", 200, o.Payload)
}

func (o *UpdatePaymentMethodCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
