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

// CreateCouponReader is a Reader for the CreateCoupon structure.
type CreateCouponReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCouponReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateCouponOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateCouponOK creates a CreateCouponOK with default headers values
func NewCreateCouponOK() *CreateCouponOK {
	return &CreateCouponOK{}
}

/*CreateCouponOK handles this case with default header values.

createCoupon response
*/
type CreateCouponOK struct {
	Payload *models.Coupon
}

func (o *CreateCouponOK) Error() string {
	return fmt.Sprintf("[POST /coupons][%d] createCouponOK  %+v", 200, o.Payload)
}

func (o *CreateCouponOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Coupon)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
