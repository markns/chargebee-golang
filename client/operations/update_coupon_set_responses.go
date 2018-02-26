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

// UpdateCouponSetReader is a Reader for the UpdateCouponSet structure.
type UpdateCouponSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCouponSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateCouponSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCouponSetOK creates a UpdateCouponSetOK with default headers values
func NewUpdateCouponSetOK() *UpdateCouponSetOK {
	return &UpdateCouponSetOK{}
}

/*UpdateCouponSetOK handles this case with default header values.

updateCouponSet response
*/
type UpdateCouponSetOK struct {
	Payload *models.CouponSetResponse
}

func (o *UpdateCouponSetOK) Error() string {
	return fmt.Sprintf("[POST /coupon_sets/{id}/update][%d] updateCouponSetOK  %+v", 200, o.Payload)
}

func (o *UpdateCouponSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CouponSetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}