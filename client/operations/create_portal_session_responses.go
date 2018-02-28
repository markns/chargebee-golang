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

// CreatePortalSessionReader is a Reader for the CreatePortalSession structure.
type CreatePortalSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePortalSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreatePortalSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePortalSessionOK creates a CreatePortalSessionOK with default headers values
func NewCreatePortalSessionOK() *CreatePortalSessionOK {
	return &CreatePortalSessionOK{}
}

/*CreatePortalSessionOK handles this case with default header values.

createPortalSession response
*/
type CreatePortalSessionOK struct {
	Payload *models.PortalSessionResponse
}

func (o *CreatePortalSessionOK) Error() string {
	return fmt.Sprintf("[POST /portal_sessions][%d] createPortalSessionOK  %+v", 200, o.Payload)
}

func (o *CreatePortalSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortalSessionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
