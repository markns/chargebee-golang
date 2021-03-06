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

// RetrievePortalSessionReader is a Reader for the RetrievePortalSession structure.
type RetrievePortalSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrievePortalSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetrievePortalSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrievePortalSessionOK creates a RetrievePortalSessionOK with default headers values
func NewRetrievePortalSessionOK() *RetrievePortalSessionOK {
	return &RetrievePortalSessionOK{}
}

/*RetrievePortalSessionOK handles this case with default header values.

retrievePortalSession response
*/
type RetrievePortalSessionOK struct {
	Payload *models.PortalSessionResponse
}

func (o *RetrievePortalSessionOK) Error() string {
	return fmt.Sprintf("[GET /portal_sessions/{id}][%d] retrievePortalSessionOK  %+v", 200, o.Payload)
}

func (o *RetrievePortalSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortalSessionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
