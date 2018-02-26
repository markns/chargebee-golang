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

// RetrieveCouponCodeReader is a Reader for the RetrieveCouponCode structure.
type RetrieveCouponCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveCouponCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetrieveCouponCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveCouponCodeOK creates a RetrieveCouponCodeOK with default headers values
func NewRetrieveCouponCodeOK() *RetrieveCouponCodeOK {
	return &RetrieveCouponCodeOK{}
}

/*RetrieveCouponCodeOK handles this case with default header values.

retrieveCouponCode response
*/
type RetrieveCouponCodeOK struct {
	Payload *models.CouponCodeResponse
}

func (o *RetrieveCouponCodeOK) Error() string {
	return fmt.Sprintf("[GET /coupon_codes/{id}][%d] retrieveCouponCodeOK  %+v", 200, o.Payload)
}

func (o *RetrieveCouponCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CouponCodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
