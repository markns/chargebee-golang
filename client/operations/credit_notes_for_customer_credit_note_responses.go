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

// CreditNotesForCustomerCreditNoteReader is a Reader for the CreditNotesForCustomerCreditNote structure.
type CreditNotesForCustomerCreditNoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreditNotesForCustomerCreditNoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreditNotesForCustomerCreditNoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreditNotesForCustomerCreditNoteOK creates a CreditNotesForCustomerCreditNoteOK with default headers values
func NewCreditNotesForCustomerCreditNoteOK() *CreditNotesForCustomerCreditNoteOK {
	return &CreditNotesForCustomerCreditNoteOK{}
}

/*CreditNotesForCustomerCreditNoteOK handles this case with default header values.

creditNotesForCustomerCreditNote response
*/
type CreditNotesForCustomerCreditNoteOK struct {
	Payload *models.CreditNoteResponse
}

func (o *CreditNotesForCustomerCreditNoteOK) Error() string {
	return fmt.Sprintf("[GET /customers/{id}/credit_notes][%d] creditNotesForCustomerCreditNoteOK  %+v", 200, o.Payload)
}

func (o *CreditNotesForCustomerCreditNoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreditNoteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
