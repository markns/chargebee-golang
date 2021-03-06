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

// UpdateAddressReader is a Reader for the UpdateAddress structure.
type UpdateAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAddressOK creates a UpdateAddressOK with default headers values
func NewUpdateAddressOK() *UpdateAddressOK {
	return &UpdateAddressOK{}
}

/*UpdateAddressOK handles this case with default header values.

updateAddress response
*/
type UpdateAddressOK struct {
	Payload *models.AddressResponse
}

func (o *UpdateAddressOK) Error() string {
	return fmt.Sprintf("[POST /addresses][%d] updateAddressOK  %+v", 200, o.Payload)
}

func (o *UpdateAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddressResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
