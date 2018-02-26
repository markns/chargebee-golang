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

// CreateInvoiceReader is a Reader for the CreateInvoice structure.
type CreateInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateInvoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateInvoiceOK creates a CreateInvoiceOK with default headers values
func NewCreateInvoiceOK() *CreateInvoiceOK {
	return &CreateInvoiceOK{}
}

/*CreateInvoiceOK handles this case with default header values.

createInvoice response
*/
type CreateInvoiceOK struct {
	Payload *models.InvoiceResponse
}

func (o *CreateInvoiceOK) Error() string {
	return fmt.Sprintf("[POST /invoices][%d] createInvoiceOK  %+v", 200, o.Payload)
}

func (o *CreateInvoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}