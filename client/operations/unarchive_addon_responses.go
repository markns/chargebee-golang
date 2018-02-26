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

// UnarchiveAddonReader is a Reader for the UnarchiveAddon structure.
type UnarchiveAddonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnarchiveAddonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUnarchiveAddonOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUnarchiveAddonOK creates a UnarchiveAddonOK with default headers values
func NewUnarchiveAddonOK() *UnarchiveAddonOK {
	return &UnarchiveAddonOK{}
}

/*UnarchiveAddonOK handles this case with default header values.

unarchiveAddon response
*/
type UnarchiveAddonOK struct {
	Payload *models.Addon
}

func (o *UnarchiveAddonOK) Error() string {
	return fmt.Sprintf("[POST /addons/{id}/unarchive][%d] unarchiveAddonOK  %+v", 200, o.Payload)
}

func (o *UnarchiveAddonOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Addon)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
