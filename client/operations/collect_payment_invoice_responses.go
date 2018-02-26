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

// CollectPaymentInvoiceReader is a Reader for the CollectPaymentInvoice structure.
type CollectPaymentInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectPaymentInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCollectPaymentInvoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCollectPaymentInvoiceOK creates a CollectPaymentInvoiceOK with default headers values
func NewCollectPaymentInvoiceOK() *CollectPaymentInvoiceOK {
	return &CollectPaymentInvoiceOK{}
}

/*CollectPaymentInvoiceOK handles this case with default header values.

collectPaymentInvoice response
*/
type CollectPaymentInvoiceOK struct {
	Payload *models.InvoiceResponse
}

func (o *CollectPaymentInvoiceOK) Error() string {
	return fmt.Sprintf("[POST /invoices/{id}/collect_payment][%d] collectPaymentInvoiceOK  %+v", 200, o.Payload)
}

func (o *CollectPaymentInvoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
