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

// RemoveCreditNoteInvoiceReader is a Reader for the RemoveCreditNoteInvoice structure.
type RemoveCreditNoteInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveCreditNoteInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveCreditNoteInvoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveCreditNoteInvoiceOK creates a RemoveCreditNoteInvoiceOK with default headers values
func NewRemoveCreditNoteInvoiceOK() *RemoveCreditNoteInvoiceOK {
	return &RemoveCreditNoteInvoiceOK{}
}

/*RemoveCreditNoteInvoiceOK handles this case with default header values.

removeCreditNoteInvoice response
*/
type RemoveCreditNoteInvoiceOK struct {
	Payload *models.InvoiceResponse
}

func (o *RemoveCreditNoteInvoiceOK) Error() string {
	return fmt.Sprintf("[POST /invoices/{id}/remove_credit_note][%d] removeCreditNoteInvoiceOK  %+v", 200, o.Payload)
}

func (o *RemoveCreditNoteInvoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
