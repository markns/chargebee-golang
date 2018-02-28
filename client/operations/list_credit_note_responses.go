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

// ListCreditNoteReader is a Reader for the ListCreditNote structure.
type ListCreditNoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCreditNoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListCreditNoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListCreditNoteOK creates a ListCreditNoteOK with default headers values
func NewListCreditNoteOK() *ListCreditNoteOK {
	return &ListCreditNoteOK{}
}

/*ListCreditNoteOK handles this case with default header values.

listCreditNote response
*/
type ListCreditNoteOK struct {
	Payload *models.ListCreditNoteOKBody
}

func (o *ListCreditNoteOK) Error() string {
	return fmt.Sprintf("[GET /credit_notes][%d] listCreditNoteOK  %+v", 200, o.Payload)
}

func (o *ListCreditNoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListCreditNoteOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
