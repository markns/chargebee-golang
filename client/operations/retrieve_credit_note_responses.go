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

// RetrieveCreditNoteReader is a Reader for the RetrieveCreditNote structure.
type RetrieveCreditNoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveCreditNoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetrieveCreditNoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveCreditNoteOK creates a RetrieveCreditNoteOK with default headers values
func NewRetrieveCreditNoteOK() *RetrieveCreditNoteOK {
	return &RetrieveCreditNoteOK{}
}

/*RetrieveCreditNoteOK handles this case with default header values.

retrieveCreditNote response
*/
type RetrieveCreditNoteOK struct {
	Payload *models.CreditNoteResponse
}

func (o *RetrieveCreditNoteOK) Error() string {
	return fmt.Sprintf("[GET /credit_notes/{id}][%d] retrieveCreditNoteOK  %+v", 200, o.Payload)
}

func (o *RetrieveCreditNoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreditNoteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
