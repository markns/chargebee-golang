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

// ListHostedPageReader is a Reader for the ListHostedPage structure.
type ListHostedPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListHostedPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListHostedPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListHostedPageOK creates a ListHostedPageOK with default headers values
func NewListHostedPageOK() *ListHostedPageOK {
	return &ListHostedPageOK{}
}

/*ListHostedPageOK handles this case with default header values.

listHostedPage response
*/
type ListHostedPageOK struct {
	Payload *models.HostedPageResponse
}

func (o *ListHostedPageOK) Error() string {
	return fmt.Sprintf("[GET /hosted_pages][%d] listHostedPageOK  %+v", 200, o.Payload)
}

func (o *ListHostedPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostedPageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
