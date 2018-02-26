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

// ListCommentReader is a Reader for the ListComment structure.
type ListCommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListCommentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListCommentOK creates a ListCommentOK with default headers values
func NewListCommentOK() *ListCommentOK {
	return &ListCommentOK{}
}

/*ListCommentOK handles this case with default header values.

listComment response
*/
type ListCommentOK struct {
	Payload *models.CommentResponse
}

func (o *ListCommentOK) Error() string {
	return fmt.Sprintf("[GET /comments][%d] listCommentOK  %+v", 200, o.Payload)
}

func (o *ListCommentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
