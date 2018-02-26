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

// ListOrderReader is a Reader for the ListOrder structure.
type ListOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListOrderOK creates a ListOrderOK with default headers values
func NewListOrderOK() *ListOrderOK {
	return &ListOrderOK{}
}

/*ListOrderOK handles this case with default header values.

listOrder response
*/
type ListOrderOK struct {
	Payload *models.OrderResponse
}

func (o *ListOrderOK) Error() string {
	return fmt.Sprintf("[GET /orders][%d] listOrderOK  %+v", 200, o.Payload)
}

func (o *ListOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
