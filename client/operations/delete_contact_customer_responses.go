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

// DeleteContactCustomerReader is a Reader for the DeleteContactCustomer structure.
type DeleteContactCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteContactCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteContactCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteContactCustomerOK creates a DeleteContactCustomerOK with default headers values
func NewDeleteContactCustomerOK() *DeleteContactCustomerOK {
	return &DeleteContactCustomerOK{}
}

/*DeleteContactCustomerOK handles this case with default header values.

deleteContactCustomer response
*/
type DeleteContactCustomerOK struct {
	Payload *models.Customer
}

func (o *DeleteContactCustomerOK) Error() string {
	return fmt.Sprintf("[POST /customers/{id}/delete_contact][%d] deleteContactCustomerOK  %+v", 200, o.Payload)
}

func (o *DeleteContactCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
