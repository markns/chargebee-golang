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

// UpdateContactCustomerReader is a Reader for the UpdateContactCustomer structure.
type UpdateContactCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateContactCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateContactCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateContactCustomerOK creates a UpdateContactCustomerOK with default headers values
func NewUpdateContactCustomerOK() *UpdateContactCustomerOK {
	return &UpdateContactCustomerOK{}
}

/*UpdateContactCustomerOK handles this case with default header values.

updateContactCustomer response
*/
type UpdateContactCustomerOK struct {
	Payload *models.CustomerResponse
}

func (o *UpdateContactCustomerOK) Error() string {
	return fmt.Sprintf("[POST /customers/{id}/update_contact][%d] updateContactCustomerOK  %+v", 200, o.Payload)
}

func (o *UpdateContactCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
