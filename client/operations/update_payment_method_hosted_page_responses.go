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

// UpdatePaymentMethodHostedPageReader is a Reader for the UpdatePaymentMethodHostedPage structure.
type UpdatePaymentMethodHostedPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePaymentMethodHostedPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdatePaymentMethodHostedPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePaymentMethodHostedPageOK creates a UpdatePaymentMethodHostedPageOK with default headers values
func NewUpdatePaymentMethodHostedPageOK() *UpdatePaymentMethodHostedPageOK {
	return &UpdatePaymentMethodHostedPageOK{}
}

/*UpdatePaymentMethodHostedPageOK handles this case with default header values.

updatePaymentMethodHostedPage response
*/
type UpdatePaymentMethodHostedPageOK struct {
	Payload *models.HostedPageResponse
}

func (o *UpdatePaymentMethodHostedPageOK) Error() string {
	return fmt.Sprintf("[POST /hosted_pages/update_payment_method][%d] updatePaymentMethodHostedPageOK  %+v", 200, o.Payload)
}

func (o *UpdatePaymentMethodHostedPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostedPageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}