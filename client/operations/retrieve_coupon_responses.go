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

// RetrieveCouponReader is a Reader for the RetrieveCoupon structure.
type RetrieveCouponReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveCouponReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetrieveCouponOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveCouponOK creates a RetrieveCouponOK with default headers values
func NewRetrieveCouponOK() *RetrieveCouponOK {
	return &RetrieveCouponOK{}
}

/*RetrieveCouponOK handles this case with default header values.

retrieveCoupon response
*/
type RetrieveCouponOK struct {
	Payload *models.CouponResponse
}

func (o *RetrieveCouponOK) Error() string {
	return fmt.Sprintf("[GET /coupons/{id}][%d] retrieveCouponOK  %+v", 200, o.Payload)
}

func (o *RetrieveCouponOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CouponResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
