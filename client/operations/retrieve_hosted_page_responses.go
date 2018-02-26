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

// RetrieveHostedPageReader is a Reader for the RetrieveHostedPage structure.
type RetrieveHostedPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveHostedPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetrieveHostedPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveHostedPageOK creates a RetrieveHostedPageOK with default headers values
func NewRetrieveHostedPageOK() *RetrieveHostedPageOK {
	return &RetrieveHostedPageOK{}
}

/*RetrieveHostedPageOK handles this case with default header values.

retrieveHostedPage response
*/
type RetrieveHostedPageOK struct {
	Payload *models.HostedPageResponse
}

func (o *RetrieveHostedPageOK) Error() string {
	return fmt.Sprintf("[GET /hosted_pages/{id}][%d] retrieveHostedPageOK  %+v", 200, o.Payload)
}

func (o *RetrieveHostedPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostedPageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
