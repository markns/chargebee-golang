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

// RetrieveInvoiceReader is a Reader for the RetrieveInvoice structure.
type RetrieveInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetrieveInvoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveInvoiceOK creates a RetrieveInvoiceOK with default headers values
func NewRetrieveInvoiceOK() *RetrieveInvoiceOK {
	return &RetrieveInvoiceOK{}
}

/*RetrieveInvoiceOK handles this case with default header values.

retrieveInvoice response
*/
type RetrieveInvoiceOK struct {
	Payload *models.InvoiceResponse
}

func (o *RetrieveInvoiceOK) Error() string {
	return fmt.Sprintf("[GET /invoices/{id}][%d] retrieveInvoiceOK  %+v", 200, o.Payload)
}

func (o *RetrieveInvoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
