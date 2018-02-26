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

// DeductPromotionalCreditReader is a Reader for the DeductPromotionalCredit structure.
type DeductPromotionalCreditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeductPromotionalCreditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeductPromotionalCreditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeductPromotionalCreditOK creates a DeductPromotionalCreditOK with default headers values
func NewDeductPromotionalCreditOK() *DeductPromotionalCreditOK {
	return &DeductPromotionalCreditOK{}
}

/*DeductPromotionalCreditOK handles this case with default header values.

deductPromotionalCredit response
*/
type DeductPromotionalCreditOK struct {
	Payload *models.PromotionalCreditResponse
}

func (o *DeductPromotionalCreditOK) Error() string {
	return fmt.Sprintf("[POST /promotional_credits/deduct][%d] deductPromotionalCreditOK  %+v", 200, o.Payload)
}

func (o *DeductPromotionalCreditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PromotionalCreditResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}