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

// UpdateAddonReader is a Reader for the UpdateAddon structure.
type UpdateAddonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAddonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateAddonOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAddonOK creates a UpdateAddonOK with default headers values
func NewUpdateAddonOK() *UpdateAddonOK {
	return &UpdateAddonOK{}
}

/*UpdateAddonOK handles this case with default header values.

updateAddon response
*/
type UpdateAddonOK struct {
	Payload *models.AddonResponse
}

func (o *UpdateAddonOK) Error() string {
	return fmt.Sprintf("[POST /addons/{id}][%d] updateAddonOK  %+v", 200, o.Payload)
}

func (o *UpdateAddonOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddonResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
