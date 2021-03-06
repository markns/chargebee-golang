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

// InvoicesForCustomerInvoiceReader is a Reader for the InvoicesForCustomerInvoice structure.
type InvoicesForCustomerInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvoicesForCustomerInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewInvoicesForCustomerInvoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInvoicesForCustomerInvoiceOK creates a InvoicesForCustomerInvoiceOK with default headers values
func NewInvoicesForCustomerInvoiceOK() *InvoicesForCustomerInvoiceOK {
	return &InvoicesForCustomerInvoiceOK{}
}

/*InvoicesForCustomerInvoiceOK handles this case with default header values.

invoicesForCustomerInvoice response
*/
type InvoicesForCustomerInvoiceOK struct {
	Payload *models.InvoicesForCustomerInvoiceOKBody
}

func (o *InvoicesForCustomerInvoiceOK) Error() string {
	return fmt.Sprintf("[GET /customers/{id}/invoices][%d] invoicesForCustomerInvoiceOK  %+v", 200, o.Payload)
}

func (o *InvoicesForCustomerInvoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvoicesForCustomerInvoiceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
