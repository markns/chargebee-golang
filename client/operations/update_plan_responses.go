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

// UpdatePlanReader is a Reader for the UpdatePlan structure.
type UpdatePlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdatePlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePlanOK creates a UpdatePlanOK with default headers values
func NewUpdatePlanOK() *UpdatePlanOK {
	return &UpdatePlanOK{}
}

/*UpdatePlanOK handles this case with default header values.

updatePlan response
*/
type UpdatePlanOK struct {
	Payload *models.Plan
}

func (o *UpdatePlanOK) Error() string {
	return fmt.Sprintf("[POST /plans/{id}][%d] updatePlanOK  %+v", 200, o.Payload)
}

func (o *UpdatePlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Plan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
