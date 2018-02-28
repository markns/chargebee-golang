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

// AcknowledgeHostedPageReader is a Reader for the AcknowledgeHostedPage structure.
type AcknowledgeHostedPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AcknowledgeHostedPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAcknowledgeHostedPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAcknowledgeHostedPageOK creates a AcknowledgeHostedPageOK with default headers values
func NewAcknowledgeHostedPageOK() *AcknowledgeHostedPageOK {
	return &AcknowledgeHostedPageOK{}
}

/*AcknowledgeHostedPageOK handles this case with default header values.

acknowledgeHostedPage response
*/
type AcknowledgeHostedPageOK struct {
	Payload *models.HostedPageResponse
}

func (o *AcknowledgeHostedPageOK) Error() string {
	return fmt.Sprintf("[POST /hosted_pages/{id}/acknowledge][%d] acknowledgeHostedPageOK  %+v", 200, o.Payload)
}

func (o *AcknowledgeHostedPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostedPageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
