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

// SetPromotionalCreditReader is a Reader for the SetPromotionalCredit structure.
type SetPromotionalCreditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetPromotionalCreditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetPromotionalCreditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetPromotionalCreditOK creates a SetPromotionalCreditOK with default headers values
func NewSetPromotionalCreditOK() *SetPromotionalCreditOK {
	return &SetPromotionalCreditOK{}
}

/*SetPromotionalCreditOK handles this case with default header values.

setPromotionalCredit response
*/
type SetPromotionalCreditOK struct {
	Payload *models.PromotionalCreditResponse
}

func (o *SetPromotionalCreditOK) Error() string {
	return fmt.Sprintf("[POST /promotional_credits/set][%d] setPromotionalCreditOK  %+v", 200, o.Payload)
}

func (o *SetPromotionalCreditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PromotionalCreditResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}