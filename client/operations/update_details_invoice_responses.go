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

// UpdateDetailsInvoiceReader is a Reader for the UpdateDetailsInvoice structure.
type UpdateDetailsInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDetailsInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateDetailsInvoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateDetailsInvoiceOK creates a UpdateDetailsInvoiceOK with default headers values
func NewUpdateDetailsInvoiceOK() *UpdateDetailsInvoiceOK {
	return &UpdateDetailsInvoiceOK{}
}

/*UpdateDetailsInvoiceOK handles this case with default header values.

updateDetailsInvoice response
*/
type UpdateDetailsInvoiceOK struct {
	Payload *models.InvoiceResponse
}

func (o *UpdateDetailsInvoiceOK) Error() string {
	return fmt.Sprintf("[POST /invoices/{id}/update_details][%d] updateDetailsInvoiceOK  %+v", 200, o.Payload)
}

func (o *UpdateDetailsInvoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
