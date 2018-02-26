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

// ActivatePortalSessionReader is a Reader for the ActivatePortalSession structure.
type ActivatePortalSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActivatePortalSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewActivatePortalSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActivatePortalSessionOK creates a ActivatePortalSessionOK with default headers values
func NewActivatePortalSessionOK() *ActivatePortalSessionOK {
	return &ActivatePortalSessionOK{}
}

/*ActivatePortalSessionOK handles this case with default header values.

activatePortalSession response
*/
type ActivatePortalSessionOK struct {
	Payload *models.PortalSessionResponse
}

func (o *ActivatePortalSessionOK) Error() string {
	return fmt.Sprintf("[POST /portal_sessions/{id}/activate][%d] activatePortalSessionOK  %+v", 200, o.Payload)
}

func (o *ActivatePortalSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortalSessionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
