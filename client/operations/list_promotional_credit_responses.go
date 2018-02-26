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

// ListPromotionalCreditReader is a Reader for the ListPromotionalCredit structure.
type ListPromotionalCreditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPromotionalCreditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListPromotionalCreditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListPromotionalCreditOK creates a ListPromotionalCreditOK with default headers values
func NewListPromotionalCreditOK() *ListPromotionalCreditOK {
	return &ListPromotionalCreditOK{}
}

/*ListPromotionalCreditOK handles this case with default header values.

listPromotionalCredit response
*/
type ListPromotionalCreditOK struct {
	Payload *models.PromotionalCredit
}

func (o *ListPromotionalCreditOK) Error() string {
	return fmt.Sprintf("[GET /promotional_credits][%d] listPromotionalCreditOK  %+v", 200, o.Payload)
}

func (o *ListPromotionalCreditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PromotionalCredit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
