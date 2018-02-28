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

// RetrieveCouponSetReader is a Reader for the RetrieveCouponSet structure.
type RetrieveCouponSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveCouponSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetrieveCouponSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveCouponSetOK creates a RetrieveCouponSetOK with default headers values
func NewRetrieveCouponSetOK() *RetrieveCouponSetOK {
	return &RetrieveCouponSetOK{}
}

/*RetrieveCouponSetOK handles this case with default header values.

retrieveCouponSet response
*/
type RetrieveCouponSetOK struct {
	Payload *models.CouponSetResponse
}

func (o *RetrieveCouponSetOK) Error() string {
	return fmt.Sprintf("[GET /coupon_sets/{id}][%d] retrieveCouponSetOK  %+v", 200, o.Payload)
}

func (o *RetrieveCouponSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CouponSetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
