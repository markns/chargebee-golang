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

// DeleteCouponSetReader is a Reader for the DeleteCouponSet structure.
type DeleteCouponSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCouponSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCouponSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCouponSetOK creates a DeleteCouponSetOK with default headers values
func NewDeleteCouponSetOK() *DeleteCouponSetOK {
	return &DeleteCouponSetOK{}
}

/*DeleteCouponSetOK handles this case with default header values.

deleteCouponSet response
*/
type DeleteCouponSetOK struct {
	Payload *models.CouponSetResponse
}

func (o *DeleteCouponSetOK) Error() string {
	return fmt.Sprintf("[POST /coupon_sets/{id}/delete][%d] deleteCouponSetOK  %+v", 200, o.Payload)
}

func (o *DeleteCouponSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CouponSetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}