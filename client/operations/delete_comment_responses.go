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

// DeleteCommentReader is a Reader for the DeleteComment structure.
type DeleteCommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCommentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCommentOK creates a DeleteCommentOK with default headers values
func NewDeleteCommentOK() *DeleteCommentOK {
	return &DeleteCommentOK{}
}

/*DeleteCommentOK handles this case with default header values.

deleteComment response
*/
type DeleteCommentOK struct {
	Payload *models.CommentResponse
}

func (o *DeleteCommentOK) Error() string {
	return fmt.Sprintf("[POST /comments/{id}/delete][%d] deleteCommentOK  %+v", 200, o.Payload)
}

func (o *DeleteCommentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
