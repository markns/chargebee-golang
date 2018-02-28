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

// CreatePlanReader is a Reader for the CreatePlan structure.
type CreatePlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreatePlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePlanOK creates a CreatePlanOK with default headers values
func NewCreatePlanOK() *CreatePlanOK {
	return &CreatePlanOK{}
}

/*CreatePlanOK handles this case with default header values.

createPlan response
*/
type CreatePlanOK struct {
	Payload *models.PlanResponse
}

func (o *CreatePlanOK) Error() string {
	return fmt.Sprintf("[POST /plans][%d] createPlanOK  %+v", 200, o.Payload)
}

func (o *CreatePlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
