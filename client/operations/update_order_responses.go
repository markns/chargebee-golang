// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"bitbucket.org/gridarrow/chargebee-api/models"
)

// UpdateOrderReader is a Reader for the UpdateOrder structure.
type UpdateOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateOrderOK creates a UpdateOrderOK with default headers values
func NewUpdateOrderOK() *UpdateOrderOK {
	return &UpdateOrderOK{}
}

/*UpdateOrderOK handles this case with default header values.

updateOrder response
*/
type UpdateOrderOK struct {
	Payload *models.Order
}

func (o *UpdateOrderOK) Error() string {
	return fmt.Sprintf("[POST /orders/{id}][%d] updateOrderOK  %+v", 200, o.Payload)
}

func (o *UpdateOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
