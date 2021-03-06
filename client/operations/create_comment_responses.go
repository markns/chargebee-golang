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

// CreateCommentReader is a Reader for the CreateComment structure.
type CreateCommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateCommentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateCommentOK creates a CreateCommentOK with default headers values
func NewCreateCommentOK() *CreateCommentOK {
	return &CreateCommentOK{}
}

/*CreateCommentOK handles this case with default header values.

createComment response
*/
type CreateCommentOK struct {
	Payload *models.CommentResponse
}

func (o *CreateCommentOK) Error() string {
	return fmt.Sprintf("[POST /comments][%d] createCommentOK  %+v", 200, o.Payload)
}

func (o *CreateCommentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
