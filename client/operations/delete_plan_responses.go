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

// DeletePlanReader is a Reader for the DeletePlan structure.
type DeletePlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeletePlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePlanOK creates a DeletePlanOK with default headers values
func NewDeletePlanOK() *DeletePlanOK {
	return &DeletePlanOK{}
}

/*DeletePlanOK handles this case with default header values.

deletePlan response
*/
type DeletePlanOK struct {
	Payload *models.PlanResponse
}

func (o *DeletePlanOK) Error() string {
	return fmt.Sprintf("[POST /plans/{id}/delete][%d] deletePlanOK  %+v", 200, o.Payload)
}

func (o *DeletePlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
